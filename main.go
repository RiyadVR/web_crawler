package main

import (
	"fmt"
	"os"
)

func main() {
	baseURL := os.Args[1:]
	if len(baseURL) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	if len(baseURL) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	fmt.Printf("starting crawl of: %s", baseURL[0])

}
