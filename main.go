package main

import (
	"awesomeProject/app/controller/notes"
	"awesomeProject/app/controller/send"
	"awesomeProject/app/controller/users"
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
	r.Use(middleware.AuthJWT())

	notesGroup := r.Group("/notes")
	notesGroup.Use(middleware.AuthJWT())
	{
		notesGroup.POST("/add", notes.Add)
		notesGroup.POST("/update", notes.Update)
		notesGroup.POST("/index", notes.Index)
		notesGroup.POST("/detail", notes.Detail)
	}

	r.POST("/send", send.Mail, middleware.AuthJWT())

	r.POST("/users/login", users.Login)
	r.POST("users/add", users.Add, middleware.AuthJWT())

	r.NoRoute(func(c *gin.Context) {
		response.Abort404(c)
	})

	r.Run(":8089")
}
