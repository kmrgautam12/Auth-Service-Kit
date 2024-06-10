package authmech

import (
	utils "Pay-AI/financial-transaction-server/Utils"

	"github.com/gin-gonic/gin"
)

func AuthenticationMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorize"})
			return
		}
		_, pub, err := ParsePublicPrivateKey()
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorize"})
			return
		}

		valid := utils.VerifyRSAJWTToken(pub, token)
		if valid {
			c.Next()
			return
		}
		c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorize"})
		return

	}

}

func CompressResponseV1(c *gin.Context, response interface{}) {
	if c.Header("Accept", "*/*") == true {

	}

}
