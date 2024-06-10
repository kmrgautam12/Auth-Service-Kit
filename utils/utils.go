package utils

import (
	"crypto/rsa"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GetPing(c *gin.Context) {

	c.JSON(200, gin.H{"status": "healthy"})
}

func MustParseUrl(u string) *url.URL {

	url, err := url.Parse(u)
	if err != nil {
		return nil
	}
	return url
}

func VerifyRSAJWTToken(pub *rsa.PublicKey, token string) bool {

	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return "Unexpected method is called", nil
		}
		return pub, nil

	})
	if err != nil {
		return false
	}

	if t.Valid {
		return true
	}
	return false

}
