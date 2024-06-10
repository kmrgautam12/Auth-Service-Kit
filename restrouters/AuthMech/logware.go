package authmech

import (
	utils "Pay-AI/financial-transaction-server/Utils"
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
