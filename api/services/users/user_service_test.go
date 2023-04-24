package services

import (
	"fmt"
	"testing"

	"github.com/edmiltonVinicius/register-steps/api/dtos"
	"github.com/edmiltonVinicius/register-steps/api/handlers/contracts"
	"github.com/edmiltonVinicius/register-steps/domain"
	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	suite.Suite
	service UserService
}

func (suite *UserServiceTestSuite) SetupSuite() {
	domain.StartDependencies()
	suite.service = NewUserService()
}

func (suite *UserServiceTestSuite) TearDownSuite() {
	domain.ClearTable("users")
	domain.ClearRedis()
	domain.DownDependencies()
}

// Should create a new user
func (suite *UserServiceTestSuite) Test01_SuccessFullCreation() {
	data := dtos.CreateUserInputDTO{
		UserType:       "freelancer",
		FirstName:      "Andre",
		LastName:       "Silva",
		Email:          "andre-silva@gmail.com",
		Password:       "12345678",
		RepeatPassword: "12345678",
		Country:        "brasil",
	}

	e, u := suite.service.Create(&data)

	suite.Nil(e)
	suite.Equal(fmt.Sprintf("%s %s", data.FirstName, data.LastName), u.UserName)
}

// Should return error be email already exists
func (suite *UserServiceTestSuite) Test02_DuplicatedEmail() {
	data := dtos.CreateUserInputDTO{
		UserType:       "freelancer",
		FirstName:      "Andre",
		LastName:       "Silva",
		Email:          "andre-silva@gmail.com",
		Password:       "12345678",
		RepeatPassword: "12345678",
		Country:        "brasil",
	}

	e, u := suite.service.Create(&data)

	suite.Nil(u)
	suite.Equal([]contracts.ContractError{{
		Field:   "Email",
		Message: "User already exists.",
	}}, e)
}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}
