## Table of Contents

-   [Overview](#overview)
-   [Features](#features)
-   [Installation](#installation)
-   [Usage](#usage)

## Overview

This project is a web scraper built in Go (Golang) using the [Colly](https://github.com/gocolly/colly) framework for extracting quotes from the website [http://quotes.toscrape.com](http://quotes.toscrape.com/). The scraped data is then stored in a PostgreSQL database using the [GORM](https://gorm.io/) ORM.

## Features

-   **Concurrent Scraping:** The scraper uses goroutines to concurrently scrape multiple pages of quotes, improving efficiency.
    
-   **Data Storage:** Quotes, authors, and tags are stored in a PostgreSQL database, providing a persistent storage solution.
    
-   **Database Interaction:** The project demonstrates the use of the GORM ORM for database interactions, including querying, adding, and checking for database emptiness.

## Installation
To get started with the project, follow the installation steps below.
### Prerequisites

Before installing the project, ensure that you have the following prerequisites installed:

- **Go Version 1.19:** 

You need Go version 1.19 or later. Check your Go version using the following command:

```bash
go version` 
```  
If Go is not installed or you have an older version, you can download and install it from the official Go website: [https://golang.org/dl/](https://golang.org/dl/)


-   **Make:** 

Make is used for building and managing the project. Verify if Make is installed with:
    
```bash
	make --version
```    
If Make is not installed, you can typically install it using your system's package manager:

```bash
	sudo apt-get install make
```   

-  **Docker:** 

The project relies on Docker for containerization. Check if Docker is installed:

```bash
	docker --version 
```  
If Docker is not installed, you can download and install it from the official Docker website: [https://www.docker.com/get-started](https://www.docker.com/get-started).

-  **Docker Compose:**

Docker Compose is used for defining and running multi-container Docker applications. Check if Docker Compose is installed:


```bash
	docker-compose --version
```

If Docker Compose is not installed, you can download and install it from the official Docker Compose website: [https://docs.docker.com/compose/install/](https://docs.docker.com/compose/install/).

### Installation Steps

To install the project, follow these steps:

```bash
	git clone https://github.com/AbderraoufKhorchani/web-scraper.git
	cd web-scraper
	make -C project install
```

**Note:** Ensure that ports 5432 and 8080 are available and not in use by other applications on your system.


## Usage

- **Run the Project:**
    
    After completing the installation steps, run the project using the following command:
    
```bash
	make -C project up_build
```
    
The scraping of the website and saving to the database occurs automatically the first time you run the project.
    
- **Retrieve All Quotes:**
    
    To retrieve all quotes, you can use the following endpoint:
    
```bash
	curl http://localhost:8080/all
```
    
-  **Retrieve Quotes Based on Tag:**
    
    To retrieve quotes based on a specific tag, replace "desired_tag" with the desired tag:
    
```bash
	curl http://localhost:8080/tag/desired_tag
```    
    
- **Retrieve Quotes Based on Author:**
    
    To retrieve quotes based on a specific author, replace "author_name" with the desired author:
    
 ```bash
     curl http://localhost:8080/author/author_name
```
    
- **Retrieve All Tags:**
    
    To retrieve all tags, use the following endpoint:
    
```bash
	curl http://localhost:8080/tags
```


