package user

import (
	"awesomeProject/app/request/userReq"
	"awesomeProject/app/service/user"
	"awesomeProject/pkg/jwt"
	"awesomeProject/pkg/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

var us = new(userService.UserService)

func Add(c *gin.Context) {
	var addReq userReq.Add
	if err := c.Bind(&addReq); err != nil {
		response.BadRequest(c, err)
		return
	}

	if err := us.Add(addReq); err != nil {
		response.BadRequest(c, err)
		return
	}

	response.Success(c)
}

func Index(c *gin.Context) {
	var searchReq userReq.Search
	if err := c.Bind(&searchReq); err != nil {
		response.BadRequest(c, err)
		return
	}

	list, total, err := us.Index(searchReq)
	if err != nil {
		response.BadRequest(c, err)
	}

	response.Data(c, gin.H{
		"list":  list,
		"total": total,
	})
}

func Detail(c *gin.Context) {
	var detailReq userReq.Detail
	if err := c.Bind(&detailReq); err != nil {
		response.BadRequest(c, err)
		return
	}

	data, err := us.GetById(detailReq.ID)
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	response.Data(c, data)
}

func Update(c *gin.Context) {
	var updateReq userReq.Update
	if err := c.ShouldBind(&updateReq); err != nil {
		response.BadRequest(c, err)
		return
	}

	err := us.Update(updateReq)
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	response.Success(c)
}

func Delete(c *gin.Context) {
	var deleteReq userReq.Delete
	if err := c.Bind(&deleteReq); err != nil {
		response.BadRequest(c, err)
		return
	}

	if err := us.Delete(deleteReq.ID); err != nil {
		response.BadRequest(c, err)
		return
	}

	response.Success(c)
}

func Login(c *gin.Context) {
	var loginReq userReq.Login
	if err := c.Bind(&loginReq); err != nil {
		response.BadRequest(c, err)
		return
	}

	result, err := us.GetByLogin(loginReq.Account, loginReq.Password)
	if err != nil {
		response.Abort403(c, "用户名或密码错误")
		return
	}

	token, err := jwt.Get(strconv.Itoa(int(result.ID)), result.Name)
	if err != nil {
		response.Abort500(c, "token生成失败")
		return
	}
	data := map[string]any{
		"user":  result,
		"token": token,
	}

	response.JSON(c, data)
}
