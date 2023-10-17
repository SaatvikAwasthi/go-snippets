package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-snips/constants"
	"go-snips/models"
	"go-snips/service"
	"go-snips/template"
	"go-snips/utils"
	"net/http"
)

type handler struct {
	service  service.Service
	template template.Template
}

type Handler interface {
	Get(ctx *gin.Context)
	GetHTML(ctx *gin.Context)
	getParams(ctx *gin.Context) (string, string, error)
}

func New(serve service.Service, temp template.Template) Handler {
	return &handler{
		service:  serve,
		template: temp,
	}
}

func (h *handler) Get(ctx *gin.Context) {
	defer utils.RecoverPanic()

	ctx.Writer.Header().Set(constants.ContentTypeHeader, constants.ApplicationJson)

	fName, lName, err := h.getParams(ctx)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	message := h.service.GetMessage(fName, lName)

	greet := models.Greeting{Message: message}

	ctx.JSON(http.StatusOK, greet)
}

func (h *handler) GetHTML(ctx *gin.Context) {
	defer utils.RecoverPanic()

	fName, lName, err := h.getParams(ctx)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	message := h.service.GetMessage(fName, lName)

	err = h.template.GetHTML(message, ctx.Writer)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

}

func (h *handler) getParams(ctx *gin.Context) (string, string, error) {
	fName := ctx.Param(constants.ParamFirstName)
	if fName == constants.EmptyString {
		err := errors.New(constants.ErrorEmptyValue)
		return constants.EmptyString, constants.EmptyString, err
	}

	lName := ctx.Param(constants.ParamLastName)
	if lName == constants.EmptyString {
		err := errors.New(constants.ErrorEmptyValue)
		return constants.EmptyString, constants.EmptyString, err
	}

	return fName, lName, nil
}
