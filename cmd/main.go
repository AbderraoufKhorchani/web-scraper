package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/AbderraoufKhorchani/web-scraper/internal/handlers"
	"github.com/AbderraoufKhorchani/web-scraper/internal/scraper"
	"github.com/AbderraoufKhorchani/web-scraper/web"
)

var counts int64

const (
	dsn     = "host=localhost port=5432 user=postgres password=password dbname=quotes sslmode=disable timezone=UTC connect_timeout=5"
	webPort = "8080"
)

func main() {
	conn := connectToDB()
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

func connectToDB() *gorm.DB {
	for {
		connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Println("Postgres not yet ready ...")
			counts++
		} else {
			log.Println("Connected to Postgres!")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for two seconds....")
		time.Sleep(2 * time.Second)
		continue
	}
}
