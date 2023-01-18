package json

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

func Serve(project string, repo string, bundle GitHubBundle) {

	fmt.Println("Server Starting")

	mux := http.NewServeMux()

	mux.HandleFunc("/home", homeHandler)
	mux.HandleFunc("/bugreports", func(w http.ResponseWriter, r *http.Request) {
		bugReportHandler(w, r, bundle.Issues)
	})
	mux.HandleFunc("/milestones", func(w http.ResponseWriter, r *http.Request) {
		milestonesHandler(w, r, bundle.Milestones)
	})
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		usersHandler(w, r, bundle.Users)
	})

	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	const templ = `
	<h1>Home</h1>
	<ul>
	  <li><a href="/bugreports">Bug Reports</a></li>
	  <li><a href="/milestones">Milestones</a></li>
	  <li><a href="/users">Users</a></li>
	</ul>
`
	report, err := template.New("home").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ)
	if err != nil {
		log.Fatal(err)
	}

	report.Execute(w, IssueBodyParams{})
}

func bugReportHandler(w http.ResponseWriter, r *http.Request, issues []GitHubIssue) {
	const templ = `
	<h1>Bug Reports</h1>
	<table>
	<tr style='text-align: left'>
	<th>#</th>
	<th>State</th>
	<th>User</th>
	<th>Title</th>
	</tr>
	</table>
    <footer>
    <p><a href="/home">Home</a></p>
	</footer> 
`
	report, err := template.New("bugreports").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ)
	if err != nil {
		log.Fatal(err)
	}

	report.Execute(w, IssueBodyParams{})
}

func milestonesHandler(w http.ResponseWriter, r *http.Request, milestones []GitHubMilestone) {
	const templ = `
	<h1>Miilestones</h1>
	<table>
	<tr style='text-align: left'>
	<th>#</th>
	<th>State</th>
	<th>User</th>
	<th>Title</th>
	</tr>
	</table>
    <footer>
    <p><a href="/home">Home</a></p>
	</footer> 
`
	report, err := template.New("milestones").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ)
	if err != nil {
		log.Fatal(err)
	}

	report.Execute(w, IssueBodyParams{})
}

func usersHandler(w http.ResponseWriter, r *http.Request, users []GitHubUser) {
	const templ = `
	<h1>Users</h1>
	<table>
	<tr style='text-align: left'>
	<th>#</th>
	<th>Username</th>
	<th>Email</th>
	<th>Website</th>
	</tr>
	</table>
    <footer>
    <p><a href="/home">Home</a></p>
	</footer> 
`
	report, err := template.New("users").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ)
	if err != nil {
		log.Fatal(err)
	}

	report.Execute(w, IssueBodyParams{})
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 2)
}
