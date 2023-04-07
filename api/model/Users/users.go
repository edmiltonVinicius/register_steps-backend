package model_users

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	FirstName string   `gorm:"column:first_name; not null; serializer:json; comment:First name the user"`
	LastName  string   `gorm:"column:last_name; not null; serializer:json; comment:Last name the user"`
	Email     string   `gorm:"unique; not null; serializer:json; comment:E-mail the user"`
	Password  string   `gorm:"password; not null; serializer:json; comment:Password the user"`
	Country   string   `gorm:"country; not null; serializer:json; comment:Country the user"`
}

func (model *Users) TableName() string {
	return "users"
}