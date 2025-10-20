package main

import (
	"log"
	"net/url"
)

type PageData struct {
	URL            string
	H1             string
	FirstParagraph string
	OutgoingLinks  []string
	ImageURLs      []string
}

func extractPageData(html, pageURL string) PageData {
	baseURL, err := url.Parse(pageURL)
	if err != nil {
		log.Fatalf("parsing URL error: %v", err)
	}

	h1 := getH1FromHTML(html)
	firstParagraph := getFirstParagraphFromHTML(html)
	outGoingLinks, err := getURLsFromHTML(html, baseURL)
	if err != nil {
		log.Fatalf("extracting URLs error: %v", err)
	}
	imageURLs, err := getImagesFromHTML(html, baseURL)
	if err != nil {
		log.Fatalf("extracting images error: %v", err)
	}
	return PageData{
		URL:            pageURL,
		H1:             h1,
		FirstParagraph: firstParagraph,
		OutgoingLinks:  outGoingLinks,
		ImageURLs:      imageURLs,
	}
}
