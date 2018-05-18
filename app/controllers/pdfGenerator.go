package controllers

import (
	"bytes"
	"fmt"
	"golang-questionnaire/app/models"
	"os"
	"path"
	"time"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/jung-kurt/gofpdf"
	"github.com/labstack/echo"
)

var (
	pdfPath = path.Join(".", "pdf")
)

type (
	IPdfGenerator interface {
		GeneratePdf(context echo.Context, result *models.Result)
		GetFile(id string) (pdfFilePath string, pdfName string)
	}
	PdfGenerator struct{}
)

func (gen *PdfGenerator) GeneratePdf(c echo.Context, res *models.Result) {
	start := time.Now()

	id := res.ResultId

	if _, err := os.Stat(pdfPath); os.IsNotExist(err) {
		os.Mkdir(pdfPath, 755)
	}

	pdfFilePath, _ := gen.getFilePath(id.String())

	if _, err := os.Stat(pdfFilePath); os.IsNotExist(err) {
		go gen.generateWkhtmlPdf(c, res, pdfFilePath)
	}

	elapsed := time.Since(start)
	c.Echo().Logger.Printf("PDF generation took %s", elapsed)
}

func (gen *PdfGenerator) generateFpdf(c echo.Context, res *models.Result, pdfFilePath string) {
	id := res.ResultId.String()

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

func (gen *PdfGenerator) generateWkhtmlPdf(c echo.Context, res *models.Result, pdfFilePath string) {
	log := c.Echo().Logger
	id := res.ResultId.String()

	buf := new(bytes.Buffer)
	err := c.Echo().Renderer.Render(buf, "results", res, c)
	if err != nil {
		log.Errorf("Error rendering the result %v", id)
	}

	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Error(err)
	}

	// Add one page from an URL
	pdfg.AddPage(wkhtmltopdf.NewPageReader(buf))

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Error(err)
	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile(pdfFilePath)
	if err != nil {
		log.Error(err)
	}
}

func (gen *PdfGenerator) getFilePath(id string) (pdfFilePath string, pdfName string) {
	pdfName = fmt.Sprintf("results-%s.pdf", id)
	pdfFilePath = path.Join(pdfPath, pdfName)
	return
}

func (gen *PdfGenerator) GetFile(id string) (pdfFilePath string, pdfName string) {
	pdfName = fmt.Sprintf("results-%s.pdf", id)
	pdfFilePath = path.Join(pdfPath, pdfName)
	return
}
