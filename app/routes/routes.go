package routes

import (
	results "golang-questionnaire/app/controllers/results"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func Init(server *echo.Echo, db *gorm.DB) {
	server.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", "GO Standalone")
	})

	results := results.NewResults(db)

	// Results
	server.POST("/submit", results.SubmitResults)

	server.GET("/result/:id", results.ViewResults)
	server.GET("/pdf/:id", results.GetResultsPdf)
}
