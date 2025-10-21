package main

import (
	"fmt"
	"log"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	parsedRawBaseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawBaseURL, err)
		return
	}
	parsedRawCurrentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}

	if parsedRawBaseURL.Host != parsedRawCurrentURL.Host {
		fmt.Println("not same domain")
		return
	}

	normalizeRawCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		log.Printf("error normalizing the current url: %v", err)
		return
	}

	_, ok := pages[normalizeRawCurrentURL]
	if ok {
		pages[normalizeRawCurrentURL]++
		return
	}

	pages[normalizeRawCurrentURL] = 1

	fmt.Printf("Crawling: %s\n", rawCurrentURL)
	html, err := getHTML(rawCurrentURL)
	if err != nil {
		log.Printf("error getting the html: %v, url: %s", err, rawCurrentURL)
		return
	}

	pageData := extractPageData(html, rawCurrentURL)
	for _, link := range pageData.OutgoingLinks {
		crawlPage(rawBaseURL, link, pages)
	}
}
