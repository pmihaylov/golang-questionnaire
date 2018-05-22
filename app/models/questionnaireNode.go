package models

type QuestionnaireNode struct {
	BaseModel
	ParentNodeId    int `gorm:"NOT NULL"`
	QuestionnaireId int `gorm:"NOT NULL"`
	AnswerId        int
}
