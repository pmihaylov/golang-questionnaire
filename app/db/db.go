package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/pmihaylov/golang-questionnaire/app/models"
)

/**
user: root
pass: p0eGTfZPG4Ug

db: questionnaire

host: postgres-ft.cxrosmahhi34.eu-central-1.rds.amazonaws.com
port: 5432
*/

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open("postgres", "host=postgres-ft.cxrosmahhi34.eu-central-1.rds.amazonaws.com port=5432 user=root dbname=questionnaire password=p0eGTfZPG4Ug")
	if err != nil {
		panic("Failed to connect database")
	}

	// Migrate the schema
	DB.AutoMigrate(&models.Result{})
	DB.AutoMigrate(&models.Question{})
}
