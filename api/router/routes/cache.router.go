package routes

import (
	cacheHandler "github.com/edmiltonVinicius/register-steps/api/handler/cache"
	"github.com/gin-gonic/gin"
)

func Cache(r *gin.RouterGroup) {
	r.DELETE("/cache", cacheHandler.CleanCache)
}
