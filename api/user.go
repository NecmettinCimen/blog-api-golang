package blog

import (
	"blog-api-golang/database"
	"blog-api-golang/helper"
	"blog-api-golang/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

func UserApiSetRoutes(r *gin.Engine) {
	r.POST("/api/v1/login", login)
}

// CreatBlog example
//
//		@Accept			json
//		@Produce		json
//		@Tags         	auth
//		@Router			/login [post]
//	 	@Param			model body models.Login true "model"
//		@Success		200 {object} models.Login
func login(c *gin.Context) {
	var body models.Login
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	db := database.ConnectPostgres()

	var user models.User
	db.Model(&models.User{}).Where(&body).Scan(&user)

	if user.UserName != "" {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			helper.IdentityKey: user.ID,
		})
		tokenStr, err := token.SigningString()
		if err != nil {
		}

		c.JSON(http.StatusOK, gin.H{"token": tokenStr})

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
	}
}
