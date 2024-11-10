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

}
