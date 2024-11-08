package blog

import (
	"blog-api-golang/database"
	"blog-api-golang/models"
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"strconv"
)

func BlogApiSetRoutes(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	r.GET("/api/v1/blog", GetBlogList)

	auth := r.Group("").Use(authMiddleware.MiddlewareFunc())

	auth.POST("/api/v1/blog", CreatBlog)
	auth.PUT("/api/v1/blog", UpdateBlog)
	auth.DELETE("/api/v1/blog", DeleteBlog)
}

// GetBlogList example
//
//	@Accept			json
//	@Produce		json
//	@Tags         	blog
//	@Router			/blog [get]
//	@Param			take query int true "model"
//	@Param			skip query int true "model"
//	@Success		200 {object} []models.Blog
func GetBlogList(c *gin.Context) {
	take, err := strconv.Atoi(c.Query("take"))
	skip, err := strconv.Atoi(c.Query("skip"))
	if err != nil {
	}

	db := database.ConnectPostgres()

	if skip != 0 {
		db = db.Offset(skip)
	}
	if take != 0 {
		db = db.Limit(take)
	}

	var blogs []models.Blog
	db.Find(&blogs)

	var total int
	db.Table("blogs").Select("count(*) as total").Scan(&total)
	var response models.ListResponse
	response.Total = total
	response.List = blogs

	c.IndentedJSON(200, response)
}

// CreatBlog example
//
//		@Accept			json
//		@Produce		json
//		@Tags         	blog
//		@Router			/blog [post]
//	 	@Param			model body models.Blog true "model"
//		@Success		200 {object} models.Blog
func CreatBlog(c *gin.Context) {
	var body models.Blog
	c.BindJSON(&body)

	db := database.ConnectPostgres()

	result := db.Create(&body)

	c.IndentedJSON(200, result)
}

// UpdateBlog example
//
//		@Accept			json
//		@Produce		json
//		@Tags         	blog
//		@Router			/blog [put]
//	 	@Param			model body models.Blog true "model"
//		@Success		200 {object} models.Blog
func UpdateBlog(c *gin.Context) {
	var body models.Blog
	c.BindJSON(&body)

	db := database.ConnectPostgres()

	result := db.Updates(&body)

	c.IndentedJSON(200, result)
}

// DeleteBlog example
//
//		@Accept			json
//		@Produce		json
//		@Tags         	blog
//		@Router			/blog [delete]
//	 	@Param			model body models.Blog true "model"
//		@Success		200 {object} models.Blog
func DeleteBlog(c *gin.Context) {
	var body models.Blog
	c.BindJSON(&body)

	db := database.ConnectPostgres()

	result := db.Delete(&body)

	c.IndentedJSON(200, result)
}
