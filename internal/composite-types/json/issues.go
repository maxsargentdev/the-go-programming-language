package json

import (
	"fmt"
	"log"
	"os"
	"time"
)

func RunSearch() {
	result, err := searchIssues(os.Args[2:])
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

func RunCreate() {
	pathParams := issuePathParams{}
	bodyParams := createIssueBodyParams{}
	createIssue(pathParams, bodyParams)
}

func RunRead() {
	readIssue()
}

func RunUpdate() {
	updateIssue()
}

func RunLock() {
	lockIssue()
}

// Helper functions
func printRange(items []*Issue) {
	for _, item := range items {
		fmt.Printf("#%-5d %9.9s %.55s %v\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
}
