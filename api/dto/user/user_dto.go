package dto

// @TODO: add eq=freelancer,eq=client into user_type validate

type CreateUserInputDTO struct {
	UserType       string `json:"user_type" validate:"required"`
	FirstName      string `json:"first_name" validate:"required,min=3,max=20"`
	LastName       string `json:"last_name" validate:"required,min=3,max=20"`
	Email          string `json:"email" validate:"required,email"`
	Password       string `json:"password" validate:"required,containsany=!@#$%&*(),eqfield=RepeatPassword"`
	RepeatPassword string `json:"repeat_password" validate:"required,containsany=!@#$%&*()"`
	Country        string `json:"country" validate:"required"`
}

type CreateUserOutPutDTO struct {
	UserName string `json:"userName"`
}
