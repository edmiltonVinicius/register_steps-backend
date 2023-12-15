package handler

import (
	"net/http"

	contracts "github.com/edmiltonVinicius/register-steps/api/handler/contract"
	"github.com/gin-gonic/gin"
)

func HealthCheck(g *gin.Context) {
	r := contracts.JsonResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data: map[string]string{
			"Message":       "Welcome to the health-check endpoints server register-steps",
			"Administrator": "Edmilton Vinciius Pansanato",
		},
	}

	r.SendJsonResponse(g)
}
