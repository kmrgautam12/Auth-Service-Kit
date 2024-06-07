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

	// auth := r.Group("/auth")
	// {
	// 	auth.Use(authmech.Middleware())
	// }
	t := r.Group("/v1/token")
	{
		t.POST("/", authmech.GenerateToken)

	}
	p := r.Group("/ping")
	{
		p.GET("/", utils.GetPing)
	}

}
