# Quote Scraper API


This project is a web scraper built in Go (Golang) using the Colly framework for extracting quotes from the website [http://quotes.toscrape.com](http://quotes.toscrape.com). The scraped data is then stored in a PostgreSQL database using the GORM ORM. In addition to web scraping, the project serves an API with Gin to access quotes and tags easily.


## Features
  -   Grabs quotes, tags and authors from [http://quotes.toscrape.com](http://quotes.toscrape.com).
   -   Sets up a slick API using the Gin framework for easy access to your quote collection. 
   -   Includes swagger documentation so you can effortlessly explore and interact with the API.


## Getting Started

### Prerequisites

To run this authentication service, ensure you have the following prerequisites installed:

- Go programming language
- PostgreSQL database 

### Installation

1. Clone the repository:

```bash
   git clone https://github.com/your-username/authentication-service.git
   cd authentication-service
```
2. Install dependencies:
```go
   go install ./...
```
3. Configuration

Configure the service by setting up the PostgreSQL connection details by editing the DSN in the main.go file.


## API Documentation<a id="api-doc"></a>

The API is thoroughly documented using Swagger annotations. Access the Swagger UI at [http://localhost:8080/docs/index.html](http://localhost:8080/docs/index.html) to explore and interact with the API.