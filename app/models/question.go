package models

import (
	"github.com/jinzhu/gorm"
)

type Question struct {
	gorm.Model
	LibraryId      int    `json:"libraryId" gorm:"NOT NULL"`
	Required       bool   `json:"required"`
	Text           string `json:"title" gorm:"Type:TEXT"`
	QuestionTypeId int    `json:"questionTypeId" gorm:"NOT NULL"`
}
