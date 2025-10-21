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
	h1 := getH1FromHTML(html)
	firstParagraph := getFirstParagraphFromHTML(html)

	baseURL, err := url.Parse(pageURL)
	if err != nil {
		return PageData{
			URL:            pageURL,
			H1:             h1,
			FirstParagraph: firstParagraph,
			OutgoingLinks:  nil,
			ImageURLs:      nil,
		}
	}

	outGoingLinks, err := getURLsFromHTML(html, baseURL)
	if err != nil {
		log.Printf("extracting URLs error: %v", err)
	}
	imageURLs, err := getImagesFromHTML(html, baseURL)
	if err != nil {
		log.Printf("extracting images error: %v", err)
	}
	return PageData{
		URL:            pageURL,
		H1:             h1,
		FirstParagraph: firstParagraph,
		OutgoingLinks:  outGoingLinks,
		ImageURLs:      imageURLs,
	}
}
