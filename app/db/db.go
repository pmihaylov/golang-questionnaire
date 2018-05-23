package db

import (
	"fmt"
	"golang-questionnaire/app/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"

	"github.com/labstack/echo"
)

func getConnectionString() string {
	dbConf := viper.GetStringMapString("db.postgresIreland")

	connectionString := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s",
		dbConf["host"],
		dbConf["port"],
		dbConf["db"],
		dbConf["user"],
		dbConf["password"],
	)

	return connectionString
}

func runAutoMigrations(db *gorm.DB) {
	// Migrate the schema
	db.AutoMigrate(
		&models.Answer{},
		&models.Identification{},
		&models.Library{},
		&models.Questionnaire{},
		&models.QuestionnaireNode{},
		&models.Question{},
		&models.QuestionType{},
	)

	db.Model(&models.Answer{}).AddForeignKey("question_id", "questions(id)", "NO ACTION", "NO ACTION")

	db.Model(&models.Question{}).AddForeignKey("library_id", "libraries(id)", "NO ACTION", "NO ACTION")
	db.Model(&models.Question{}).AddForeignKey("question_type_id", "question_types(id)", "NO ACTION", "NO ACTION")

	db.Model(&models.Questionnaire{}).AddForeignKey("library_id", "libraries(id)", "NO ACTION", "NO ACTION")
	db.Model(&models.Questionnaire{}).AddForeignKey("identification_id", "identifications(id)", "NO ACTION", "NO ACTION")
	// db.Model(&models.Questionnaire{}).AddForeignKey("entry_node_id", "questionnaire_nodes(id)", "NO ACTION", "NO ACTION")

	db.Model(&models.QuestionnaireNode{}).AddForeignKey("questionnaire_id", "questionnaires(id)", "NO ACTION", "NO ACTION")
	// db.Model(&models.QuestionnaireNode{}).AddForeignKey("parent_node_id", "questionnaire_nodes(id)", "NO ACTION", "NO ACTION")
	db.Model(&models.QuestionnaireNode{}).AddForeignKey("answer_id", "answers(id)", "NO ACTION", "NO ACTION")
}

func Init(server *echo.Echo) *gorm.DB {

	connectionString := getConnectionString()

	db, err := gorm.Open("postgres", connectionString)

	// db.LogMode(true)

	if err != nil {
		server.Logger.Fatalf("failed to connect database: %v", err)
	}

	runAutoMigrations(db)

	return db
}
