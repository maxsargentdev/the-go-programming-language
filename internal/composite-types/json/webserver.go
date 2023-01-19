package json

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)


func QueryGitHubAPI() GitHubBundle {
	return GitHubBundle{
		Issues: []GitHubIssue{},
		Users: []GitHubUser{
			{Id: 0, Login: "test_user", HtmlUrl: "test_user.github.io"},
		},
		Milestones: []GitHubMilestone{},
	}
}

func Serve(project string, repo string) {

	fmt.Println("Querying GitHub API....")

	bundle := QueryGitHubAPI()

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

	http.ListenAndServe(":8080", mux)
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

	report.Execute(w, issues)
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

	report.Execute(w, milestones)
}

func usersHandler(w http.ResponseWriter, r *http.Request, users []GitHubUser) {
	const templ = `
	<h1>Users</h1>
	<table>
	<tr style='text-align: left'>
	<th>#</th>
	<th>Id</th>
	<th>Login</th>
	<th>URL</th>
	</tr>
	{{range .}}
	<tr>
	<td>{{.Id}}</td>
	<td>{{.Login}}</td>
	<td><a href='{{.HtmlUrl}}'>{{.HtmlUrl}}</a></td>
	</tr>
	{{end}}
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

	report.Execute(w, users)
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 2)
}
