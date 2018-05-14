package models

import (
	"github.com/jinzhu/gorm"
)

type Question struct {
	gorm.Model
	QuestionId string `json:"questionId"`
	Answers    []interface{} `json:"answers" gorm:"-"`
}
