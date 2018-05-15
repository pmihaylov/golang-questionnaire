package results

import (
	"bytes"
	"fmt"
	"golang-questionnaire/app/db"
	"golang-questionnaire/app/models"
	"net/http"
	"path"

	"github.com/google/uuid"
	"github.com/jung-kurt/gofpdf"
	"github.com/labstack/echo"
	"os"
)

var pdfPath = path.Join(".", "pdf")

func SubmitResults(c echo.Context) error {

	uiid := uuid.New()

	res := &models.Result{
		ResultId: uiid,
		Title:    fmt.Sprintf("Title of result %v", uiid),
	}
	if err := c.Bind(res); err != nil {
		return err
	}

	storeResults(c, res)

	return c.JSON(http.StatusCreated, res)

}

func storeResults(c echo.Context, res *models.Result) error {
	db.DB.Create(res)

	id := res.ResultId

	if _, err := os.Stat(pdfPath); os.IsNotExist(err) {
		os.Mkdir(pdfPath, 755)
	}

	pdfFilePath, _ := getFilePath(id.String())

	if _, err := os.Stat(pdfFilePath); os.IsNotExist(err) {
		buf := new(bytes.Buffer)
		err := c.Echo().Renderer.Render(buf, "results", res, c)
		if err != nil {
			c.Echo().Logger.Errorf("Error rendering the result %v", id)
		}

		pdf := gofpdf.New("P", "mm", "A4", "")
		pdf.AddPage()
		pdf.SetLeftMargin(45)
		pdf.SetFontSize(14)
		_, lineHt := pdf.GetFontSize()
		html := pdf.HTMLBasicNew()
		html.Write(lineHt, buf.String())

		err = pdf.OutputFileAndClose(pdfFilePath)
		if err != nil {
			c.Echo().Logger.Errorf("Error generating pdf for the result %v: %v", id, err)
		}
	}

	return nil
}

func getResults(uiid string, c echo.Context) (res models.Result, err error) {
	id, err := uuid.Parse(uiid)

	if err != nil {
		c.Logger().Printf("Wrong result id %v", uiid)
	}

	db.DB.First(&res, "result_id = ?", id)
	return
}

func ViewResult(c echo.Context) error {
	id := c.Param("id")

	res, err := getResults(id, c)
	if err != nil {
		c.Logger().Errorf("Error getting the result %v", id)
	}

	return c.JSON(http.StatusOK, res)
}

func GetResultsPdf(c echo.Context) error {
	id := c.Param("id")
	pdfFilePath, pdfName := getFilePath(id)

	if _, err := os.Stat(pdfFilePath); os.IsNotExist(err) {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.Attachment(pdfFilePath, pdfName)
}

func getFilePath(id string) (pdfFilePath string, pdfName string) {
	pdfName = fmt.Sprintf("results-%s.pdf", id)
	pdfFilePath = path.Join(pdfPath, pdfName)
	return
}
