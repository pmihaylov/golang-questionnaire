package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Result struct {
	gorm.Model
	ResultId  uuid.UUID  `json:"resultId"`
	Title  string  `json:"title"`
	Questions []Question `json:"questions" gorm:"-"`
}
