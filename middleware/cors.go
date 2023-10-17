package middleware

import (
	"go-snips/constants"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func EnableCors() gin.HandlerFunc {
	return func(context *gin.Context) {
		origin := context.Request.Header.Get("Origin")
		if len(origin) != 0 {
			context.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			context.Writer.Header().Set("Access-Control-Allow-Headers",
				strings.Join([]string{
					constants.Origin,
					constants.Accept,
					constants.ContentTypeHeader,
					constants.Authorization,
					constants.DateUsed,
					constants.XRequestedWith,
					constants.Cookie,
				}, ","))
			context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			context.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		}
		if context.Request.Method == http.MethodOptions {
			log.Printf("Error: Cors OPTIONS Mthord not allowed")
			return
		}
		context.Next()
	}
}
