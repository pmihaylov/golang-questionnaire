package models

// QuestionnaireNode model
type QuestionnaireNode struct {
	BaseModel
	ParentNodeID    int `gorm:"NOT NULL"`
	QuestionnaireID int `gorm:"NOT NULL"`
	AnswerID        int
}
