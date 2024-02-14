package handler

import (
	"net/http"

	contracts "github.com/edmiltonVinicius/register-steps/api/handler/contract"
	"github.com/gin-gonic/gin"
)

// @Summary		Health check
// @Description	This endpoint is used to check if the server is running
// @Tags			Health check
// @Produce		json
// @Success		200	{object}	contract.JsonResponse
// @Failure		400	{object}	contract.JsonResponse
// @Router			/health-check [get]
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
