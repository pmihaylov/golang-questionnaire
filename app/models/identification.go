package models

type Identification struct {
	BaseModel
	Value string `gorm:"NOT NULL"`
}
