package json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const SearchIssuesURL = "https://api.github.com/search/issues"
const CreateIssueURL = "https://api.github.com/repos/%s/%s/issues"
const ReadIssueURL = "https://api.github.com/repos/%s/%s/issues/%s"
const UpdateIssueURL = "https://api.github.com/repos/%s/%s/issues/%s"
const LockIssueURL = "https://api.github.com/repos/%s/%s/issues/%s/lock"
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
	createURL := fmt.Sprintf(CreateIssueURL, pathParams.Owner, pathParams.Repo)

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
	readURL := fmt.Sprintf(ReadIssueURL, pathParams.Owner, pathParams.Repo, pathParams.IssueNumber)

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
	updateURL := fmt.Sprintf(UpdateIssueURL, pathParams.Owner, pathParams.Repo, pathParams.IssueNumber)

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
	lockURL := fmt.Sprintf(LockIssueURL, pathParams.Owner, pathParams.Repo, pathParams.IssueNumber)

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

// These functions are for the webserver ex4.14

var issueSearchURL = "https://api.github.com/repos/%s/%s/issues"         //owner & repo
var milestoneSearchURL = "https://api.github.com/repos/%s/%s/milestones" //owner & repo
var getUserURL = " https://api.github.com/users/%s"                      // username

type GitHubIssue struct {
	Id            int    `json:"id"`
	NodeId        string `json:"node_id"`
	Url           string `json:"url"`
	RepositoryUrl string `json:"repository_url"`
	LabelsUrl     string `json:"labels_url"`
	CommentsUrl   string `json:"comments_url"`
	EventsUrl     string `json:"events_url"`
	HtmlUrl       string `json:"html_url"`
	Number        int    `json:"number"`
	State         string `json:"state"`
	Title         string `json:"title"`
	Body          string `json:"body"`
	User          struct {
		Login             string `json:"login"`
		Id                int    `json:"id"`
		NodeId            string `json:"node_id"`
		AvatarUrl         string `json:"avatar_url"`
		GravatarId        string `json:"gravatar_id"`
		Url               string `json:"url"`
		HtmlUrl           string `json:"html_url"`
		FollowersUrl      string `json:"followers_url"`
		FollowingUrl      string `json:"following_url"`
		GistsUrl          string `json:"gists_url"`
		StarredUrl        string `json:"starred_url"`
		SubscriptionsUrl  string `json:"subscriptions_url"`
		OrganizationsUrl  string `json:"organizations_url"`
		ReposUrl          string `json:"repos_url"`
		EventsUrl         string `json:"events_url"`
		ReceivedEventsUrl string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"user"`
	Labels []struct {
		Id          int    `json:"id"`
		NodeId      string `json:"node_id"`
		Url         string `json:"url"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Color       string `json:"color"`
		Default     bool   `json:"default"`
	} `json:"labels"`
	Assignee struct {
		Login             string `json:"login"`
		Id                int    `json:"id"`
		NodeId            string `json:"node_id"`
		AvatarUrl         string `json:"avatar_url"`
		GravatarId        string `json:"gravatar_id"`
		Url               string `json:"url"`
		HtmlUrl           string `json:"html_url"`
		FollowersUrl      string `json:"followers_url"`
		FollowingUrl      string `json:"following_url"`
		GistsUrl          string `json:"gists_url"`
		StarredUrl        string `json:"starred_url"`
		SubscriptionsUrl  string `json:"subscriptions_url"`
		OrganizationsUrl  string `json:"organizations_url"`
		ReposUrl          string `json:"repos_url"`
		EventsUrl         string `json:"events_url"`
		ReceivedEventsUrl string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"assignee"`
	Assignees []struct {
		Login             string `json:"login"`
		Id                int    `json:"id"`
		NodeId            string `json:"node_id"`
		AvatarUrl         string `json:"avatar_url"`
		GravatarId        string `json:"gravatar_id"`
		Url               string `json:"url"`
		HtmlUrl           string `json:"html_url"`
		FollowersUrl      string `json:"followers_url"`
		FollowingUrl      string `json:"following_url"`
		GistsUrl          string `json:"gists_url"`
		StarredUrl        string `json:"starred_url"`
		SubscriptionsUrl  string `json:"subscriptions_url"`
		OrganizationsUrl  string `json:"organizations_url"`
		ReposUrl          string `json:"repos_url"`
		EventsUrl         string `json:"events_url"`
		ReceivedEventsUrl string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"assignees"`
	Milestone struct {
		Url         string `json:"url"`
		HtmlUrl     string `json:"html_url"`
		LabelsUrl   string `json:"labels_url"`
		Id          int    `json:"id"`
		NodeId      string `json:"node_id"`
		Number      int    `json:"number"`
		State       string `json:"state"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Creator     struct {
			Login             string `json:"login"`
			Id                int    `json:"id"`
			NodeId            string `json:"node_id"`
			AvatarUrl         string `json:"avatar_url"`
			GravatarId        string `json:"gravatar_id"`
			Url               string `json:"url"`
			HtmlUrl           string `json:"html_url"`
			FollowersUrl      string `json:"followers_url"`
			FollowingUrl      string `json:"following_url"`
			GistsUrl          string `json:"gists_url"`
			StarredUrl        string `json:"starred_url"`
			SubscriptionsUrl  string `json:"subscriptions_url"`
			OrganizationsUrl  string `json:"organizations_url"`
			ReposUrl          string `json:"repos_url"`
			EventsUrl         string `json:"events_url"`
			ReceivedEventsUrl string `json:"received_events_url"`
			Type              string `json:"type"`
			SiteAdmin         bool   `json:"site_admin"`
		} `json:"creator"`
		OpenIssues   int       `json:"open_issues"`
		ClosedIssues int       `json:"closed_issues"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
		ClosedAt     time.Time `json:"closed_at"`
		DueOn        time.Time `json:"due_on"`
	} `json:"milestone"`
	Locked           bool   `json:"locked"`
	ActiveLockReason string `json:"active_lock_reason"`
	Comments         int    `json:"comments"`
	PullRequest      struct {
		Url      string `json:"url"`
		HtmlUrl  string `json:"html_url"`
		DiffUrl  string `json:"diff_url"`
		PatchUrl string `json:"patch_url"`
	} `json:"pull_request"`
	ClosedAt  interface{} `json:"closed_at"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	ClosedBy  struct {
		Login             string `json:"login"`
		Id                int    `json:"id"`
		NodeId            string `json:"node_id"`
		AvatarUrl         string `json:"avatar_url"`
		GravatarId        string `json:"gravatar_id"`
		Url               string `json:"url"`
		HtmlUrl           string `json:"html_url"`
		FollowersUrl      string `json:"followers_url"`
		FollowingUrl      string `json:"following_url"`
		GistsUrl          string `json:"gists_url"`
		StarredUrl        string `json:"starred_url"`
		SubscriptionsUrl  string `json:"subscriptions_url"`
		OrganizationsUrl  string `json:"organizations_url"`
		ReposUrl          string `json:"repos_url"`
		EventsUrl         string `json:"events_url"`
		ReceivedEventsUrl string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"closed_by"`
	AuthorAssociation string `json:"author_association"`
	StateReason       string `json:"state_reason"`
}

func getBugReports(project string, repo string) (returnIssues []GitHubIssue) {
	client := resty.New()

	resp, _ := client.R().
		EnableTrace().
		Get(fmt.Sprintf(issueSearchURL, project, repo))

	json.Unmarshal(resp.Body(), &returnIssues)

	return returnIssues
}

type GitHubMilestone struct {
	Url         string `json:"url"`
	HtmlUrl     string `json:"html_url"`
	LabelsUrl   string `json:"labels_url"`
	Id          int    `json:"id"`
	NodeId      string `json:"node_id"`
	Number      int    `json:"number"`
	State       string `json:"state"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Creator     struct {
		Login             string `json:"login"`
		Id                int    `json:"id"`
		NodeId            string `json:"node_id"`
		AvatarUrl         string `json:"avatar_url"`
		GravatarId        string `json:"gravatar_id"`
		Url               string `json:"url"`
		HtmlUrl           string `json:"html_url"`
		FollowersUrl      string `json:"followers_url"`
		FollowingUrl      string `json:"following_url"`
		GistsUrl          string `json:"gists_url"`
		StarredUrl        string `json:"starred_url"`
		SubscriptionsUrl  string `json:"subscriptions_url"`
		OrganizationsUrl  string `json:"organizations_url"`
		ReposUrl          string `json:"repos_url"`
		EventsUrl         string `json:"events_url"`
		ReceivedEventsUrl string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"creator"`
	OpenIssues   int       `json:"open_issues"`
	ClosedIssues int       `json:"closed_issues"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	ClosedAt     time.Time `json:"closed_at"`
	DueOn        time.Time `json:"due_on"`
}

func getMilestones(project string, repo string) (returnMilestones []GitHubMilestone) {
	fmt.Println("Got milestones")

	client := resty.New()

	resp, _ := client.R().
		EnableTrace().
		Get(fmt.Sprintf(milestoneSearchURL, project, repo))

	json.Unmarshal(resp.Body(), &returnMilestones)

	return returnMilestones
}

type GitHubUser struct {
	Login             string `json:"login"`
	Id                int    `json:"id"`
	NodeId            string `json:"node_id"`
	AvatarUrl         string `json:"avatar_url"`
	GravatarId        string `json:"gravatar_id"`
	Url               string `json:"url"`
	HtmlUrl           string `json:"html_url"`
	FollowersUrl      string `json:"followers_url"`
	FollowingUrl      string `json:"following_url"`
	GistsUrl          string `json:"gists_url"`
	StarredUrl        string `json:"starred_url"`
	SubscriptionsUrl  string `json:"subscriptions_url"`
	OrganizationsUrl  string `json:"organizations_url"`
	ReposUrl          string `json:"repos_url"`
	EventsUrl         string `json:"events_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

func getUsers(project string, repo string) {
	fmt.Println("Got users")
}

type GitHubBundle struct {
	Issues     []GitHubIssue
	Users      []GitHubUser
	Milestones []GitHubMilestone
}
