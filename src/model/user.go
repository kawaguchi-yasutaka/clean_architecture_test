package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email    string
	Name     string
	Password string
}
