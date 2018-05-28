package models

import (
	"github.com/jinzhu/gorm"
)

// BaseModel is a model which is an adaptor for Gorm fields e.g. contains ID, created_at, etc.
type BaseModel struct {
	gorm.Model
}
