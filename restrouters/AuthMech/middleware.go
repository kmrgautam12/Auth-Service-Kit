package authmech

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Middleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		token := c.GetHeader("Authrization")
		if token == "" {
			c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorize"})
		}
		jwtToken, err := jwt.ParseWithClaims(token,
			jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
				return t, nil
			},
		)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"message": "Unauthoriza"})
		}
		if jwtToken.Valid {
			c.Next()
		}

	}

}
