package main

import (
	"blog-api-golang/api"
	_ "blog-api-golang/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
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
	r := gin.New()

	// use ginSwagger middleware to serve the API docs
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.GET("/api/v1/blog", api.GetBlogList)

	r.Run()
}
