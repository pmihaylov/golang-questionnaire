package models

import (
	"github.com/jinzhu/gorm"
	"github.com/google/uuid"
)

type Question struct {
	gorm.Model
	QuestionId uuid.UUID `json:"questionId"`
	Answers    []interface{} `json:"answers" gorm:"type:text[]"`
}
