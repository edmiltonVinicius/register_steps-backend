package dtos

type CreateUserInputDTO struct {
	UserType       string `json:"user_type" validate:"required"`
	FirstName      string `json:"first_name" validate:"required,min=3,max=20"`
	LastName       string `json:"last_name" validate:"required,min=3,max=20"`
	Email          string `json:"email" validate:"required,email"`
	Password       string `json:"password" validate:"required"`
	RepeatPassword string `json:"repeat_password" validate:"required"`
	Country        string `json:"country" validate:"required"`
}

type CreateUserOutPutDTO struct {
	UserName string `json:"userName"`
}
