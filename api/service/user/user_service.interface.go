package services

import (
	dto "github.com/edmiltonVinicius/register-steps/api/dto/user"
	"github.com/edmiltonVinicius/register-steps/api/entity"
	"github.com/edmiltonVinicius/register-steps/api/handler/contract"
)

type IUserService interface {
	Create(data *dto.CreateUserInputDTO) (err []contract.ContractError)
	FindByEmail(email string) (user entity.User, errs []contract.ContractError)
}
