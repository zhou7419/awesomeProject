package main

import (
	"awesomeProject/app/middleware"
	"awesomeProject/pkg/db"
	"awesomeProject/pkg/logger"
	"awesomeProject/pkg/timer"
	"awesomeProject/route"
	"github.com/gin-gonic/gin"
)

func main() {

	logger.InitLogger()
	//logger.InitGormLogger()
	go timer.Timer()

	db.DbConnect()
	db.RdbConnect()

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.GinLogger())

	route.InitRoute(r)

	err := r.Run(":8089")
	if err != nil {
		return
	}
}
