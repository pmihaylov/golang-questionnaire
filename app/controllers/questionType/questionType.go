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

	controller.DB.Create(item)

	return c.NoContent(http.StatusCreated)
}

func (controller *QuestionTypeController) Read(c echo.Context) error {
	id := c.Param("id")
	item := new(models.QuestionType)

	controller.DB.First(&item, "id = ?", id)

	if item.ID == 0 {
		return controllers.HttpNotFound(c)
	}

	return c.JSONPretty(http.StatusOK, &item, " ")
}

func (controller *QuestionTypeController) List(c echo.Context) error {
	items := new([]models.QuestionType)

	controller.DB.Find(&items)

	if len(*items) == 0 {
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
