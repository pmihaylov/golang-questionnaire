package app

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
	"golang-questionnaire/app/db"
	"golang-questionnaire/app/routes"
	"html/template"
	"io"
	"path"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var Server *echo.Echo

func Config() {
	viper.SetConfigType("json")
	viper.SetConfigName("app-config")
	viper.AddConfigPath(path.Join(".", "app", "config"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func Init() {
	db.Init()
	defer db.DB.Close()

	templateRenderer := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	Server = echo.New()

	Server.Renderer = templateRenderer

	// Middleware
	Server.Use(middleware.Logger())
	Server.Use(middleware.Recover())

	routes.Init(Server)

	Server.Logger.Fatal(Server.Start(":8080"))
}
