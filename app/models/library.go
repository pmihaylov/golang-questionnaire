package models

import (
	"github.com/jinzhu/gorm"
)

type Library struct {
	gorm.Model
	Name string
}
