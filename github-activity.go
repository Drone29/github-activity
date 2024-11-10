package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments!")
		fmt.Println(os.Args[0], "<username>")
		return
	}

	// url := fmt.Sprintf("https://api.github.com/users/%s/events", os.Args[1])

}
