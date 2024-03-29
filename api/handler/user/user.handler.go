package handler

import (
	"encoding/json"
	"net/http"

	dto "github.com/edmiltonVinicius/register-steps/api/dto/user"
	"github.com/edmiltonVinicius/register-steps/api/handler/contract"
	"github.com/gin-gonic/gin"

	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

// @Summary		Create user
// @Description	Create a new user
// @Tags			Users
// @Accept			json
// @Produce		json
// @Param			user	body		dto.CreateUserInputDTO	true	"Data required to create a new user"
// @Success		201	{object}	contract.JsonResponse
// @Failure		400	{object}	contract.JsonResponse
// @Failure		500	{object}	contract.JsonResponse
// @Router			/users [post]
func (uh *UserHandler) CreateUser(g *gin.Context) {
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

	errCreate := uh.service.Create(&body)
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

// @Summary		Get user by email
// @Tags			Users
// @Accept			json
// @Produce		json
// @Param			email	path		string	true	"Email of user searched"
// @Success		200	{object}	contract.JsonResponse
// @Failure		400	{object}	contract.JsonResponse
// @Failure		500	{object}	contract.JsonResponse
// @Router			/users/:email [get]
func (uh *UserHandler) GetByEmail(g *gin.Context) {
	email := g.Param("email")

	res, err := uh.service.FindByEmail(email)
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
