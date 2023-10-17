package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-snips/constants"
	handler2 "go-snips/handler"
	"go-snips/models"
	"go-snips/router"
	"go-snips/service"
	"go-snips/template"
	"go-snips/utils"
	"log"
	"net/http"
)

func Server(port string) {
	defer utils.RecoverPanic()

	url, routerHandler := router.SetupRouter(port)
	routerHandler = getHandlers(routerHandler)

	err := http.ListenAndServe(url, routerHandler)
	if err != nil {
		log.Fatalf("Error is setting up router. %e", err)
	}
}

func getHandlers(router *gin.Engine) *gin.Engine {
	router.NoRoute(func(c *gin.Context) {
		notRouteError := models.Error{
			Message: "Invalid Error",
			Error:   fmt.Errorf("no route exists for : %s", c.Request.URL),
			Code:    http.StatusNotFound,
		}
		c.JSON(http.StatusNotFound, notRouteError)
	})

	serv := getService()
	temp := getTemplate()
	handler := handler2.New(serv, temp)

	router.Handle(http.MethodGet, fmt.Sprintf("/greetHTML/:%s/:%s", constants.ParamFirstName, constants.ParamLastName), handler.GetHTML)
	router.Handle(http.MethodGet, fmt.Sprintf("/greet/:%s/:%s", constants.ParamFirstName, constants.ParamLastName), handler.Get)

	return router
}

func getService() service.Service {
	return service.New()
}

func getTemplate() template.Template {
	return template.New()
}
