package authmech

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(c *gin.Context) {

	pri, _, err := GeneratePublicPrivateKeys()
	if err != nil {
		c.JSON(400, gin.H{"message1": "Unauthorize"})
		return
	}
	t := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"name": "Pay-AI",
		"exp":  time.Now().Add(time.Minute * 25).Unix(),
	})
	tokenStr, err := t.SignedString(pri)
	if err != nil {
		c.JSON(400, gin.H{"message": "Token generation failed"})
		return
	}
	c.JSON(200, gin.H{"token": tokenStr})
	return
}
