package handlers_user

import (
	"encoding/json"
	"net/http"

	dto "github.com/edmiltonVinicius/register-steps/api/dto/user"
	"github.com/edmiltonVinicius/register-steps/api/handler/contract"
	userService "github.com/edmiltonVinicius/register-steps/api/service/user"
	"github.com/gin-gonic/gin"

	_ "github.com/swaggo/files"       // swagger embed files
	_ "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var service userService.IUserService

func init() {
	service = userService.NewUserService()
}

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

// @Summary		Get user by email
// @Tags			Users
// @Accept			json
// @Produce		json
// @Param			email	path		string	true	"Email of user searched"
// @Success		200	{object}	contract.JsonResponse
// @Failure		400	{object}	contract.JsonResponse
// @Failure		500	{object}	contract.JsonResponse
// @Router			/users/:email [get]
func GetByEmail(g *gin.Context) {
	email := g.Param("email")

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
