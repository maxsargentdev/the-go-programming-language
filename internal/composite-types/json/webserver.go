package json

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

func Serve(project string, repo string) {

	fmt.Println("Server Starting")

	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/bugreports", bugReportHandler)
	http.HandleFunc("/milestones", milestonesHandler)
	http.HandleFunc("/users", usersHandler)

	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	const templ = `
	<h1>Home</h1>
	<ul>
	  <li>Bugreports</li>
	  <li>Milestones</li>
	  <li>Users</li>
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

func bugReportHandler(w http.ResponseWriter, r *http.Request) {
	const templ = `
	<h1>Bugreports</h1>`

	report, err := template.New("bugreports").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ)
	if err != nil {
		log.Fatal(err)
	}

	report.Execute(w, IssueBodyParams{})
}

func milestonesHandler(w http.ResponseWriter, r *http.Request) {
	const templ = `
	<h1>Milestones</h1>`

	report, err := template.New("milestones").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ)
	if err != nil {
		log.Fatal(err)
	}

	report.Execute(w, IssueBodyParams{})
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	const templ = `
	<h1>Users</h1>`

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
