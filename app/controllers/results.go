package controllers

import (
	"fmt"
	"golang-questionnaire/app/models"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type (
	IResults interface {
		SubmitResults(c echo.Context) error
		ViewResults(c echo.Context) error
		GetResultsPdf(c echo.Context) error
	}
	Results struct {
		DB     *gorm.DB
		PdfGen IPdfGenerator
	}
)

func (r *Results) SubmitResults(c echo.Context) error {

	uiid := uuid.New()

	res := &models.Result{
		ResultId: uiid,
		Title:    fmt.Sprintf("Title of result %v", uiid),
	}
	if err := c.Bind(res); err != nil {
		return err
	}

	r.DB.Create(res)
	r.PdfGen.GeneratePdf(c, res)

	return c.JSON(http.StatusCreated, res)

}

func (r *Results) ViewResults(c echo.Context) error {
	uiid := c.Param("id")
	res := &models.Result{}
	var count int

	r.DB.First(&res, "result_id = ?", uiid).Count(&count)
	if count == 0 {
		c.Logger().Errorf("No results with id %v", uiid)
	}

	return c.Render(http.StatusOK, "results", &res)
}

func (r *Results) GetResultsPdf(c echo.Context) error {
	id := c.Param("id")
	pdfFilePath, pdfName := r.PdfGen.GetFile(id)

	if _, err := os.Stat(pdfFilePath); os.IsNotExist(err) {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.Attachment(pdfFilePath, pdfName)
}
