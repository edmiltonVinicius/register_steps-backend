package router

import (
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/edmiltonVinicius/register-steps/api/docs"
	"github.com/edmiltonVinicius/register-steps/api/router/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/edmiltonVinicius/register-steps/api/middleware"
	"github.com/edmiltonVinicius/register-steps/config"
	"github.com/gin-gonic/gin"
)

//	@title			Swagger Register steps API
//	@version		1.0
//	@description	This is a documentation for use API, write in Golang.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API support, use or new implementations
//	@contact.url	https://github.com/edmiltonVinicius
//	@contact.email	edmilton.vinicius2@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/v1

// @securityDefinitions.basic	BasicAuth
func LoadRoutes() {

	if config.Environment.ENV == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.Use(middleware.Recovery())

	g := router.Group("/v1")

	routes.Health(g)
	routes.Cache(g)
	routes.Users(g)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	s := &http.Server{
		Addr:           fmt.Sprintf(":%s", config.Environment.SERVER_PORT),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
