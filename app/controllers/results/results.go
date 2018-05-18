package results

import (
	"fmt"
	"golang-questionnaire/app/controllers/pdfGenerator"
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
		db     *gorm.DB
		pdfGen pdfGenerator.IPdf
	}
)

func (r *Results) SubmitResults(c echo.Context) error {

	uiid := uuid.New()

	res := &models.Result{
		ResultId: uiid,
		Title:    fmt.Sprintf("Title of result %v", uiid),
	}
	if err := c.Bind(res); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	r.db.Create(res)
	r.pdfGen.GeneratePdf(c, res)

	return c.JSON(http.StatusCreated, res)

}

func (r *Results) ViewResults(c echo.Context) error {
	uiid := c.Param("id")
	res := &models.Result{}
	var count int

	r.db.First(&res, "result_id = ?", uiid).Count(&count)
	if count == 0 {
		c.Logger().Errorf("No results with id %v", uiid)
		return c.JSON(http.StatusNotFound, `{"not found":true}`)
	}

	return c.Render(http.StatusOK, "results", &res)
}

func (r *Results) GetResultsPdf(c echo.Context) error {
	id := c.Param("id")
	pdfFilePath, pdfName := r.pdfGen.GetFileInfo(id)

	if _, err := os.Stat(pdfFilePath); os.IsNotExist(err) {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.Attachment(pdfFilePath, pdfName)
}

func NewResults(db *gorm.DB) IResults {
	results := &Results{
		db,
		pdfGenerator.NewPdf(),
	}

	return results
}

/*var instance IResults
var once *sync.Once

func GetResultsInstance(db *gorm.DB) IResults {
	once.Do(func() {
		instance = NewResults(db)
	})
	return instance
}*/
