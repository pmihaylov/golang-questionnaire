package pdf

import (
	"bytes"
	"errors"
	"fmt"
	"golang-questionnaire/app/helpers"
	"golang-questionnaire/app/models"
	"log"
	"os"
	"path"
	"time"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/google/uuid"
	"github.com/labstack/echo"
)

var (
	pdfPath = path.Join("public", "pdf")
)

type (
	IPdfGenerator interface {
		GeneratePdf(context echo.Context, result *models.Result) error
		GetFileInfo(id string) (pdfFilePath string, pdfName string)
		GenerateWkhtmlPdf(c echo.Context, res *models.Result, pdfFilePath string) error
		HtmlToPdf(c echo.Context, buffer *bytes.Buffer, pdfFilePath string) error
	}
	PdfGenerator struct {
		generator *wkhtmltopdf.PDFGenerator
	}
)

func NewPdfGenerator() IPdfGenerator {
	generator, err := wkhtmltopdf.NewPDFGenerator()

	if err != nil {
		log.Fatal(err)
	}

	gen := &PdfGenerator{generator}

	return gen
}

func (gen *PdfGenerator) GeneratePdf(c echo.Context, res *models.Result) error {

	id := res.ResultId
	if id == uuid.Nil {
		return errors.New("result: empty id")
	}

	if _, err := os.Stat(pdfPath); os.IsNotExist(err) {
		os.Mkdir(pdfPath, 755)
	}

	pdfFilePath, _ := gen.GetFileInfo(id.String())

	if _, err := os.Stat(pdfFilePath); os.IsNotExist(err) {
		go gen.GenerateWkhtmlPdf(c, res, pdfFilePath)
	}

	return nil
}

func (gen *PdfGenerator) GenerateWkhtmlPdf(c echo.Context, res *models.Result, pdfFilePath string) (err error) {
	start := time.Now()

	id := res.ResultId.String()

	buf := new(bytes.Buffer)
	err = c.Echo().Renderer.Render(buf, "results", res, c)
	if err != nil {
		c.Logger().Error("Error rendering the result %v", id)
		return helpers.InternalServerError(c, err)
	}

	// Add one page from an URL
	gen.generator.AddPage(wkhtmltopdf.NewPageReader(buf))

	// Create PDF document in internal buffer
	err = gen.generator.Create()
	if err != nil {
		c.Logger().Error(err)
		return
	}

	// Write buffer contents to file on disk
	err = gen.generator.WriteFile(pdfFilePath)
	if err != nil {
		c.Logger().Error(err)
		return
	}

	elapsed := time.Since(start)
	c.Logger().Printf("PDF generation took %s", elapsed)

	return nil
}

func (gen *PdfGenerator) HtmlToPdf(c echo.Context, buffer *bytes.Buffer, pdfFilePath string) error {
	// start := time.Now()

	if _, err := os.Stat(pdfPath); os.IsNotExist(err) {
		os.Mkdir(pdfPath, 755)
	}

	gen.generator.AddPage(wkhtmltopdf.NewPageReader(buffer))
	err := gen.generator.Create()
	if err != nil {
		c.Logger().Error(err)
		return helpers.InternalServerError(c, err)
	}

	// Write buffer contents to file on disk
	err = gen.generator.WriteFile(pdfFilePath)
	if err != nil {
		return helpers.InternalServerError(c, err)
	}

	// elapsed := time.Since(start)
	// log.Printf("PDF generation took %s", elapsed)

	return nil
}

func (gen *PdfGenerator) GetFileInfo(name string) (pdfFilePath string, pdfName string) {
	pdfName = fmt.Sprintf("%s.pdf", name)
	pdfFilePath = path.Join(pdfPath, pdfName)
	return
}
