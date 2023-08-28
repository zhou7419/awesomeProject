package noteReq

import "awesomeProject/app/request"

type Add struct {
	ContentType int    `form:"content_type" binding:"required"`
	RemindType  int    `form:"remind_type" binding:"required"`
	Date        string `form:"date" binding:"required"`
}

type Index struct {
	request.Page
}

type Detail struct {
	ID string `form:"id" binding:"required"`
}

type Update struct {
	ID          int    `form:"id" binding:"required"`
	ContentType int    `form:"content_type" binding:"required"`
	RemindType  int    `form:"remind_type" binding:"required"`
	Date        string `form:"date" binding:"required"`
}

type Delete struct {
	ID string `form:"id" binding:"required"`
}
