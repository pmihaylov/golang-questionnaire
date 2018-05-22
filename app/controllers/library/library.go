package library

import (
	"golang-questionnaire/app/controllers"
	"golang-questionnaire/app/controllers/base"
	"golang-questionnaire/app/models"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type (
	ILibraryController interface {
		base.IController
	}
	LibraryController struct {
		base.Controller
	}
)

func (controller *LibraryController) Create(c echo.Context) error {
	item := new(models.Library)
	if err := c.Bind(item); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	controller.DB.Create(item)

	return c.NoContent(http.StatusCreated)
}

func (controller *LibraryController) Read(c echo.Context) error {
	id := c.Param("id")
	items := new(models.Library)

	controller.DB.First(&items, "id = ?", id)

	if items.ID == 0 {
		return controllers.HttpNotFound(c)
	}

	return c.JSONPretty(http.StatusOK, &items, " ")
}

func (controller *LibraryController) List(c echo.Context) error {
	items := new([]models.Library)

	controller.DB.Find(&items)

	if len(*items) == 0 {
		return controllers.HttpNotFound(c)
	}

	return c.JSON(http.StatusOK, &items)
}

func New(db *gorm.DB) ILibraryController {
	controller := &LibraryController{
		base.New(db),
	}

	return controller
}
