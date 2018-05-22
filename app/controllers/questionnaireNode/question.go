package questionnaireNode

import (
	"golang-questionnaire/app/controllers"
	"golang-questionnaire/app/controllers/base"
	"golang-questionnaire/app/models"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type (
	IQuestionnaireNodeController interface {
		base.IController
	}
	QuestionnaireNodeController struct {
		base.Controller
	}
)

func (controller *QuestionnaireNodeController) Create(c echo.Context) error {

	items := new(models.QuestionnaireNode)
	if err := c.Bind(items); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	controller.DB.Create(items)

	return c.NoContent(http.StatusCreated)
}

func (controller *QuestionnaireNodeController) Read(c echo.Context) error {
	id := c.Param("id")
	items := new(models.QuestionnaireNode)

	controller.DB.First(&items, "id = ?", id)

	if items.ID == 0 {
		return controllers.HttpNotFound(c)
	}

	return c.JSONPretty(http.StatusOK, &items, " ")
}

func (controller *QuestionnaireNodeController) List(c echo.Context) error {
	items := new([]models.QuestionnaireNode)

	if len(*items) == 0 {
		return controllers.HttpNotFound(c)
	}

	controller.DB.Find(&items)

	return c.JSON(http.StatusOK, &items)
}

func New(db *gorm.DB) IQuestionnaireNodeController {
	controller := &QuestionnaireNodeController{
		base.New(db),
	}

	return controller
}
