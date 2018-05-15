package routes

import (
	"net/http"

	"github.com/labstack/echo"

	"golang-questionnaire/app/controllers/results"
)

func Init(server *echo.Echo) {

	server.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", "GO Standalone")
	})

	// Results
	server.POST("/submit", results.SubmitResults)

	server.GET("/result/:id", results.ViewResult)
	server.GET("/pdf/:id", results.GetResultsPdf)
}
