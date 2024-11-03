package main

import (
	blog "blog-api-golang/api"
	"blog-api-golang/database"
	_ "blog-api-golang/docs"
	"blog-api-golang/helper"
	"blog-api-golang/models"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

//	@title			Blog Api Golang
//	@version		1.0
//	@description	This is a sample go server blog api.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	https://necmettincimen.com.tr
//	@contact.email	mail@necmettincimen.com.tr

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:8080
// @BasePath	/api/v1
//
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				Description for what is this security definition being used
func main() {
	db := database.ConnectPostgres()

	err := db.AutoMigrate(&models.Blog{}, &models.User{})
	if err != nil {
		log.Fatal(err)
	}
	r := gin.New()

	auth, err := jwt.New(helper.SetJWTMid())

	r.Use(handlerMiddleWare(auth))

	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	blog.UserApiSetRoutes(r)
	blog.BlogApiSetRoutes(r, auth)
	err = r.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func handlerMiddleWare(authMiddleware *jwt.GinJWTMiddleware) gin.HandlerFunc {
	return func(context *gin.Context) {
		errInit := authMiddleware.MiddlewareInit()
		if errInit != nil {
			log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
		}
	}
}
