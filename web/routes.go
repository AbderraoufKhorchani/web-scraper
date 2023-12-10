package web

import (
	"github.com/AbderraoufKhorchani/web-scraper/internal/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET"}
	config.AllowHeaders = []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"}
	config.ExposeHeaders = []string{"Link"}
	config.AllowCredentials = true
	config.MaxAge = 300
	r.Use(cors.New(config))

	r.GET("/all", handlers.GetAll)
	r.GET("/author/:author", handlers.GetByAuthor)
	r.GET("/tag/:tag", handlers.GetByTag)
	r.GET("/tags", handlers.GetAllTags)

	return r
}
