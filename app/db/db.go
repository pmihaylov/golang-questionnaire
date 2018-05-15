package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"fmt"
	"github.com/spf13/viper"
	"golang-questionnaire/app/models"
)

var DB *gorm.DB

func Init() {

	connectionString := getConnectionString()

	var err error
	DB, err = gorm.Open("postgres", connectionString)
	if err != nil {
		fmt.Errorf("failed to connect database: %v", err)
	}

	// Migrate the schema
	DB.AutoMigrate(
		&models.Result{},
		&models.Question{},
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
