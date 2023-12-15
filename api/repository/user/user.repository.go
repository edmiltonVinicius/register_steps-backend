package repository

import (
	dto "github.com/edmiltonVinicius/register-steps/api/dto/user"
	"github.com/edmiltonVinicius/register-steps/api/entity"
	"github.com/edmiltonVinicius/register-steps/config"
	cache "github.com/edmiltonVinicius/register-steps/redis"
)

const tableName = "users"

type userRepository struct {
}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}

func (ur *userRepository) Create(data *dto.CreateUserInputDTO) (err error) {
	u := entity.User{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Password:  data.Password,
		Country:   data.Country,
	}
	res := config.DB.Table(tableName).Create(&u)
	if res.Error != nil {
		err = res.Error
		return
	}
	_ = cache.SetJSon(data.Email, &u, cache.TLL_TEN_DAYS)
	return
}

func (ur *userRepository) FindByEmail(email string, user *entity.User) (err error) {
	err = cache.GetJSon(email, &user)
	if err == nil {
		return
	}
	err = nil

	er := config.DB.Table(tableName).Limit(1).Where("email = ?", email).Scan(&user)
	if er.Error != nil {
		err = er.Error
		return
	}

	_ = cache.SetJSon(user.Email, &user, cache.TLL_TEN_DAYS)
	return
}
