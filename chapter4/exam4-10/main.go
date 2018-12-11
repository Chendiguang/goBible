package main

import (
	"fmt"
	"goBible/chapter4/github"
	"log"
	"os"
	"sort"
	"time"
)

func SortResult(result *github.IssuesSearchResult) {
	sort.Slice(result.Items, func(i, j int) bool {
		return result.Items[i].CreatedAt.Before(result.Items[j].CreatedAt)
	})
}

// Issues prints a table of GitHub issues matching the search terms.
func main() {
	result, err := github.SearchIssues(os.Args[3:])
	if err != nil {
		log.Fatal(err)
	}
	SortResult(result)
	var month, year, distant []*github.Issue
	for v, i := range result.Items {
		if i.CreatedAt.After(time.Now().Add(time.Hour * 24 * 30 * -1)) {
			month = append(month, result.Items[v])
		} else if i.CreatedAt.After(time.Now().Add(time.Hour * 24 * 365 * -1)) {
			year = append(year, result.Items[v])
		} else {
			distant = append(distant, result.Items[v])
		}
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	fmt.Printf("%d issues in Month\n", len(month))
	for _, item := range month {
		fmt.Printf("#%-5d %s %9.9s %.55s\n",
			item.Number, item.CreatedAt, item.User.Login, item.Title)
	}

	fmt.Printf("%d issues in Year\n", len(month))
	for _, item := range year {
		fmt.Printf("#%-5d %s %9.9s %.55s\n",
			item.Number, item.CreatedAt, item.User.Login, item.Title)
	}

	fmt.Printf("%d issues Over Year\n", len(month))
	for _, item := range distant {
		fmt.Printf("#%-5d %s %9.9s %.55s\n",
			item.Number, item.CreatedAt, item.User.Login, item.Title)
	}

}
