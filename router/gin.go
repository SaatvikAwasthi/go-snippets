package router

import (
	"github.com/gin-gonic/gin"
	"go-snips/constants"
	"go-snips/middleware"
	"log"
)

func SetupRouter(port string) (string, *gin.Engine) {
	routerUrl := constants.KeySeperator + port

	router := newRouter(constants.EmptyString)
	log.Printf("Service loaded successfully")

	return routerUrl, router
}

func newRouter(appMode string) *gin.Engine {
	switch appMode {
	case constants.Prod:
		gin.SetMode(gin.ReleaseMode)
		break
	case constants.Test:
		gin.SetMode(gin.TestMode)
		break
	default:
		gin.SetMode(gin.DebugMode)
		break
	}
	router := gin.Default()
	router.Use(middleware.LoggerMiddleware())
	router.Use(middleware.EnableCors())
	return router
}
