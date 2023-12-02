package handlers_users

import (
	"encoding/json"
	"net/http"

	dto "github.com/edmiltonVinicius/register-steps/api/dto/user"
	"github.com/edmiltonVinicius/register-steps/api/handlers/contracts"
	service "github.com/edmiltonVinicius/register-steps/api/services/users"
	"github.com/gin-gonic/gin"
)

func Create(g *gin.Context) {
	var body dto.CreateUserInputDTO
	decoder := json.NewDecoder(g.Request.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&body)
	if err != nil {
		r := contracts.JsonResponse{
			StatusCode: http.StatusBadRequest,
			Errors: []contracts.ContractError{{
				Field:   "Body",
				Message: err.Error(),
			}},
			Data: nil,
		}
		r.SendJsonResponse(g)
		return
	}

	service := service.NewUserService()

	errCreate, res := service.Create(&body)
	if errCreate != nil {
		r := contracts.JsonResponse{
			StatusCode: http.StatusBadRequest,
			Errors:     errCreate,
			Data:       nil,
		}
		r.SendJsonResponse(g)
		return
	}

	r := contracts.JsonResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data:       res,
	}
	r.SendJsonResponse(g)
}

func GetByEmail(g *gin.Context) {
	email := g.Param("email")
	service := service.NewUserService()

	res, err := service.FindByEmail(email)
	if err != nil {
		r := contracts.JsonResponse{
			StatusCode: http.StatusBadRequest,
			Errors:     err,
			Data:       nil,
		}
		r.SendJsonResponse(g)
		return
	}

	r := contracts.JsonResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data:       res,
	}
	r.SendJsonResponse(g)
}
