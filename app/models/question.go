package models

type Question struct {
	BaseModel
	LibraryId      int `gorm:"NOT NULL"`
	Required       bool
	Text           string `gorm:"Type:TEXT"`
	QuestionTypeId int    `gorm:"NOT NULL"`
}
