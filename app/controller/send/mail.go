package send

import (
	"awesomeProject/auth"
	"awesomeProject/pkg/mail"
	"awesomeProject/pkg/response"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type DataRequest struct {
	Type string
	Time int
}

func Mail(c *gin.Context) {
	str := c.PostForm("str")
	key := "87ca2f3b550d6b51"
	res, err := auth.Verify(str, key)
	if err != nil {
		response.Abort403(c, err.Error())
		return
	}

	dataRequest := DataRequest{}
	err = json.Unmarshal([]byte(res), &dataRequest)
	if err != nil {
		response.Abort500(c, err.Error())
		return
	}

	err = mail.Send("测试标题", "测试内容-"+dataRequest.Type, "")
	if err != nil {
		response.Abort500(c, err.Error())
		return
	}

	response.Success(c)
}
