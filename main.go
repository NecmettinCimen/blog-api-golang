package main

import (
	"blog-api-golang/api"
	"blog-api-golang/database"
	_ "blog-api-golang/docs"
	"blog-api-golang/models"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
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
func main() {
	db := database.ConnectPostgres()

	err := db.AutoMigrate(&models.Blog{})
	if err != nil {
		log.Fatal(err)
	}
	r := gin.New()

	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	blog.SetRoutes(r)
	r.Run()
}
