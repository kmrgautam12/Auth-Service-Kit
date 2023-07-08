package controllers

import (
	users "Book-Rental-Service/domain/Users"
	"Book-Rental-Service/services"
	"Book-Rental-Service/utils"
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {

	reqBodyByte, err := utils.GetRequestBodyMap(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
	}
	reqBodyMap := make(map[string]interface{}, 0)
	json.Unmarshal(reqBodyByte, &reqBodyMap)
	fmt.Println("user request body--", reqBodyMap["UserName"])
	fmt.Println("user request body--", reqBodyMap["PassWord"])

	userName := reqBodyMap["UserName"].(string)
	password := reqBodyMap["PassWord"].(string)
	passByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
	}

	hashPassword := string(passByte)

	fmt.Println("Hash password to store:", hashPassword)
	err = users.StoreUserCredentials(userName, hashPassword)
	if err != nil {
		c.JSON(500, gin.H{"message": "internal server error"})
	}
	err = services.CreateUserService(reqBodyMap)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return

	}
	c.JSON(201, gin.H{"message": "user created"})

}
func GetUser(c *gin.Context) {

	userName := c.Query("username")
	fmt.Println("user id from params--", userName)
	users, exist := users.GetUser(userName)
	if !exist || len(users) == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"message": "user doesn't exist"})
		return
	}
	fmt.Println("is user present flag--", exist)
	c.JSON(http.StatusOK, gin.H{"data": users})

}

func DeleteUser(c *gin.Context) {
	userName := c.Query("username")
	fmt.Println("user id from params---", userName)
	err := users.DeleteUser(userName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "unable to delete user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})

}
