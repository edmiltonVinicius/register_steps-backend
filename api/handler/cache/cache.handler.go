package handler_cache

import (
	"net/http"

	"github.com/edmiltonVinicius/register-steps/api/handler/contract"
	"github.com/edmiltonVinicius/register-steps/config"
	"github.com/gin-gonic/gin"
)

// @Summary		Clear cache
// @Description	This endpoint is used to clear the cache
// @Tags			Cache
// @Produce		json
// @Success		201	{object}	contract.JsonResponse
// @Failure		400	{object}	contract.JsonResponse
// @Router			/cache [delete]
func CleanCache(g *gin.Context) {
	var r contract.JsonResponse
	res, err := config.RedisClient.FlushAll(config.Environment.CTX).Result()

	if err != nil {
		r = contract.JsonResponse{
			StatusCode: http.StatusBadRequest,
			Errors: []contract.ContractError{{
				Field:   "Redis error",
				Message: err.Error(),
			}},
			Data: nil,
		}
	} else {
		r = contract.JsonResponse{
			StatusCode: http.StatusOK,
			Errors:     nil,
			Data:       res,
		}
	}

	r.SendJsonResponse(g)
}
