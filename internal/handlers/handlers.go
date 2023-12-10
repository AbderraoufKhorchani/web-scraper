package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {

	var quoteInstance Quote

	all, err := quoteInstance.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}
	c.JSON(http.StatusOK, all)

}

func GetAllTags(c *gin.Context) {

	var quoteInstance Quote

	all, err := quoteInstance.GetAllTags()

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}
	c.JSON(http.StatusOK, all)

}

func GetByAuthor(c *gin.Context) {

	author := c.Param("author")

	var quoteInstance Quote

	byAuthor, err := quoteInstance.GetByAuthor(author)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}
	c.JSON(http.StatusOK, byAuthor)

}

func GetByTag(c *gin.Context) {

	tag := c.Param("tag")

	var quoteInstance Quote

	byTag, err := quoteInstance.GetByTag(tag)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}
	c.JSON(http.StatusOK, byTag)

}
