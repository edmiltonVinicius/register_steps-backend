package repository

import (
	dto "github.com/edmiltonVinicius/register-steps/api/dto/user"
	"github.com/edmiltonVinicius/register-steps/api/entity"
)

type IUserRepository interface {
	Create(data *dto.CreateUserInputDTO) (err error)
	FindByEmail(email string, user *entity.User) (err error)
}
