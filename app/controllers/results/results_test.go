package results

import (
	"database/sql"
	"encoding/json"
	"golang-questionnaire/app/models"
	"html/template"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"path"
	"strings"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var (
	submitJson = `{ "questions": [ { "questionId": "6f6e200d-e4f7-4b90-aad6-cae937c0d044", "answers" :  [1,4] }, { "questionId": "0268b968-d96c-411f-8857-ae967d261df3", "answers" :  [2] }, { "questionId": "29e0506f-3321-4ea0-857d-1f026d8b57ca", "answers" :  ["Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vestibulum ac lacus in libero vulputate sagittis in quis odio. Quisque felis tellus, fringilla nec sapien non, porta commodo leo. In congue neque id eros elementum scelerisque. Nulla dolor libero, commodo in elit at, sodales fermentum ante. Proin augue enim, finibus in scelerisque sed, elementum non magna. Nam dui dui, consequat ut efficitur ut, euismod in leo. Proin in odio luctus neque eleifend hendrerit. Integer ipsum dolor, sagittis eu venenatis eget, sagittis id orci. Integer hendrerit dapibus erat, vel pretium metus ultricies sed. Nulla scelerisque mollis dictum. Nunc aliquet ultricies purus ac sodales. Aenean fringilla pharetra tellus, eget dignissim urna tristique ut."] } ] }`
	resultJson = ` { "id": 28, "created_at": "2018-05-15 11:33:30.552288", "updated_at": "2018-05-15 11:33:30.552288", "deleted_at": null, "result_id": "09364436-18f4-4b14-9f4b-dbfa216a1831", "title": "Title of result 09364436-18f4-4b14-9f4b-dbfa216a1831" } `
)

type MockPdfGenerator struct{}

func (gen *MockPdfGenerator) GeneratePdf(c echo.Context, res *models.Result) error       { return nil }
func (gen *MockPdfGenerator) GetFileInfo(id string) (pdfFilePath string, pdfName string) { return }
func (gen *MockPdfGenerator) GenerateWkhtmlPdf(c echo.Context, res *models.Result, pdfFilePath string) error {
	return nil
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newDB() (sqlmock.Sqlmock, *gorm.DB) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("can't create sqlmock: %s", err)
	}

	gormDB, gerr := gorm.Open("postgres", db)
	if gerr != nil {
		log.Fatalf("can't open gorm connection: %s", err)
	}
	gormDB.LogMode(true)

	return mock, gormDB.Set("gorm:update_column", true)
}

func getResultsWithMocks(db *gorm.DB) IResults {
	results := &Results{
		db:     db,
		pdfGen: &MockPdfGenerator{},
	}

	return results
}

func TestNewResults(t *testing.T) {
	_, db := newDB()
	results := NewResults(db)

	assert.IsType(t, &Results{}, results)
}

func TestResults_SubmitResults(t *testing.T) {
	mock, db := newDB()
	defer assert.NoError(t, mock.ExpectationsWereMet())
	e := echo.New()

	req := httptest.NewRequest(echo.POST, "/submit", strings.NewReader(submitJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	results := getResultsWithMocks(db)

	rows := sqlmock.NewRows([]string{"id"}).
		AddRow("1")

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT INTO \"results\" (.+)").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(rows)
	mock.ExpectCommit()

	// Assertions
	if assert.NoError(t, results.SubmitResults(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestResults_SubmitResultsError(t *testing.T) {
	mock, db := newDB()
	defer assert.NoError(t, mock.ExpectationsWereMet())
	e := echo.New()

	req := httptest.NewRequest(echo.POST, "/submit", strings.NewReader("malformed"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	results := getResultsWithMocks(db)

	rows := sqlmock.NewRows([]string{"id"}).
		AddRow("1")

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT INTO \"results\" (.+)").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(rows)
	mock.ExpectCommit()

	// Assertions
	if assert.NoError(t, results.SubmitResults(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestResults_ViewResults(t *testing.T) {
	mock, db := newDB()
	defer assert.NoError(t, mock.ExpectationsWereMet())
	e := echo.New()

	result := &models.Result{}
	json.Unmarshal([]byte(resultJson), result)

	templateRenderer := &Template{
		templates: template.Must(template.ParseGlob(path.Join("..", "..", "..", "public", "views", "*.html"))),
	}

	e.Renderer = templateRenderer

	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/result/:id")
	c.SetParamNames("id")
	c.SetParamValues(result.ResultId.String())

	results := getResultsWithMocks(db)

	rows := sqlmock.
		NewRows([]string{"id", "result_id"}).
		AddRow(rand.Int(), result.ResultId.String())

	countRows := sqlmock.
		NewRows([]string{"count"}).
		AddRow(1)

	mock.ExpectQuery("SELECT \\* FROM \"results\".*").WillReturnRows(rows)
	mock.ExpectQuery("SELECT count\\(\\*\\) FROM \"results\".*").WillReturnRows(countRows)

	if assert.NoError(t, results.ViewResults(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestResults_ViewResultsFailure(t *testing.T) {
	mock, db := newDB()
	defer assert.NoError(t, mock.ExpectationsWereMet())
	e := echo.New()

	result := &models.Result{}
	json.Unmarshal([]byte(resultJson), result)

	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/result/:id")
	c.SetParamNames("id")
	c.SetParamValues("")

	results := getResultsWithMocks(db)

	countRows := sqlmock.
		NewRows([]string{"count"}).
		AddRow(0)

	mock.ExpectQuery("SELECT \\* FROM \"results\".*").WillReturnError(sql.ErrNoRows)
	mock.ExpectQuery("SELECT count\\(\\*\\) FROM \"results\".*").WillReturnRows(countRows)

	if assert.NoError(t, results.ViewResults(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
	}
}

func TestResults_GetResultsPdf(t *testing.T) {
	_, db := newDB()
	e := echo.New()
	results := getResultsWithMocks(db)

	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/pdf/:id")
	c.SetParamNames("id")
	c.SetParamValues("")

	if assert.NoError(t, results.GetResultsPdf(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
	}
}
