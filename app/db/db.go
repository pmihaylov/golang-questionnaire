package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"fmt"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"golang-questionnaire/app/models"
)

var DB *gorm.DB

func Init(server *echo.Echo) {
	connectionString := getConnectionString()

	var err error
	DB, err = gorm.Open("postgres", connectionString)

	if err != nil {
		server.Logger.Fatalf("failed to connect database: %v", err)
	}

	// Migrate the schema
	DB.AutoMigrate(
		// &models.Result{},
		new(models.Library),
		new(models.QuestionType),
		new(models.Identification),
		new(models.Question),
		new(models.Questionnaire),
	)
}

func getConnectionString() (connectionString string) {
	dbConf := viper.GetStringMapString("db.postgresIreland")

	connectionString = fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s",
		dbConf["host"],
		dbConf["port"],
		dbConf["db"],
		dbConf["user"],
		dbConf["password"],
	)

	return
}
