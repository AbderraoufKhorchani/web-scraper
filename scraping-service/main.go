package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/AbderraoufKhorchani/web-scraper/scraping-service/cmd/api"
	"github.com/AbderraoufKhorchani/web-scraper/scraping-service/data"
	"github.com/AbderraoufKhorchani/web-scraper/scraping-service/scraper"
)

var counts int64

const webPort = "80"

type Config struct {
	DB      *gorm.DB
	Api     api.Api
	Models  data.Models
	Scraper scraper.Scraper
}

func main() {
	fmt.Println("1")
	conn := connectToDB()
	if conn == nil {
		log.Panic("Can't connect to Postgres!")
	}
	fmt.Println("2")

	app := Config{
		DB:      conn,
		Models:  data.New(conn),
		Api:     api.Api{},
		Scraper: scraper.Scraper{},
	}
	fmt.Println("3")
	app.Scraper.Scrape()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.Api.Routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}

func connectToDB() *gorm.DB {
	dsn := os.Getenv("DSN")
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
