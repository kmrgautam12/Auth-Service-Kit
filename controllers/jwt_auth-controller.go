package controllers

import (
	jwt_auth "Book-Rental-Service/JWT-Auth"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Token(c *gin.Context) {
	userName := c.Query("username")
	password := c.Query("password")
	fmt.Println("username--", userName, "password--", password)
	token, err := jwt_auth.GenerateToken(userName, password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "unable to generate token"})
	}
	c.JSON(http.StatusOK, gin.H{"token": token})

}
