package models

// Identification model
type Identification struct {
	BaseModel
	Value          string `gorm:"NOT NULL"`
	Questionnaires []Questionnaire
}
