package services

import (
	"testing"

	dto "github.com/edmiltonVinicius/register-steps/api/dto/user"
	"github.com/edmiltonVinicius/register-steps/api/handler/contract"
	"github.com/edmiltonVinicius/register-steps/config"
	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	suite.Suite
	service IUserService
}

func (suite *UserServiceTestSuite) SetupSuite() {
	config.StartDependencies()
	suite.service = NewUserService()
}

func (suite *UserServiceTestSuite) TearDownSuite() {
	config.ClearTable("users")
	config.ClearRedis()
	config.DownDependencies()
}

// Should create a new user
func (suite *UserServiceTestSuite) Test01_SuccessFullCreation() {
	data := dto.CreateUserInputDTO{
		UserType:       "freelancer",
		FirstName:      "Andre",
		LastName:       "Silva",
		Email:          "andre-silva@gmail.com",
		Password:       "12345678",
		RepeatPassword: "12345678",
		Country:        "brasil",
	}

	e := suite.service.Create(&data)

	suite.Nil(e)
	// suite.Equal(fmt.Sprintf("%s %s", data.FirstName, data.LastName), u.UserName)
}

// Should return error be email already exists
func (suite *UserServiceTestSuite) Test02_DuplicatedEmail() {
	data := dto.CreateUserInputDTO{
		UserType:       "freelancer2",
		FirstName:      "Andre2",
		LastName:       "Silva",
		Email:          "andre-silva2@gmail.com",
		Password:       "12345678",
		RepeatPassword: "12345678",
		Country:        "brasil",
	}

	e := suite.service.Create(&data)

	// suite.Nil(u)
	suite.Equal([]contract.ContractError{{
		Field:   "Email",
		Message: "User already exists.",
	}}, e)
}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}
