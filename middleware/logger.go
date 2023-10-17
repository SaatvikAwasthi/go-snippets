package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		startTime := time.Now()
		context.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := context.Request.Method
		reqUri := context.Request.RequestURI
		statusCode := context.Writer.Status()
		clientIP := context.ClientIP()
		log.Printf("| %3d | %13v | %15s | %s | %s | %v",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
			context.Errors,
		)
	}
}
