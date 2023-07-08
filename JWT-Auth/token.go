package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type JWTClaims struct {
	UserName string
	Password string
	jwt.StandardClaims
}

var jwtKey = []byte("supersecretkey")

func GenerateTokenRS256(c *gin.Context) (string, error) {
	keyGenrateKey, _ := rsa.GenerateKey(rand.Reader, 2048)

	fmt.Println("key generated is ", keyGenrateKey)
	privateKeyPem, _ := os.Create("private_key.pem")
	_ = pem.Encode(privateKeyPem, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(keyGenrateKey),
	})
	privateKeyPem.Close()

	publicKey := keyGenrateKey.PublicKey
	rsaKey, _ := rsa.EncryptOAEP(sha256.New(), rand.Reader, &publicKey, []byte("super secret bytes"), nil)
	rsaKeyBase64 := base64.StdEncoding.EncodeToString(rsaKey)
	return string(rsaKeyBase64), nil

}
func GenerateToken(username string, password string) (token string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaims{
		UserName: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	fmt.Print("jwt claims---", claims)
	JWTToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println("token jwt --", JWTToken)
	tokenString, err := JWTToken.SignedString(jwtKey)
	if err != nil {
		fmt.Println("error while signing string--", err.Error())
		return "", err
	}
	fmt.Println("jwt token string--", tokenString)
	return tokenString, nil

}

func ValidateToken(jwttoken string) (bool, error) {
	fmt.Println("validate token called")
	token, err := jwt.ParseWithClaims(jwttoken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		fmt.Println("error block called")
		return false, errors.New("invalid token")
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return false, errors.New("invalid token")

	}
	fmt.Println("claims returned---", claims.UserName, "--", claims.Password)
	if claims.ExpiresAt < time.Now().Local().Unix() {
		fmt.Println("token expired check login---")
		return false, errors.New("token expired")
	}
	fmt.Println("token is valid")
	return true, nil
}
