package handler

import (
	services "github.com/edmiltonVinicius/register-steps/api/service/user"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service services.IUserService
}

func NewUserHandler(us *services.IUserService) IUserHandler {
	return &UserHandler{
		service: *us,
	}
}

type IUserHandler interface {
	CreateUser(g *gin.Context)
	GetByEmail(g *gin.Context)
}
