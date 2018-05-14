package routes

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/pmihaylov/golang-questionnaire/app/controllers/results"
)

func Init(server *echo.Echo) {

	server.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", "GO Standalone")
	})

	// Results
	server.GET("/result/:id", results.ViewResult)
	server.POST("/submit", results.SubmitResults)
}
