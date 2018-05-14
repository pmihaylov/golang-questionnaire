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

	dbConf := viper.GetStringMapString("db.postgres")

	connectionString := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s",
		dbConf["host"],
		dbConf["port"],
		dbConf["db"],
		dbConf["user"],
		dbConf["password"],
	)

	var err error
	DB, err = gorm.Open("postgres", connectionString)
	if err != nil {
		panic("Failed to connect database")
	}

	// Migrate the schema
	DB.AutoMigrate(
		&models.Result{},
		&models.Question{},
	)
}
