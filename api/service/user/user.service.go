package services

import (
	"errors"

	dto "github.com/edmiltonVinicius/register-steps/api/dto/user"
	"github.com/edmiltonVinicius/register-steps/api/entity"
	"github.com/edmiltonVinicius/register-steps/api/handler/contract"
	repository "github.com/edmiltonVinicius/register-steps/api/repository/user"
	"github.com/edmiltonVinicius/register-steps/api/utils"
	"github.com/edmiltonVinicius/register-steps/config"
)

type userService struct{}

func NewUserService() IUserService {
	return &userService{}
}

func (u *userService) Create(data *dto.CreateUserInputDTO) (errs []contract.ContractError) {
	err := config.Validate.Struct(data)
	if err != nil {
		if errors.As(err, &config.ValidationErrors) {
			errs = config.RunValidator(config.ValidationErrors)
		} else {
			errs = []contract.ContractError{{
				Field:   "Internal error",
				Message: err.Error(),
			}}
		}
		return
	}

	if data.Password != data.RepeatPassword {
		errs = []contract.ContractError{{
			Field:   "Password",
			Message: "Password and repeat password is not the same.",
		}}
		return
	}

	hash, e := utils.GenerateHashString(data.Password, 14)
	if e != nil {
		errs = []contract.ContractError{{
			Field:   "Password",
			Message: "Error in validation of password.",
		}}
		return
	}
	data.Password = hash

	ur := repository.NewUserRepository()
	var us entity.User

	err = ur.FindByEmail(data.Email, &us)
	if err != nil {
		errs = []contract.ContractError{{
			Field:   "Email",
			Message: "Error to validate",
		}}
		return
	}

	if us.Email != "" {
		errs = []contract.ContractError{{
			Field:   "Email",
			Message: "User already exists.",
		}}
		return
	}

	err = ur.Create(data)
	if err != nil {
		errs = []contract.ContractError{{
			Field:   "Create",
			Message: "Error in create new user.",
		}}
	}
	return
}

func (u *userService) FindByEmail(email string) (user entity.User, errs []contract.ContractError) {
	ur := repository.NewUserRepository()

	err := ur.FindByEmail(email, &user)
	if err != nil {
		errs = []contract.ContractError{{
			Field:   "Error finding user",
			Message: "Error find user, please try again.",
		}}
	}
	return
}
