package helpers

import (
	"net/http"

	"github.com/labstack/echo"
)

func HttpNotFound(c echo.Context) error {
	return c.JSONBlob(http.StatusNotFound, []byte(`{"message": "Not Found"}`))
}

func InternalServerError(c echo.Context, err error) error {
	return c.JSON(http.StatusInternalServerError, err)
}
