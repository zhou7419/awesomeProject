package route

import (
	"awesomeProject/app/controller/note"
	"awesomeProject/app/controller/send"
	"awesomeProject/app/controller/user"
	"awesomeProject/app/middleware"
	"awesomeProject/pkg/response"
	"github.com/gin-gonic/gin"
)

func InitRoute(r *gin.Engine) {
	noteGroup := r.Group("/notes")
	noteGroup.Use(middleware.AuthJWT())
	{
		noteGroup.POST("/add", note.Add)
		noteGroup.POST("/index", note.Index)
		noteGroup.POST("/detail", note.Detail)
		noteGroup.POST("/update", note.Update)
		noteGroup.POST("/delete", note.Delete)
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
}
