package router

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/edmiltonVinicius/register-steps/api/handlers"
	handlerUser "github.com/edmiltonVinicius/register-steps/api/handlers/user"
	"github.com/edmiltonVinicius/register-steps/api/middleware"
	"github.com/edmiltonVinicius/register-steps/domain"
	"github.com/gin-gonic/gin"
)

func LoadRoutes() {
	router := gin.Default()

	router.Use(middleware.Recovery())

	r := router.Group("/v1")
	r.GET("/health-check", handlers.HealthCheck)

	user := r.Group("/users")
	user.POST("/", handlerUser.Create)
	user.GET("/:email", handlerUser.GetByEmail)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	s := &http.Server{
		Addr:           fmt.Sprintf(":%s", domain.Environment.SERVER_PORT),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
