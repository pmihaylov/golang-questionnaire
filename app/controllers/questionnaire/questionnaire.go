package questionnaire

import (
	"fmt"
	"golang-questionnaire/app/controllers"
	"golang-questionnaire/app/controllers/base"
	"golang-questionnaire/app/models"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type (
	IQuestionnaireController interface {
		base.IController
	}
	QuestionnaireController struct {
		base.Controller
	}
)

func (controller *QuestionnaireController) Create(c echo.Context) error {

	item := new(models.Questionnaire)

	if err := c.Bind(item); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	controller.DB.Create(item)

	if controller.DB.Error != nil {
		fmt.Printf("%v", controller.DB.Error)
	}

	return c.NoContent(http.StatusCreated)
}

func (controller *QuestionnaireController) Read(c echo.Context) error {
	id := c.Param("id")
	item := new(models.Questionnaire)

	controller.DB.First(&item, "id = ?", id)

	if item.ID == 0 {
		return controllers.HttpNotFound(c)
	}

	return c.JSONPretty(http.StatusOK, &item, " ")
}

func (controller *QuestionnaireController) List(c echo.Context) error {
	items := new([]models.Questionnaire)

	controller.DB.Find(&items)

	if len(*items) == 0 {
		return controllers.HttpNotFound(c)
	}

	return c.JSON(http.StatusOK, &items)
}

func New(db *gorm.DB) IQuestionnaireController {
	controller := &QuestionnaireController{
		base.New(db),
	}

	return controller
}
