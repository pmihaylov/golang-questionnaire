package answer

import (
	"golang-questionnaire/app/controllers"
	"golang-questionnaire/app/controllers/base"
	"golang-questionnaire/app/models"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type (
	IAnswerController interface {
		base.IController
	}
	AnswerController struct {
		base.Controller
	}
)

func (controller *AnswerController) Create(c echo.Context) error {

	item := new(models.Answer)
	if err := c.Bind(item); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	controller.DB.Create(item)

	return c.NoContent(http.StatusCreated)
}

func (controller *AnswerController) Read(c echo.Context) error {
	id := c.Param("id")
	item := new(models.Answer)

	controller.DB.First(&item, "id = ?", id)

	if item.ID == 0 {
		return controllers.HttpNotFound(c)
	}

	return c.JSONPretty(http.StatusOK, &item, " ")
}

func (controller *AnswerController) List(c echo.Context) error {
	items := new([]models.Answer)

	if len(*items) == 0 {
		return controllers.HttpNotFound(c)
	}

	controller.DB.Find(&items)

	return c.JSON(http.StatusOK, &items)
}

func New(db *gorm.DB) IAnswerController {
	controller := &AnswerController{
		base.New(db),
	}

	return controller
}
