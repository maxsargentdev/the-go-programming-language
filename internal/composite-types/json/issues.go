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
	recentItems := []*Issue
	laterItems := []*Issue
	latestItems := []*Issue
	oneMonthFrom := now.AddDate(0, 1, 0)
	oneYearFrom := now.AddDate(1, 0, 0)
	for _, item := range result.Items {
		//fmt.Printf("#%-5d %9.9s %.55s %v\n",
		//	item.Number, item.User.Login, item.Title, item.CreatedAt)
		if item.CreatedAt < oneMonthFrom {
			recentItems = append(recentItems, item)
		} else if oneYearFrom > item.CreatedAt > oneMonthFrom {
			laterItems = append(laterItems, item)
		} else {
			latestItems = append(latestItems, item)
		}
	}
}
