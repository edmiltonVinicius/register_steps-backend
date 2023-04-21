package model

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	FirstName string `gorm:"column:first_name; not null;"`
	LastName  string `gorm:"column:last_name; not null;"`
	Email     string `gorm:"unique; not null;"`
	Password  string `gorm:"password; not null;"`
	Country   string `gorm:"country; not null;"`
}
