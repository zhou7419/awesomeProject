package note

import (
	"awesomeProject/app/model"
)

type Note struct {
	model.BaseModel
	UserId      int    `json:"user_id,omitempty" `
	ContentType int    `json:"content_type,omitempty"`
	RemindType  int    `json:"remind_type,omitempty"`
	Date        string `json:"date,omitempty"`
	model.BaseModelTimestamp
}
