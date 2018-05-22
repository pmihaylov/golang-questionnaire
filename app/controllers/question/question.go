package question

import (
	"golang-questionnaire/app/controllers"
	"golang-questionnaire/app/controllers/base"
	"golang-questionnaire/app/models"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type (
	IQuestionController interface {
		base.IController
	}
	QuestionController struct {
		base.Controller
	}
)

func (controller *QuestionController) Create(c echo.Context) error {

	item := new(models.Question)
	if err := c.Bind(item); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	controller.DB.Create(item)

	return c.NoContent(http.StatusCreated)
}

func (controller *QuestionController) Read(c echo.Context) error {
	id := c.Param("id")
	item := new(models.Question)

	controller.DB.First(&item, "id = ?", id)

	if item.ID == 0 {
		return controllers.HttpNotFound(c)
	}

	return c.JSONPretty(http.StatusOK, &item, " ")
}

func (controller *QuestionController) List(c echo.Context) error {
	items := new([]models.Question)

	if len(*items) == 0 {
		return controllers.HttpNotFound(c)
	}

	controller.DB.Find(&items)

	return c.JSON(http.StatusOK, &items)
}

func New(db *gorm.DB) IQuestionController {
	controller := &QuestionController{
		base.New(db),
	}

	return controller
}
