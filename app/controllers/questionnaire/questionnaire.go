package questionnaire

import (
	"bytes"
	"errors"
	"fmt"
	"golang-questionnaire/app/controllers"
	"golang-questionnaire/app/controllers/base"
	"golang-questionnaire/app/controllers/pdfGenerator"
	"golang-questionnaire/app/models"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type (
	IQuestionnaireController interface {
		base.IController
		View(c echo.Context) error
		Pdf(c echo.Context) error
	}
	QuestionnaireController struct {
		base.Controller
		pdfGen pdfGenerator.IPdf
	}
)

func (controller *QuestionnaireController) Create(c echo.Context) error {

	item := new(models.Questionnaire)

	if err := c.Bind(item); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := controller.DB.Create(item).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusCreated)
}

func (controller *QuestionnaireController) getQuestionnaire(id string) (*models.Questionnaire, error) {
	item := new(models.Questionnaire)

	if controller.DB.
		Preload("QuestionnaireNodes").
		Preload("Library").
		First(&item, "id = ?", id).RecordNotFound() {
		return item, errors.New("questionnaire record not found")
	}

	return item, nil
}

func (controller *QuestionnaireController) Read(c echo.Context) error {
	id := c.Param("id")

	item, err := controller.getQuestionnaire(id)
	if err != nil {
		return controllers.HttpNotFound(c)
	}

	return c.JSON(http.StatusOK, item)
}

func (controller *QuestionnaireController) View(c echo.Context) error {
	id := c.Param("id")

	item, err := controller.getQuestionnaire(id)
	if err != nil {
		return controllers.HttpNotFound(c)
	}

	return c.Render(http.StatusOK, "questionnaire", item)
}

func (controller *QuestionnaireController) Pdf(c echo.Context) error {
	id := c.Param("id")

	item, err := controller.getQuestionnaire(id)
	if err != nil {
		return controllers.HttpNotFound(c)
	}

	pdfFilePath, pdfName := controller.pdfGen.GetFileInfo(fmt.Sprintf("questionnaire-%v", item.ID))

	buf := new(bytes.Buffer)
	err = c.Echo().Renderer.Render(buf, "questionnaire", item, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	controller.pdfGen.HtmlToPdf(c, buf, pdfFilePath)

	if _, err := os.Stat(pdfFilePath); os.IsNotExist(err) {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.Attachment(pdfFilePath, pdfName)
}

func (controller *QuestionnaireController) List(c echo.Context) error {
	items := new([]models.Questionnaire)

	if controller.DB.
		Preload("QuestionnaireNodes").
		Preload("Library").
		Find(&items).RecordNotFound() {
		return controllers.HttpNotFound(c)
	}

	return c.JSON(http.StatusOK, &items)
}

func New(db *gorm.DB) IQuestionnaireController {
	controller := &QuestionnaireController{
		base.New(db),
		pdfGenerator.NewPdf(),
	}

	return controller
}
