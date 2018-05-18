package pdfGenerator

import (
	"bytes"
	"fmt"
	"golang-questionnaire/app/models"
	"log"
	"os"
	"path"
	"time"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/labstack/echo"
)

var (
	pdfPath = path.Join(".", "pdf")
)

type (
	IPdf interface {
		GeneratePdf(context echo.Context, result *models.Result)
		GetFileInfo(id string) (pdfFilePath string, pdfName string)
	}
	Pdf struct {
		generator *wkhtmltopdf.PDFGenerator
	}
)

func NewPdf() IPdf {
	generator, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	gen := &Pdf{generator}

	return gen
}

func (gen *Pdf) GeneratePdf(c echo.Context, res *models.Result) {

	id := res.ResultId

	if _, err := os.Stat(pdfPath); os.IsNotExist(err) {
		os.Mkdir(pdfPath, 755)
	}

	pdfFilePath, _ := gen.GetFileInfo(id.String())

	if _, err := os.Stat(pdfFilePath); os.IsNotExist(err) {
		go gen.generateWkhtmlPdf(c, res, pdfFilePath)
	}
}

func (gen *Pdf) generateWkhtmlPdf(c echo.Context, res *models.Result, pdfFilePath string) {
	start := time.Now()

	log := c.Echo().Logger
	id := res.ResultId.String()

	buf := new(bytes.Buffer)
	err := c.Echo().Renderer.Render(buf, "results", res, c)
	if err != nil {
		log.Errorf("Error rendering the result %v", id)
	}

	// Add one page from an URL
	gen.generator.AddPage(wkhtmltopdf.NewPageReader(buf))

	// Create PDF document in internal buffer
	err = gen.generator.Create()
	if err != nil {
		log.Error(err)
	}

	// Write buffer contents to file on disk
	err = gen.generator.WriteFile(pdfFilePath)
	if err != nil {
		log.Error(err)
	}

	elapsed := time.Since(start)
	log.Printf("PDF generation took %s", elapsed)
}

func (gen *Pdf) GetFileInfo(id string) (pdfFilePath string, pdfName string) {
	pdfName = fmt.Sprintf("results-%s.pdf", id)
	pdfFilePath = path.Join(pdfPath, pdfName)
	return
}
