package restrouters

import (
	utils "Pay-AI/financial-transaction-server/Utils"
	authmech "Pay-AI/financial-transaction-server/restrouters/AuthMech"

	"github.com/gin-gonic/gin"
)

func Handler() {

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	t2 := r.Group("/v1/token")
	{
		t2.POST("/", authmech.GetTokenV1)

	}
	t := r.Group("/v2/token")
	{
		t.POST("/", authmech.GenerateToken)

	}
	r.Use(authmech.Middleware())

	p := r.Group("/ping")
	{
		p.GET("/", utils.GetPing)
	}
	r.Run(":8080")

}
