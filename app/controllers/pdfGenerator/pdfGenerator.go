package pdfGenerator

import (
	"bytes"
	"errors"
	"fmt"
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
	pdfPath = path.Join(".", "pdf")
)

type (
	IPdfGenerator interface {
		Args() []string
		//AddPage(p *wkhtmltopdf.PageReader)
	}
	IPdf interface {
		GeneratePdf(context echo.Context, result *models.Result) error
		GetFileInfo(id string) (pdfFilePath string, pdfName string)
		GenerateWkhtmlPdf(c echo.Context, res *models.Result, pdfFilePath string) error
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

func (gen *Pdf) GeneratePdf(c echo.Context, res *models.Result) error {

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

func (gen *Pdf) GenerateWkhtmlPdf(c echo.Context, res *models.Result, pdfFilePath string) (err error) {
	start := time.Now()

	log := c.Echo().Logger
	id := res.ResultId.String()

	buf := new(bytes.Buffer)
	err = c.Echo().Renderer.Render(buf, "results", res, c)
	if err != nil {
		log.Errorf("Error rendering the result %v", id)
		return
	}

	// Add one page from an URL
	gen.generator.AddPage(wkhtmltopdf.NewPageReader(buf))

	// Create PDF document in internal buffer
	err = gen.generator.Create()
	if err != nil {
		log.Error(err)
		return
	}

	// Write buffer contents to file on disk
	err = gen.generator.WriteFile(pdfFilePath)
	if err != nil {
		log.Error(err)
		return
	}

	elapsed := time.Since(start)
	log.Printf("PDF generation took %s", elapsed)

	return nil
}

func (gen *Pdf) GetFileInfo(id string) (pdfFilePath string, pdfName string) {
	pdfName = fmt.Sprintf("results-%s.pdf", id)
	pdfFilePath = path.Join(pdfPath, pdfName)
	return
}
