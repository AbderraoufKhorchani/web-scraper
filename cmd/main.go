package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AbderraoufKhorchani/web-scraper/internal/handlers"
	"github.com/AbderraoufKhorchani/web-scraper/internal/scraper"
	"github.com/AbderraoufKhorchani/web-scraper/web"
)

const (
	dsn     = "host=localhost port=5432 user=postgres password=password dbname=quotes sslmode=disable timezone=UTC connect_timeout=5"
	webPort = "8080"
)

func main() {
	conn := handlers.ConnectToDB(dsn)
	if conn == nil {
		log.Panic("Can't connect to Postgres!")
	}

	handlers.New(conn)

	scraper.Scrape()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: web.Routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}
