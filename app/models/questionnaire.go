package models

import (
	"github.com/jinzhu/gorm"
)

type Questionnaire struct {
	gorm.Model
	LibraryId           int    `json:"libraryId" gorm:"NOT NULL"`
	IdentificationId    string `json:"identificationId" gorm:"NOT NULL"`
	IdentificationValue string `json:"identificationValue"`
	Name                string
	EntryNodeId         int
}
