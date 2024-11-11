package main

import (
	"fmt"
	"github-activity/http_handler"
	"github-activity/json_handler"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments!")
		fmt.Println(os.Args[0], "<username>")
		return
	}
	username := os.Args[1]
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	body, _, err := http_handler.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch %v\n", err)
	}

	activities, err := json_handler.ParseActivities(body)
	if err != nil {
		log.Fatalf("Failed to decode json %v\n", err)
	}

	fmt.Printf("%s activity:\n", username)
	for _, activity := range activities.Objects {
		switch activity.Type {
		case "PushEvent":
			fmt.Printf("- %s Pushed %d commits to %s\n", activity.CreatedAt, len(activity.Payload.Commits), activity.Repo.Name)
		case "CreateEvent":
			fmt.Printf("- %s Created %s %s at %s\n", activity.CreatedAt, activity.Payload.RefType, activity.Payload.Ref, activity.Repo.Name)
		case "PullRequestEvent":
			fmt.Printf("- %s Updated PR status to %s (%s)\n", activity.CreatedAt, activity.Repo.Name, activity.Payload.Action)
		case "PullRequestReviewCommentEvent":
			fmt.Printf("- %s Updated PR review comment status to %s (%s)\n", activity.CreatedAt, activity.Repo.Name, activity.Payload.Action)
		case "PullRequestReviewEvent":
			fmt.Printf("- %s Updated PR review status to %s (%s)\n", activity.CreatedAt, activity.Repo.Name, activity.Payload.Action)
		case "IssuesEvent":
			fmt.Printf("- %s Updated issue status at %s (%s)\n", activity.CreatedAt, activity.Repo.Name, activity.Payload.Issue.State)
		case "DeleteEvent":
			fmt.Printf("- %s Deleted %s %s at %s\n", activity.CreatedAt, activity.Payload.RefType, activity.Payload.Ref, activity.Repo.Name)
		}
	}

}
