package models

// QuestionType model
type QuestionType struct {
	BaseModel
	Name      string `gorm:"NOT NULL"`
	Questions []Question
}
