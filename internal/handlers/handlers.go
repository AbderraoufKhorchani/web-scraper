package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAll godoc
// @Summary Retrieves all quotes
// @Description Retrieves all quotes from the database.
// @Tags All quotes
// @Success 200 {array} baseQuote "List of quotes"
// @Failure 500 {string} string "Internal server error"
// @Router /all [get]
func GetAll(c *gin.Context) {

	all, err := GetAllDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}
	c.JSON(http.StatusOK, all)

}

// GetByTag godoc
// @Summary Retrieves quotes by tag
// @Description Retrieves quotes from the database based on the provided tag.
// @Tags Quotes by tag
// @Param tag path string true "Tag name"
// @Success 200 {array} baseQuote "List of quotes with the specified tag"
// @Failure 500 {string} string "Internal server error"
// @Router /tag/{tag} [get]
func GetByTag(c *gin.Context) {

	tag := c.Param("tag")

	byTag, err := GetByTagDB(tag)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}
	c.JSON(http.StatusOK, byTag)

}

// GetAllTags godoc
// @Summary Retrieves all tags
// @Description Retrieves all tags from the database.
// @Tags All tags
// @Success 200 {array} string "List of tags"
// @Failure 500 {string} string "Internal server error"
// @Router /tags [get]
func GetAllTags(c *gin.Context) {

	all, err := GetAllTagsDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}
	c.JSON(http.StatusOK, all)

}

// GetByAuthor godoc
// @Summary Retrieves quotes by author
// @Description Retrieves quotes from the database based on the provided author name.
// @Tags Quotes by author
// @Param author path string true "Author's name"
// @Success 200 {array} baseQuote "List of quotes by the specified author"
// @Failure 500 {string} string "Internal server error"
// @Router /author/{author} [get]
func GetByAuthor(c *gin.Context) {

	author := c.Param("author")

	byAuthor, err := GetByAuthorDB(author)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}
	c.JSON(http.StatusOK, byAuthor)

}
