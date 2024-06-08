package authmech

import (
	"Pay-AI/financial-transaction-server/constantservice"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CreateHs256Token() (string, error) {

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Audience": "local",
		"role":     "admin",
		"exp":      time.Now().Add(time.Hour * 1),
	},
	)
	saltSecretStr := constantservice.SecretString
	s := []byte(saltSecretStr["Secret"] + saltSecretStr["Salt"])
	token, err := t.SignedString(s)
	if err != nil {
		return "", err
	}
	return token, nil

}
func GetTokenV1(c *gin.Context) {
	token, err := CreateHs256Token()
	if err != nil {
		fmt.Println("Error:", err)
		c.JSON(400, gin.H{"message": "Unable to generate token"})
		return
	}
	c.JSON(200, gin.H{"token": token})
	return

}
