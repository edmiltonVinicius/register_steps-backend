package routes

import (
	"github.com/edmiltonVinicius/register-steps/api/handler"
	"github.com/gin-gonic/gin"
)

func Health(r *gin.RouterGroup) {
	r.GET("/health-check", handler.HealthCheck)
}
