package routes

import (
	"golang-questionnaire/app/controllers"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func Init(server *echo.Echo, db *gorm.DB) {

	results := &controllers.Results{
		DB:     db,
		PdfGen: &controllers.PdfGenerator{},
	}

	server.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", "GO Standalone")
	})

	// Results
	server.POST("/submit", results.SubmitResults)

	server.GET("/result/:id", results.ViewResults)
	server.GET("/pdf/:id", results.GetResultsPdf)
}
