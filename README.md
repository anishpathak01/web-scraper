Website Scraper Implementation Using Go
---------------------------------------

This project implements a web scraper using the **Colly** framework in Go. The scraper extracts information from a website, processes the data, and stores it in a structured format.

### Features

-   **Data Extraction**: Uses the Colly framework to scrape relevant data from web pages.
-   **Customizable Scraping Logic**: Allows for easy modification of scraping logic based on the website's structure.
-   **Data Storage**: The scraped data is stored in **JSON files** for easy access and further processing.

### Technologies Used

-   **Go**: Primary programming language.
-   **Colly**: Web scraping framework for Go.

### How to Run

1.  Clone the repository.
2.  Install the required Go dependencies:

    bash

    Copy code

    `go get github.com/gocolly/colly`

3.  Run the scraper:

    bash

    Copy code

    `go run main.go`

### Notes

-   Ensure you have permission to scrape the website you are targeting.
-   The scraper handles basic error cases like failed HTTP requests.