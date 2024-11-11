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

	url := fmt.Sprintf("https://api.github.com/users/%s/events", os.Args[1])
	body, _, err := http_handler.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch %v\n", err)
	}

	activities, err := json_handler.ParseActivities(body)
	if err != nil {
		log.Fatalf("Failed to decode json %v\n", err)
	}

	fmt.Println(activities)

}
