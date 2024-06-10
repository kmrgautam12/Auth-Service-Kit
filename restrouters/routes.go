package restrouters

import (
	utils "Pay-AI/financial-transaction-server/Utils"
	authmech "Pay-AI/financial-transaction-server/restrouters/AuthMech"
	"log"

	"github.com/gin-gonic/gin"
)

var logger = utils.Logger

func Handler() {

	r := gin.Default()
	r.Use(log.Logger)
	r.Use(authmech.LoggingMiddleware())
	r.Use(authmech.RecoveryMiddleware())
	t2 := r.Group("/v1/token")
	{
		t2.POST("/", authmech.GetTokenV1)

	}
	t := r.Group("/v2/token")
	{
		t.POST("/", authmech.GenerateToken)

	}
	r.Use(authmech.AuthenticationMiddleware())

	p := r.Group("/ping")
	{
		p.GET("/", utils.GetPing)
	}
	r.Run(":8080")

}
