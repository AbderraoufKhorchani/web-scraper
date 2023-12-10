package web

import (
	_ "github.com/AbderraoufKhorchani/web-scraper/docs"
	"github.com/AbderraoufKhorchani/web-scraper/internal/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	//add swagger
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/all", handlers.GetAll)
	r.GET("/author/:author", handlers.GetByAuthor)
	r.GET("/tag/:tag", handlers.GetByTag)
	r.GET("/tags", handlers.GetAllTags)

	return r
}
