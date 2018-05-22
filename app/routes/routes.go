package routes

import (
	"golang-questionnaire/app/controllers/answer"
	"golang-questionnaire/app/controllers/identification"
	"golang-questionnaire/app/controllers/library"
	"golang-questionnaire/app/controllers/question"
	"golang-questionnaire/app/controllers/questionType"
	"golang-questionnaire/app/controllers/questionnaire"
	"golang-questionnaire/app/controllers/questionnaireNode"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func Init(server *echo.Echo, db *gorm.DB) {
	server.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", "Go standalone P.O.C")
	})

	answerController := answer.New(db)
	identificationController := identification.New(db)
	libraryController := library.New(db)
	questionController := question.New(db)
	questionTypeController := questionType.New(db)
	questionnaireController := questionnaire.New(db)
	questionnaireNodeController := questionnaireNode.New(db)

	// Library
	server.POST("/lib", libraryController.Create)
	server.GET("/lib", libraryController.List)
	server.GET("/lib/:id", libraryController.Read)

	// Question Types
	server.POST("/question-type", questionTypeController.Create)
	server.GET("/question-type", questionTypeController.List)
	server.GET("/question-type/:id", questionTypeController.Read)

	// Identification
	server.POST("/id", identificationController.Create)
	server.GET("/id", identificationController.List)
	server.GET("/id/:id", identificationController.Read)

	// Questions
	server.POST("/question", questionController.Create)
	server.GET("/question", questionController.List)
	server.GET("/question/:id", questionController.Read)

	// Answers
	server.POST("/answer", answerController.Create)

	// Questionnaire
	server.POST("/questionnaire", questionnaireController.Create)
	server.GET("/questionnaire", questionnaireController.List)
	server.GET("/questionnaire/:id", questionnaireController.Read)

	// QuestionnaireNode
	server.POST("/node", questionnaireNodeController.Create)
	server.GET("/node", questionnaireNodeController.List)
	server.GET("/node/:id", questionnaireNodeController.Read)
}
