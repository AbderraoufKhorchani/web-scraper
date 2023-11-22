package api

import (
	"log"
	"net/http"

	"github.com/AbderraoufKhorchani/web-scraper/scraping-service/data"
	"github.com/gin-gonic/gin"
)

func (api *Api) GetAll(c *gin.Context) {

	var quoteInstance data.Quote

	all, err := quoteInstance.GetAll()

	if err != nil {
		api.errorJSON(c, err)
		log.Println(err)
		return
	}
	api.writeJSON(c, http.StatusAccepted, all)

}

func (api *Api) GetByAuthor(c *gin.Context) {

	author := c.Param("author")

	var quoteInstance data.Quote

	byAuthor, err := quoteInstance.GetByAuthor(author)

	if err != nil {
		api.errorJSON(c, err)
		log.Println(err)
		return
	}
	api.writeJSON(c, http.StatusAccepted, byAuthor)

}

func (api *Api) GetByTag(c *gin.Context) {

	tag := c.Param("tag")

	var quoteInstance data.Quote

	byTag, err := quoteInstance.GetByTag(tag)

	if err != nil {
		api.errorJSON(c, err)
		log.Println(err)
		return
	}
	api.writeJSON(c, http.StatusAccepted, byTag)

}
