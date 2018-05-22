package models

type QuestionType struct {
	BaseModel
	Name string `gorm:"NOT NULL"`
}
