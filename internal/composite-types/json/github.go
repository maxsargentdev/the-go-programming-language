package json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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
const GitHubApiVersion = "2022-11-28"

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

// searchGitHubIssues queries the GitHub issue tracker.
func searchGitHubIssues(terms []string) (*IssuesSearchResult, error) {
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

type IssueHeaderParams struct {
	Bearer string
}

type IssuePathParams struct {
	Owner       string
	Repo        string
	IssueNumber string
}

type IssueBodyParams struct {
	Title       string   `json:"title,omitempty"`
	Body        string   `json:"body,omitempty"`
	Assignee    string   `json:"assignee,omitempty"`
	State       string   `json:"state,omitempty"`
	StateReason string   `json:"state_reason,omitempty"`
	Milestone   string   `json:"milestone,omitempty"`
	Labels      []string `json:"labels,omitempty"`
	Assignees   []string `json:"assignees,omitempty"`
	LockReason  string   `json:"lock_reason,omitempty"`
}

type readIssueBodyParams struct {
}

// createGitHubIssue creates a new GitHub issue
func createGitHubIssue(headerParams IssueHeaderParams, pathParams IssuePathParams, bodyParams IssueBodyParams) error {

	// Interpolate our OWNER and REPO values into the URL path
	createURL := strings.Replace(CreateIssueURL, "OWNER", pathParams.Owner, 1)
	createURL = strings.Replace(createURL, "REPO", pathParams.Repo, 1)

	// Marshal body for the wire
	postBody, err := json.Marshal(bodyParams)
	if err != nil {
		return err
	}
	postBodyBytes := bytes.NewReader(postBody)

	// Create client for POST request
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, createURL, postBodyBytes)
	if err != nil {
		return err
	}

	// Add bearer
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", headerParams.Bearer))
	req.Header.Add("Accept", GitHubContentType)
	req.Header.Add("X-GitHub-Api-Version", GitHubApiVersion)

	// Fire POST request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	// No 200 or fail to close then return the error
	if resp.StatusCode != http.StatusCreated {
		err := resp.Body.Close()
		if err != nil {
			return err
		}
		return fmt.Errorf("issue creation failed: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error while reading the response bytes:", err)
	}

	// For debug
	fmt.Println(body)
	return nil
}

// readGitHubIssue reads an existing GitHub issue
func readGitHubIssue(headerParams IssueHeaderParams, pathParams IssuePathParams, bodyParams IssueBodyParams) error {
	// Interpolate our OWNER and REPO values into the URL path
	readURL := strings.Replace(ReadIssueURL, "OWNER", pathParams.Owner, 1)
	readURL = strings.Replace(readURL, "REPO", pathParams.Repo, 1)
	readURL = strings.Replace(readURL, "ISSUE_NUMBER", pathParams.IssueNumber, 1)

	// Marshal body for the wire
	getBody, err := json.Marshal(bodyParams)
	if err != nil {
		return err
	}
	getBodyBytes := bytes.NewReader(getBody)

	// Create client for GET request
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, readURL, getBodyBytes)
	if err != nil {
		return err
	}

	// Add bearer
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", headerParams.Bearer))
	req.Header.Add("Accept", GitHubContentType)
	req.Header.Add("X-GitHub-Api-Version", GitHubApiVersion)

	// Fire GET request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	// No 200 or fail to close then return the error
	if resp.StatusCode != http.StatusOK {
		err := resp.Body.Close()
		if err != nil {
			return err
		}
		return fmt.Errorf("issue read failed: %s", resp.Status)
	}

	// For debug
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// No need to deserialize the response, just display a string
	fmt.Printf(string(respBody))
	defer resp.Body.Close()
	return nil
}

// updateGitHubIssue updates an existing issue
func updateGitHubIssue(headerParams IssueHeaderParams, pathParams IssuePathParams, bodyParams IssueBodyParams) error {
	// Interpolate our OWNER and REPO values into the URL path
	updateURL := strings.Replace(UpdateIssueURL, "OWNER", pathParams.Owner, 1)
	updateURL = strings.Replace(updateURL, "REPO", pathParams.Repo, 1)
	updateURL = strings.Replace(updateURL, "ISSUE_NUMBER", pathParams.IssueNumber, 1)

	// Marshal body for the wire
	patchBody, err := json.Marshal(bodyParams)
	if err != nil {
		return err
	}
	patchBodyBytes := bytes.NewReader(patchBody)

	// Create client for PATCH request
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPatch, updateURL, patchBodyBytes)
	if err != nil {
		return err
	}

	// Add bearer
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", headerParams.Bearer))
	req.Header.Add("Accept", GitHubContentType)
	req.Header.Add("X-GitHub-Api-Version", GitHubApiVersion)

	// Fire PATCH request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	// No 200 or fail to close then return the error
	if resp.StatusCode != http.StatusOK {
		err := resp.Body.Close()
		if err != nil {
			return err
		}
		return fmt.Errorf("issue update failed: %s", resp.Status)
	}

	// For debug
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// No need to deserialize the response, just display a string
	fmt.Printf(string(respBody))
	defer resp.Body.Close()
	return nil
}

// lockGitHubIssue locks an issue, instead of deleting
func lockGitHubIssue(headerParams IssueHeaderParams, pathParams IssuePathParams, bodyParams IssueBodyParams) error {
	// Interpolate our OWNER and REPO values into the URL path
	lockURL := strings.Replace(LockIssueURL, "OWNER", pathParams.Owner, 1)
	lockURL = strings.Replace(lockURL, "REPO", pathParams.Repo, 1)
	lockURL = strings.Replace(lockURL, "ISSUE_NUMBER", pathParams.IssueNumber, 1)

	// Marshal body for the wire
	postBody, err := json.Marshal(bodyParams)
	if err != nil {
		return err
	}
	patchBodyBytes := bytes.NewReader(postBody)

	// Create client for PATCH request
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, lockURL, patchBodyBytes)
	if err != nil {
		return err
	}

	// Add bearer
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", headerParams.Bearer))
	req.Header.Add("Accept", GitHubContentType)
	req.Header.Add("X-GitHub-Api-Version", GitHubApiVersion)

	// Fire PATCH request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	// No 200 or fail to close then return the error
	if resp.StatusCode != http.StatusNoContent {
		err := resp.Body.Close()
		if err != nil {
			return err
		}
		return fmt.Errorf("issue update failed: %s", resp.Status)
	}

	// For debug
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// No need to deserialize the response, just display a string
	fmt.Printf(string(respBody))
	defer resp.Body.Close()
	return nil
}
