package authmech

import (
	utils "Pay-AI/financial-transaction-server/Utils"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var logger = utils.Logger

func LoggingMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		duration := time.Since(startTime)
		logger.Info(
			"Request completed",
			zap.String("Request method:", c.Request.Method),
			zap.String("Request Url", c.Request.RequestURI),
			zap.String("Remote address", c.ClientIP()),
			zap.Int("Response status", c.Writer.Status()),
			zap.Duration("Response time", duration),
		)

	}

}

func RecoveryMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("Panic:",
					zap.String(
						"RecoveryMiddlewareError", fmt.Sprint(err),
					))
				ctx.AbortWithStatusJSON(500, gin.H{"message": "Internal Server Error"})
				return
			}
		}()
		ctx.Next()
	}
}
