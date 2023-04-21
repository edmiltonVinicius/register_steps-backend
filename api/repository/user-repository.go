package repository

import (
	"github.com/edmiltonVinicius/register-steps/api/dtos"
	"github.com/edmiltonVinicius/register-steps/api/model"
	"github.com/edmiltonVinicius/register-steps/domain"
	cache "github.com/edmiltonVinicius/register-steps/redis"
)

const tableName = "users"

type UserRepository interface {
	Create(data *dtos.CreateUserInputDTO) (err error)
	FindByEmail(email string, user *model.Users) (err error)
}

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (ur *userRepository) Create(data *dtos.CreateUserInputDTO) (err error) {
	u := model.Users{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Password:  data.Password,
		Country:   data.Country,
	}
	res := domain.DB.Table(tableName).Create(&u)
	if res.Error != nil {
		err = res.Error
	}
	_ = cache.SetJSon(data.Email, &u, cache.TLL_TEN_DAYS)
	return
}

func (ur *userRepository) FindByEmail(email string, user *model.Users) (err error) {
	err = cache.GetJSon(email, &user)
	if err == nil {
		return
	}
	err = nil

	er := domain.DB.Table(tableName).Limit(1).Where("email = ?", email).Scan(&user)
	if er.Error != nil {
		err = er.Error
	} else {
		_ = cache.SetJSon(user.Email, &user, cache.TLL_TEN_DAYS)
	}
	return
}
