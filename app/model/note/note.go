package note

import "gorm.io/gorm"

type Note struct {
	Type string
	Date string
	gorm.Model
}
