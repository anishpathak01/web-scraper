package main

import (
	"encoding/json"
	"fmt"
	"os"

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
