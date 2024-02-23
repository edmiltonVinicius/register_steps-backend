package userRepository

import (
	dto "github.com/edmiltonVinicius/register-steps/api/dto/user"
	"github.com/edmiltonVinicius/register-steps/api/entity"
	"github.com/edmiltonVinicius/register-steps/cache"
)

type UserRepository struct {
	cache cache.ICache
}

func NewUserRepository(c *cache.ICache) IUserRepository {
	return &UserRepository{
		cache: *c,
	}
}

type IUserRepository interface {
	Create(data *dto.CreateUserInputDTO) (err error)
	FindByEmail(email string, user *entity.User) (err error)
}
