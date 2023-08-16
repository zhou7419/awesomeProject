package user

import "gorm.io/gorm"

type User struct {
	Name     string
	Email    string
	Account  string
	Password string
	gorm.Model
}
