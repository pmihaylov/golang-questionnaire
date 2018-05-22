package models

import (
	"github.com/lib/pq"
)

type Questionnaire struct {
	BaseModel
	LibraryId        int `gorm:"NOT NULL"`
	IdentificationId int `gorm:"NOT NULL"`
	Name             string
	EntryNodeId      pq.Int64Array `gorm:"Type:integer[]"`
}
