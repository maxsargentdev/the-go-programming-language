package json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const SearchIssuesURL = "https://api.github.com/search/issues"
const CreateIssueURL = "https://api.github.com/repos/OWNER/REPO/issues"
const ReadIssueURL = "https://api.github.com/repos/OWNER/REPO/issues/ISSUE_NUMBER"
const UpdateIssueURL = "https://api.github.com/repos/OWNER/REPO/issues/ISSUE_NUMBER"
const LockIssueURL = "https://api.github.com/repos/OWNER/REPO/issues/ISSUE_NUMBER/lock"
const GitHubContentType = "application/vnd.github+json"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// saerchIssuess queries the GitHub issue tracker.
func searchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(SearchIssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

type createIssuePathParams struct {
	Owner string
	Repo  string
}

type createIssueBodyParams struct {
	Title     string   `json:"title,omitempty"`
	Body      string   `json:"body,omitempty"`
	Assignee  string   `json:"assignee,omitempty"`
	Milestone string   `json:"milestone,omitempty"`
	Labels    []string `json:"labels,omitempty"`
	Assignees []string `json:"assignees,omitempty"`
}

// createIssue creates a new GitHub issue
func createIssue(pathParams createIssuePathParams, bodyParams createIssueBodyParams) error {

	// Interpolate our OWNER and REPO values into the URL path
	createURL := strings.Replace(CreateIssueURL, "OWNER", pathParams.Owner, 1)
	createURL = strings.Replace(createURL, "REPO", pathParams.Repo, 1)

	// Marshal body for the wire
	postBody, err := json.Marshal(bodyParams.Body)
	if err != nil {
		return err
	}
	postBodyBytes := bytes.NewReader(postBody)

	// Fire POST request
	resp, err := http.Post(createURL, GitHubContentType, postBodyBytes)
	if err != nil {
		return err
	}

	// No 200 or fail to close then return the error
	if resp.StatusCode != http.StatusOK {
		err := resp.Body.Close()
		if err != nil {
			return err
		}
		return fmt.Errorf("search query failed: %s", resp.Status)
	}

	// For debug
	fmt.Println(resp.Body)
	return nil
}

// readIssue reads an existing GitHub issue
func readIssue() {
	fmt.Println("Issue read")
}

// updateIssue updates an existing issue
func updateIssue() {
	fmt.Println("Issue updated")
}

// lockIssue locks an issue, instead of deleting
func lockIssue() {
	fmt.Println("Issue locked")
}
