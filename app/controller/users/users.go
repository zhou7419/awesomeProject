package users

import (
	"awesomeProject/app/model/user"
	"awesomeProject/pkg/db"
	"awesomeProject/pkg/helpers"
	"awesomeProject/pkg/jwt"
	"awesomeProject/pkg/response"
	"github.com/gin-gonic/gin"
)

type Users struct {
	Name     string `form:"name" binding:"required"`
	Email    string `form:"email" binding:"required"`
	Account  string `form:"account" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func Test(c *gin.Context) {
	response.Success(c)
}

func Add(c *gin.Context) {
	var usersRequest Users
	if err := c.Bind(&usersRequest); err != nil {
		response.BadRequest(c, err)
		return
	}

	var userModel user.User
	userModel.Name = usersRequest.Name
	userModel.Email = usersRequest.Email
	userModel.Account = usersRequest.Account
	userModel.Password = helpers.Str2Md5(usersRequest.Password)
	db.Db.Create(&userModel)

	response.Success(c)
}

type LoginRequest struct {
	Account  string `form:"account" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var loginRequest LoginRequest
	if err := c.Bind(&loginRequest); err != nil {
		response.BadRequest(c, err)
		return
	}

	var userModel user.User
	result := db.Db.Where("account = ? and password = ?", loginRequest.Account, helpers.Str2Md5(loginRequest.Password)).First(&userModel)
	if result.RowsAffected < 1 {
		response.Abort403(c, "用户名或密码错误")
		return
	}

	token, err := jwt.Get(string(userModel.ID), userModel.Name)
	if err != nil {
		response.Abort500(c, "token生成失败")
		return
	}
	data := map[string]any{
		"user":  userModel,
		"token": token,
	}

	response.JSON(c, data)
}
