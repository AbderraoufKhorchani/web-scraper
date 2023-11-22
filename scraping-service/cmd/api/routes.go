package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (app *Api) Routes() *gin.Engine {

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET"}
	config.AllowHeaders = []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"}
	config.ExposeHeaders = []string{"Link"}
	config.AllowCredentials = true
	config.MaxAge = 300
	r.Use(cors.New(config))

	r.GET("/all", app.GetAll)
	r.GET("/author/:author", app.GetByAuthor)
	r.GET("/tag/:tag", app.GetByTag)
	r.GET("/tags", app.GetAllTags)

	return r
}
