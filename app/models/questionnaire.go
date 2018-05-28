package models

// Questionnaire model
type Questionnaire struct {
	BaseModel
	IdentificationID int `gorm:"NOT NULL"`
	Name             string
	EntryNodeID      int
	LibraryID        int `gorm:"NOT NULL"`

	Library Library
	// Libraries          []*Library `gorm:"many2many:library_questionnaires"'`
	QuestionnaireNodes []QuestionnaireNode
	// !!! Example !!! EntryNodeID      pq.Int64Array `gorm:"Type:integer[]"`
}
