package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

type book struct {
	Title        string `json:"name"`
	Price        string `json:"price"`
	Availability bool   `json:"availability"`
}

func main() {
	scraper := colly.NewCollector(
		colly.AllowedDomains("books.toscrape.com"),
	)

	var Books []book

	scraper.OnHTML(".product_pod", func(e *colly.HTMLElement) {
		title := e.ChildAttr("h3 a", "title")
		price := e.ChildText(".price_color")
		stock := e.ChildText(".instock.availability")
		availability := stock == " In stock "

		currentBook := book{
			Title:        title,
			Price:        price,
			Availability: availability,
		}

		Books = append(Books, currentBook)
	})

	scraper.OnHTML(".next a", func(e *colly.HTMLElement) {
		nextPage := e.Attr("href")
		if nextPage != "" {
			// Ensure the nextPage URL includes the "catalogue" route if it's missing
			if !strings.HasPrefix(nextPage, "catalogue") {
				nextPage = fmt.Sprintf("catalogue/%s", nextPage)
			}
			err := scraper.Visit("https://books.toscrape.com/" + nextPage)
			if err != nil {
				fmt.Println("Failed to scrap next page: ", "books.toscrape.com/"+nextPage, err)
			}
		}
	})

	scraper.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request failed:", r.Request.URL, err)
	})

	err := scraper.Visit("https://books.toscrape.com")
	if err != nil {
		fmt.Println("Failed to scrap the website: ", "books.toscrape.com", err)
	}

	data, err := json.MarshalIndent(Books, "", " ")
	if err != nil {
		fmt.Println("Failed to marshal the data: ", err)
	}

	err = os.WriteFile("books.json", data, 0644)
	if err != nil {
		fmt.Println("Failed to write data to the file books.json")
	}
}
