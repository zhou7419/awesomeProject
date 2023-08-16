package notes

import (
	"awesomeProject/app/model/note"
	"awesomeProject/pkg/db"
	"awesomeProject/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Notes struct {
	Type string `form:"type" binding:"required"`
	Date string `form:"date" binding:"required"`
}

func Add(c *gin.Context) {
	var notes Notes
	if err := c.Bind(&notes); err != nil {
		response.BadRequest(c, err)
		return
	}

	data := note.Note{
		Type: notes.Type,
		Date: notes.Date,
	}

	db.Db.Create(&data)
	response.Success(c)
}

func Index(c *gin.Context) {
	var notesModel []note.Note
	db.Db.Limit(20).Offset(0).Order("id desc").Find(&notesModel)
	response.Data(c, notesModel)
}

func Detail(c *gin.Context) {
	id := c.Param("id")
	var noteModel note.Note
	db.Db.First(&noteModel, id)
	response.Data(c, noteModel)
}

func Update(c *gin.Context) {
	id := c.PostForm("id")
	var noteModel note.Note
	result := db.Db.First(&noteModel, id)
	if result.RowsAffected < 1 {
		response.Abort404(c)
		return
	}

	var notes Notes
	if err := c.Bind(&notes); err != nil {
		response.BadRequest(c, err)
		return
	}

	noteModel.Type = notes.Type
	noteModel.Date = notes.Date
	fmt.Println(&notes)
	db.Db.Model(&noteModel).Save(&noteModel)

	response.Success(c)
}
