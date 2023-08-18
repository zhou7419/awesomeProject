package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID uint64 `json:"id,omitempty" gorm:"column:id;primaryKey"`
}

type BaseModelTimestamp struct {
	//CreatedAt CustomTime `json:"created_at,omitempty" gorm:"column:created_at"`
	CreatedAt time.Time      `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt time.Time      `json:"updated_at,omitempty" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"column:deleted_at"`
}

//// CustomTime 自定义时间
//type CustomTime struct {
//	time.Time
//}
//
//// MarshalJson 如果有omitempty标签，未查询时间时返回空
//func (ct CustomTime) MarshalJson() ([]byte, error) {
//	if ct.IsZero() {
//		return []byte("null"), nil
//	}
//
//	return ct.Time.MarshalJSON()
//}
