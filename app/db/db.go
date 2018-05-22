package db

import (
	"fmt"
	"golang-questionnaire/app/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"

	"github.com/labstack/echo"
)

var DB = new(gorm.DB)

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
		new(models.Answer),
		new(models.Identification),
		new(models.Library),
		new(models.Questionnaire),
		new(models.QuestionnaireNode),
		new(models.Question),
		new(models.QuestionType),
	)
}

func Init(server *echo.Echo) {
	connectionString := getConnectionString()

	db, err := gorm.Open("postgres", connectionString)

	if err != nil {
		server.Logger.Fatalf("failed to connect database: %v", err)
	}

	runAutoMigrations(db)

	DB = db
}
