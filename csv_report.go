package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func writeCSVReport(pages map[string]PageData, filename string) error {
	if len(pages) == 0 {
		fmt.Println("No data to write to CSV")
		return nil
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("create csv: %w", err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{
		"page_url",
		"h1",
		"first_paragraph",
		"outgoing_link_urls",
		"image_urls",
	})
	if err != nil {
		return fmt.Errorf("write header: %w", err)
	}

	for _, pageData := range pages { // iterating maps in go generates random output, order is not guaranteed, so better to sort the keys first - will implement it later
		err := writer.Write([]string{
			pageData.URL,
			pageData.H1,
			pageData.FirstParagraph,
			strings.Join(pageData.OutgoingLinks, ";"),
			strings.Join(pageData.ImageURLs, ";"),
		})
		if err != nil {
			return fmt.Errorf("write row for %s: %w", pageData.URL, err)
		}
	}

	fmt.Printf("Report written to %s\n", filename)

	return nil
}
