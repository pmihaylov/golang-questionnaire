package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Result struct {
	gorm.Model
	ResultId  uuid.UUID  `json:"resultId" gorm:"type:varchar(36)"`
	Title  string  `json:"title"`
	Questions []Question `json:"questions" gorm:"-"`
}
