package user

import (
	"awesomeProject/app/model"
	"time"
)

type User struct {
	model.BaseModel
	Name      string     `json:"name,omitempty"`
	Email     string     `json:"email,omitempty"`
	Account   string     `json:"account,omitempty" gorm:"uniqueIndex"`
	Password  string     `json:"password,omitempty"`
	LastLogin *time.Time `json:"last_login,omitempty" gorm:"column:last_login;autoCreateTime:false;"`
	model.BaseModelTimestamp
}
