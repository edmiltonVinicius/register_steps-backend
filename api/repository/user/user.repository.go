package userRepository

import (
	"errors"

	dto "github.com/edmiltonVinicius/register-steps/api/dto/user"
	"github.com/edmiltonVinicius/register-steps/api/entity"
	"github.com/edmiltonVinicius/register-steps/cache"
	"github.com/edmiltonVinicius/register-steps/config"
	"gorm.io/gorm"
)

const tableName = "users"

func (ur *UserRepository) Create(data *dto.CreateUserInputDTO) (err error) {
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

	_ = ur.cache.SetJSon(data.Email, &u, cache.TLL_TEN_DAYS)
	return
}

func (ur *UserRepository) FindByEmail(email string, user *entity.User) (err error) {
	err = ur.cache.GetJSon(email, &user)
	if err == nil && user.Email != "" {
		return
	}
	err = nil

	er := config.DB.Model(&entity.User{}).Omit("password", "updated_at").Where("email = ?", email).Scan(&user)
	if er.Error != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		err = er.Error
		return
	}

	_ = ur.cache.SetJSon(user.Email, &user, cache.TLL_TEN_DAYS)
	return
}
