package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				fmt.Println("Panic captured: ", err)
			}
		}()

		c.Next()
	}
}
