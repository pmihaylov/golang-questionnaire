package questionType

import (
	"golang-questionnaire/app/controllers"
	"golang-questionnaire/app/controllers/base"
	"golang-questionnaire/app/models"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type (
	IQuestionTypeController interface {
		base.IController
	}
	QuestionTypeController struct {
		base.Controller
	}
)

func (controller *QuestionTypeController) Create(c echo.Context) error {

	item := new(models.QuestionType)
	if err := c.Bind(item); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := controller.DB.Create(item).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusCreated)
}

func (controller *QuestionTypeController) Read(c echo.Context) error {
	id := c.Param("id")
	item := new(models.QuestionType)

	if controller.DB.First(&item, "id = ?", id).RecordNotFound() {
		return controllers.HttpNotFound(c)
	}

	return c.JSONPretty(http.StatusOK, &item, " ")
}

func (controller *QuestionTypeController) List(c echo.Context) error {
	items := new([]models.QuestionType)

	if controller.DB.Find(&items).RecordNotFound() {
		return controllers.HttpNotFound(c)
	}

	return c.JSON(http.StatusOK, &items)
}

func New(db *gorm.DB) IQuestionTypeController {
	controller := &QuestionTypeController{
		base.New(db),
	}

	return controller
}
