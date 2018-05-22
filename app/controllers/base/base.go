package base

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type (
	IController interface {
		Create(c echo.Context) error
		Read(c echo.Context) error
		// Update(c echo.Context) error
		// Delete(c echo.Context) error
		List(c echo.Context) error
	}
	Controller struct {
		DB *gorm.DB
	}
)

func New(db *gorm.DB) Controller {
	return Controller{
		db,
	}
}
