package answer

import (
	"golang-questionnaire/app/controllers/base"
	"golang-questionnaire/app/helpers"
	"golang-questionnaire/app/models"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type (
	IController interface {
		base.IController
	}
	Controller struct {
		base.Controller
	}
)

func (controller *Controller) Create(c echo.Context) error {

	item := new(models.Answer)
	if err := c.Bind(item); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := controller.DB.Create(item).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusCreated)
}

func (controller *Controller) Read(c echo.Context) error {
	id := c.Param("id")
	item := new(models.Answer)

	if controller.DB.First(&item, "id = ?", id).RecordNotFound() {
		return helpers.HttpNotFound(c)
	}

	return c.JSONPretty(http.StatusOK, &item, " ")
}

func (controller *Controller) List(c echo.Context) error {
	items := new([]models.Answer)

	if controller.DB.Find(&items).RecordNotFound() {
		return helpers.HttpNotFound(c)
	}

	return c.JSON(http.StatusOK, &items)
}

func New(db *gorm.DB) IController {
	controller := &Controller{
		base.New(db),
	}

	return controller
}
