package models

type Question struct {
	BaseModel
	Required       bool
	Text           string `gorm:"Type:TEXT"`
	QuestionTypeID int    `gorm:"NOT NULL"`
	LibraryID      int    `gorm:"NOT NULL"`

	QuestionType   QuestionType
	Library        Library
	Answers        []Answer
}
