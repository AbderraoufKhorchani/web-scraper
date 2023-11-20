package main

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

func main() {
	// Create a new collector
	c := colly.NewCollector()

	// Set the URL to scrape
	url := "http://quotes.toscrape.com"

	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Define the CSS selector for quotes
	quoteSelector := "div.quote"

	// Set up the callback for handling extracted data
	c.OnHTML(quoteSelector, func(e *colly.HTMLElement) {
		// Decrement the WaitGroup counter when the goroutine completes
		defer wg.Done()

		// Extract quote text
		quoteText := e.ChildText("span.text")

		// Extract author
		author := e.ChildText("small.author")

		// Extract tags using strings.Fields
		tags := strings.Fields(e.ChildText("div.tags"))
		if len(tags) > 0 {
			tags = tags[1:]
		}

		// Print the extracted data
		fmt.Printf("Quote: %s\nAuthor: %s\nTags: %s\n\n", quoteText, author, strings.Join(tags, ", "))
	})

	// Set up error handling
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Start the scraping process with a goroutine for each page
	for i := 1; i <= 5; i++ {
		// Increment the WaitGroup counter for each goroutine
		wg.Add(1)

		go func(page int) {
			defer wg.Done()

			// Visit each page concurrently
			err := c.Visit(fmt.Sprintf("%s/page/%d/", url, page))
			if err != nil {
				log.Println("Error visiting page:", page, "-", err)
			}
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()
}
