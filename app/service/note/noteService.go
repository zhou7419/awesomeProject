package noteService

import (
	"awesomeProject/app/model/note"
	"awesomeProject/app/request/noteReq"
	"awesomeProject/pkg/db"
)

type NoteService struct {
}

func (ns *NoteService) Add(userId int, data noteReq.Add) (err error) {
	var noteModel note.Note
	noteModel.UserId = userId
	noteModel.ContentType = data.ContentType
	noteModel.RemindType = data.RemindType
	noteModel.Date = data.Date
	err = db.Db.Create(&noteModel).Error
	return
}

func (ns *NoteService) Index(data noteReq.Index) (list []note.Note, total int64, err error) {
	limit := data.Limit
	offset := (data.Page.Page - 1) * limit
	var noteModel note.Note
	query := db.Db.Model(&noteModel)
	if err = query.Count(&total).Error; err != nil {
		return
	}
	if err = query.Offset(offset).Limit(limit).Find(&list).Error; err != nil {
		return
	}
	return
}

func (ns *NoteService) GetById(id string) (NoteModel note.Note, err error) {
	field := []string{"id", "content_type", "remind_type", "date", "created_at", "updated_at"}
	err = db.Db.Select(field).Where("id = ?", id).First(&NoteModel).Error
	return
}

func (ns *NoteService) Update(data noteReq.Update) (err error) {
	var noteModel note.Note
	if err = db.Db.Where("id = ?", data.ID).First(&noteModel).Error; err != nil {
		return
	}

	noteModel.RemindType = data.RemindType
	noteModel.ContentType = data.ContentType
	noteModel.Date = data.Date
	err = db.Db.Save(&noteModel).Error
	return
}

func (ns *NoteService) Delete(id string) (err error) {
	var noteModel note.Note
	err = db.Db.Where("id = ?", id).Delete(&noteModel).Error
	return
}
