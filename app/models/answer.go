package models

// Answer model
type Answer struct {
	BaseModel
	QuestionID int    `gorm:"NOT NULL"`
	Value      string `gorm:"NOT NULL;Type:TEXT"`
}
