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

	item := new(models.QuestionnaireNode)
	if err := c.Bind(item); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := controller.DB.Create(item).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusCreated)
}

func (controller *QuestionnaireNodeController) Read(c echo.Context) error {
	id := c.Param("id")
	item := new(models.QuestionnaireNode)

	if controller.DB.First(&item, "id = ?", id).RecordNotFound() {
		return controllers.HttpNotFound(c)
	}

	return c.JSONPretty(http.StatusOK, &item, " ")
}

func (controller *QuestionnaireNodeController) List(c echo.Context) error {
	items := new([]models.QuestionnaireNode)

	if controller.DB.Find(&items).RecordNotFound() {
		return controllers.HttpNotFound(c)
	}

	return c.JSON(http.StatusOK, &items)
}

func New(db *gorm.DB) IQuestionnaireNodeController {
	controller := &QuestionnaireNodeController{
		base.New(db),
	}

	return controller
}
