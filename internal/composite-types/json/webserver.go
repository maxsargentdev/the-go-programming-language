package json

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

func Serve(project string, repo string) {
	getBugReports(project, repo)
	getMilestones(project, repo)
	getUsers(project, repo)

	fmt.Println("Server Starting")
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/index", HttpFileHandler)

	//http.HandleFunc("/", indexTemplateHandler)

	http.ListenAndServe(":8080", nil)

	renderHome()
}

func renderHome() {
	fmt.Println("Rendering home HTML.........")
}

func renderBugReports() {
	fmt.Println("Rendering HTML templates.......")
}

func renderMilestones() {
	fmt.Println("Rendering HTML templates.......")
}

func renderUsers() {
	fmt.Println("Rendering HTML templates.......")
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

	report, err := template.New("report").
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

func HttpFileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s!", r.URL.Path[1:])
	//http.ServeFile(response, request, "Index.html")
}
