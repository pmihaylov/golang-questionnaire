package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func HttpNotFound(c echo.Context) error {
	return c.JSONBlob(http.StatusNotFound, []byte(`{"message": "Not Found"}`))
}
