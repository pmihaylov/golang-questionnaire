package pdf

import (
	"encoding/json"
	"golang-questionnaire/app/models"
	"html/template"
	"io"
	"net/http/httptest"
	"os"
	"path"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	resultJSON = `{ "id": 28, "created_at": "2018-05-15 11:33:30.552288", "updated_at": "2018-05-15 11:33:30.552288", "deleted_at": null, "result_id": "09364436-18f4-4b14-9f4b-dbfa216a1831", "title": "Title of result 09364436-18f4-4b14-9f4b-dbfa216a1831" }`
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func getContext(id string) echo.Context {
	e := echo.New()

	templateRenderer := &Template{
		templates: template.Must(template.ParseGlob(path.Join("..", "..", "views", "*.html"))),
	}

	e.Renderer = templateRenderer
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	c.SetPath("/pdf/:id")
	c.SetParamNames("id")
	c.SetParamValues(id)

	return c
}

func TestPdf_GeneratePdf(t *testing.T) {
	pdf := NewPdfGenerator()
	assert := assert.New(t)

	result := &models.Result{}
	json.Unmarshal([]byte(resultJSON), result)
	result.ResultID = uuid.New()

	c := getContext(result.ResultID.String())

	assert.NoError(pdf.GeneratePdf(c, result))
}

func TestPdf_GeneratePdfFailure(t *testing.T) {
	pdf := NewPdfGenerator()
	assert := assert.New(t)

	result := &models.Result{}
	json.Unmarshal([]byte(resultJSON), result)
	result.ResultID = uuid.Nil

	c := getContext(result.ResultID.String())

	assert.Error(pdf.GeneratePdf(c, result))
}

func TestPdf_GenerateWkhtmlPdf(t *testing.T) {
	pdfPath = path.Join("test_pdfs")

	pdf := NewPdfGenerator()
	assert := assert.New(t)

	result := &models.Result{}
	json.Unmarshal([]byte(resultJSON), result)
	result.ResultID = uuid.New()

	c := getContext(result.ResultID.String())

	pdfFilePath, pdfName := pdf.GetFileInfo(result.ResultID.String())

	assert.IsType(*new(string), pdfFilePath)
	assert.NotEmpty(pdfFilePath)

	assert.IsType(*new(string), pdfName)
	assert.NotEmpty(pdfName)

	if _, err := os.Stat(pdfPath); os.IsNotExist(err) {
		os.MkdirAll(pdfPath, 755)
	}

	if assert.NoError(pdf.GenerateWkhtmlPdf(c, result, pdfFilePath)) {
		_, err := os.Stat(pdfFilePath)
		assert.NoError(err)
	}

	assert.NoError(os.RemoveAll(pdfPath))
}