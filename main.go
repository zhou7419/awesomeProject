package main

import (
	"awesomeProject/app/controller/notes"
	"awesomeProject/app/controller/send"
	"awesomeProject/app/controller/user"
	"awesomeProject/app/middleware"
	"awesomeProject/pkg/db"
	"awesomeProject/pkg/response"
	"awesomeProject/pkg/timer"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {

	go timer.Timer()

	db.DbConnect()
	db.RdbConnect()

	f, _ := os.Create("./log/gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout) // gin 日志

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	noteGroup := r.Group("/notes")
	noteGroup.Use(middleware.AuthJWT())
	{
		noteGroup.POST("/add", notes.Add)
		noteGroup.POST("/update", notes.Update)
		noteGroup.POST("/index", notes.Index)
		noteGroup.POST("/detail", notes.Detail)
	}

	userGroup := r.Group("/users")
	userGroup.Use(middleware.AuthJWT())
	{
		userGroup.POST("/add", user.Add)
		userGroup.POST("/index", user.Index)
		userGroup.POST("/detail", user.Detail)
		userGroup.POST("/update", user.Update)
		userGroup.POST("/delete", user.Delete)
	}

	noAuthGroup := r.Group("/")
	{
		noAuthGroup.POST("/users/login", user.Login)
	}

	r.POST("/send", send.Mail, middleware.AuthJWT())

	r.NoRoute(func(c *gin.Context) {
		response.Abort404(c)
	})

	err := r.Run(":8089")
	if err != nil {
		return
	}
}
