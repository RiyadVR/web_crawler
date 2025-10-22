package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) < 3 {
		fmt.Println("not enough arguments provided")
		fmt.Println("usage: crawler <baseURL> <maxConcurrency> <maxPages>")
		os.Exit(1)
	}

	if len(args) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	rawBaseURL := args[0]
	maxConcurrency, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("second argument (max concurrency) should be a number ")
		os.Exit(1)
	}
	maxPages, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("third argument (max pages) should be a number ")
		os.Exit(1)
	}

	cfg, err := configure(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
		return
	}

	fmt.Printf("starting crawl of: %s\n", rawBaseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	for normalizedURL, count := range cfg.pages {
		fmt.Printf("%v - %s\n", count.URL, normalizedURL)
	}
}
