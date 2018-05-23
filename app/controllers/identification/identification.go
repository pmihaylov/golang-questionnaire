package identification

import (
	"golang-questionnaire/app/controllers"
	"golang-questionnaire/app/controllers/base"
	"golang-questionnaire/app/models"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type (
	IIdentificationController interface {
		base.IController
	}
	IdentificationController struct {
		base.Controller
	}
)

func (controller *IdentificationController) Create(c echo.Context) error {

	item := new(models.Identification)
	if err := c.Bind(item); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := controller.DB.Create(item).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusCreated)
}

func (controller *IdentificationController) Read(c echo.Context) error {
	id := c.Param("id")
	item := new(models.Identification)

	controller.DB.First(&item, "id = ?", id)

	if item.ID == 0 {
		return controllers.HttpNotFound(c)
	}

	return c.JSONPretty(http.StatusOK, &item, " ")
}

func (controller *IdentificationController) List(c echo.Context) error {
	items := new([]models.Identification)

	if controller.DB.Find(&items).RecordNotFound() {
		return controllers.HttpNotFound(c)
	}

	return c.JSON(http.StatusOK, &items)
}

func New(db *gorm.DB) IIdentificationController {
	controller := &IdentificationController{
		base.New(db),
	}

	return controller
}
