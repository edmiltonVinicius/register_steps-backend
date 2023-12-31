package handlers_user

import (
	"encoding/json"
	"net/http"

	dto "github.com/edmiltonVinicius/register-steps/api/dto/user"
	"github.com/edmiltonVinicius/register-steps/api/handler/contract"
	userService "github.com/edmiltonVinicius/register-steps/api/service/user"
	"github.com/gin-gonic/gin"
)

func Create(g *gin.Context) {
	var body dto.CreateUserInputDTO
	decoder := json.NewDecoder(g.Request.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&body)
	if err != nil {
		r := contract.JsonResponse{
			StatusCode: http.StatusBadRequest,
			Errors: []contract.ContractError{{
				Field:   "Body",
				Message: err.Error(),
			}},
			Data: nil,
		}
		r.SendJsonResponse(g)
		return
	}

	service := userService.NewUserService()

	errCreate := service.Create(&body)
	if errCreate != nil {
		r := contract.JsonResponse{
			StatusCode: http.StatusBadRequest,
			Errors:     errCreate,
			Data:       nil,
		}
		r.SendJsonResponse(g)
		return
	}

	r := contract.JsonResponse{
		StatusCode: http.StatusCreated,
		Errors:     nil,
		Data:       nil,
	}
	r.SendJsonResponse(g)
}

func GetByEmail(g *gin.Context) {
	email := g.Param("email")
	service := userService.NewUserService()

	res, err := service.FindByEmail(email)
	if err != nil {
		r := contract.JsonResponse{
			StatusCode: http.StatusBadRequest,
			Errors:     err,
			Data:       nil,
		}
		r.SendJsonResponse(g)
		return
	}

	r := contract.JsonResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data:       res,
	}
	r.SendJsonResponse(g)
}
