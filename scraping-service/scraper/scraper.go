package scraper

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/AbderraoufKhorchani/web-scraper/scraping-service/data"
	"github.com/gocolly/colly"
)

type Scraper struct {
}

var quoteInstance data.Quote
var mu sync.Mutex

func (s *Scraper) Scrape() {
	// Lock the mutex before checking if the database is empty
	mu.Lock()
	defer mu.Unlock()

	// Check if the database is empty
	isEmpty, err := quoteInstance.DatabaseIsEmpty()
	if err != nil {
		log.Println("Error checking if the database is empty:", err)
		return
	}

	if isEmpty {
		// Create a new collector
		c := colly.NewCollector()

		// Set the base URL to scrape
		baseURL := "http://quotes.toscrape.com/page/"

		// Create a WaitGroup to wait for all goroutines to finish
		var wg sync.WaitGroup

		// Set up the callback for handling extracted data
		c.OnHTML("div.quote", func(e *colly.HTMLElement) {
			// Extract quote text
			quoteText := e.ChildText("span.text")

			// Extract author
			author := e.ChildText("small.author")

			// Extract tags
			var tags []string
			e.ForEach("div.tags a.tag", func(_ int, el *colly.HTMLElement) {
				tags = append(tags, el.Text)
			})

			quote := &data.Quote{
				Author:    author,
				QuoteText: quoteText,
			}

			// Add the quote to the database
			err := quoteInstance.AddQuoteWithTags(*quote, tags)
			if err != nil {
				log.Println("Error adding quote:", err)
			}

			// Print the extracted data
			fmt.Printf("Quote: %s\nAuthor: %s\nTags: %s\n\n", quoteText, author, strings.Join(tags, ", "))
		})

		// Set up error handling
		c.OnError(func(r *colly.Response, err error) {
			log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
		})

		// Start the scraping process with a goroutine for each page
		for i := 1; i <= 10; i++ {
			// Increment the WaitGroup counter for each goroutine
			wg.Add(1)

			go func(page int) {
				defer wg.Done()

				// Visit each page concurrently
				err := c.Visit(fmt.Sprintf("%s%d/", baseURL, page))
				if err != nil {
					log.Println("Error visiting page:", page, "-", err)
				}
			}(i)
		}

		// Wait for all goroutines to finish
		wg.Wait()
	}
}
