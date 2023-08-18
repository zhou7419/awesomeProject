package userService

import (
	"awesomeProject/app/model/user"
	"awesomeProject/app/request/userReq"
	"awesomeProject/pkg/db"
	"awesomeProject/pkg/helpers"
)

type UserService struct {
}

func (us *UserService) GetByLogin(account string, password string) (userModel *user.User, err error) {
	field := []string{"name", "email", "account", "last_login"}
	err = db.Db.Select(field).Where("account = ? AND password = ?", account, helpers.Str2Md5(password)).Find(&userModel).Error
	return
}

func (us *UserService) Add(data userReq.Add) (err error) {
	userModel := user.User{
		Name:     data.Name,
		Email:    data.Email,
		Account:  data.Account,
		Password: data.Password,
	}
	err = db.Db.Create(&userModel).Error
	return
}

func (us *UserService) Index(data userReq.Search) (list []user.User, total int64, err error) {
	limit := data.Limit
	offset := (data.Page - 1) * data.Limit
	var userModel user.User
	query := db.Db.Model(&userModel)
	if err = query.Count(&total).Error; err != nil {
		return
	}
	if err = query.Offset(offset).Limit(limit).Find(&list).Error; err != nil {
		return
	}
	return
}

func (us *UserService) GetById(id string) (userModel user.User, err error) {
	field := []string{"name", "email", "account", "last_login", "created_at", "updated_at"}
	err = db.Db.Select(field).Where("id = ?", id).First(&userModel).Error
	return
}

func (us *UserService) Update(data userReq.Update) (err error) {
	var userData user.User
	if err = db.Db.Where("id = ?", data.ID).First(&userData).Error; err != nil {
		return
	}

	userData.Name = data.Name
	userData.Email = data.Email
	err = db.Db.Save(userData).Error
	return
}

func (us *UserService) Delete(id string) (err error) {
	var userModel user.User
	err = db.Db.Where("id = ?", id).Delete(&userModel).Error
	return
}
