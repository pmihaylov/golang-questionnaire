package models

type Library struct {
	BaseModel
	Name string
	Questions []*Question `gorm:"many2many:library_questions"'`
	Questionnaires []*Questionnaire `gorm:"many2many:library_questionnaires"'`
}
