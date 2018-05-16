package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Question struct {
	gorm.Model
	QuestionId uuid.UUID     `json:"questionId"`
	Answers    []interface{} `json:"answers" gorm:"type:text[]"`
}
