package models

import (
	"github.com/jinzhu/gorm"
)

type QuestionType struct {
	gorm.Model
	Name string `gorm:"NOT NULL"`
}
