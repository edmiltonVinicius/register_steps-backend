package services

import (
	"errors"

	"github.com/edmiltonVinicius/register-steps/api/dtos"
	"github.com/edmiltonVinicius/register-steps/api/handlers/contracts"
	"github.com/edmiltonVinicius/register-steps/api/model"
	"github.com/edmiltonVinicius/register-steps/api/repository"
	"github.com/edmiltonVinicius/register-steps/domain"
)

type UserService interface {
	Create(data *dtos.CreateUserInputDTO) (err []contracts.ContractError, res *dtos.CreateUserOutPutDTO)
	FindByEmail(email string) (user model.Users, errs []contracts.ContractError)
}

type userService struct{}

func NewUserService() UserService {
	return &userService{}
}

func (u *userService) Create(data *dtos.CreateUserInputDTO) (errs []contracts.ContractError, res *dtos.CreateUserOutPutDTO) {
	err := domain.Validate.Struct(data)
	if err != nil {
		if errors.As(err, &domain.ValidationErrors) {
			errs = domain.RunValidator(domain.ValidationErrors)
		} else {
			errs = []contracts.ContractError{{
				Field:   "Internal error",
				Message: err.Error(),
			}}
		}
		return
	}

	ur := repository.NewUserRepository()
	var us model.Users

	err = ur.FindByEmail(data.Email, &us)
	if err == nil && us.Email != "" {
		errs = []contracts.ContractError{{
			Field:   "Email",
			Message: "User already exists.",
		}}
		return
	}

	err = ur.Create(data)
	if err != nil {
		errs = []contracts.ContractError{{
			Field:   "Create",
			Message: "Error in create new user.",
		}}
		return
	}

	res = &dtos.CreateUserOutPutDTO{
		UserName: data.FirstName + " " + data.LastName,
	}
	return
}

func (u *userService) FindByEmail(email string) (user model.Users, errs []contracts.ContractError) {
	ur := repository.NewUserRepository()

	err := ur.FindByEmail(email, &user)
	if err != nil {
		errs = []contracts.ContractError{{
			Field:   "Error finding user",
			Message: "Error find user, please try again.",
		}}
	}
	return
}
