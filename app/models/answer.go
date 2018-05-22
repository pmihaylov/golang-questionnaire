package models

type Answer struct {
	BaseModel
	QuestionId int    `gorm:"NOT NULL"`
	Value      string `gorm:"NOT NULL;Type:TEXT"`
}
