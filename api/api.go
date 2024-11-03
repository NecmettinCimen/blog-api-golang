package api

import (
	"blog-api-golang/models"
	"github.com/gin-gonic/gin"
)

// GetBlogList example
//
//	@Accept			json
//	@Produce		json
//	@Tags         	blog
//	@Router			/blog [get]
//	@Success		200 {object} []models.Blog
func GetBlogList(c *gin.Context) {
	var blogs = []models.Blog{}
	c.IndentedJSON(200, blogs)
}
