package json

import (
	"fmt"
	"log"
	"os"
	"time"
)

func RunSearch() {
	result, err := searchGitHubIssues(os.Args[2:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)

	now := time.Now()

	var recentItems []*Issue
	var laterItems []*Issue
	var latestItems []*Issue

	oneMonthAgo := now.AddDate(0, -1, 0)
	oneYearAgo := now.AddDate(-1, 0, 0)

	for _, item := range result.Items {
		if item.CreatedAt.After(oneMonthAgo) {
			recentItems = append(recentItems, item)
		} else if item.CreatedAt.Before(oneMonthAgo) && item.CreatedAt.After(oneYearAgo) {
			laterItems = append(laterItems, item)
		} else {
			latestItems = append(latestItems, item)
		}
	}

	fmt.Println("---------------------------------------------------------------------------")
	fmt.Println("In the last month:")
	fmt.Println("---------------------------------------------------------------------------")
	printRange(recentItems)
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Println("In the last year:")
	fmt.Println("---------------------------------------------------------------------------")
	printRange(laterItems)
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Println("Everything else:")
	fmt.Println("---------------------------------------------------------------------------")
	printRange(latestItems)
}

func RunGitHubCreate(headerParams IssueHeaderParams, pathParams IssuePathParams, bodyParams IssueBodyParams) {
	editedBodyParams := editBodyParams(bodyParams)
	err := createGitHubIssue(headerParams, pathParams, editedBodyParams)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating issue: %v\n", err)
		os.Exit(1)
	}
}

func RunGitHubRead(headerParams IssueHeaderParams, pathParams IssuePathParams, bodyParams IssueBodyParams) {
	err := readGitHubIssue(headerParams, pathParams, bodyParams)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading issue: %v\n", err)
		os.Exit(1)
	}
}

func RunGitHubUpdate(headerParams IssueHeaderParams, pathParams IssuePathParams, bodyParams IssueBodyParams) {
	editedBodyParams := editBodyParams(bodyParams)
	err := updateGitHubIssue(headerParams, pathParams, editedBodyParams)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error updating issue: %v\n", err)
		os.Exit(1)
	}
}

func RunGitHubLock(headerParams IssueHeaderParams, pathParams IssuePathParams, bodyParams IssueBodyParams) {
	err := lockGitHubIssue(headerParams, pathParams, bodyParams)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error locking issue: %v\n", err)
		os.Exit(1)
	}
}

// Helper functions
func printRange(items []*Issue) {
	for _, item := range items {
		fmt.Printf("#%-5d %9.9s %.55s %v\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
}
