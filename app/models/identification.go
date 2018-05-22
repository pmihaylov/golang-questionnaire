package models

import (
	"github.com/jinzhu/gorm"
)

type Identification struct {
	gorm.Model
	Value string `gorm:"NOT NULL"`
}
