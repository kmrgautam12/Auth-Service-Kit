package controllers

import (
	users "Book-Rental-Service/domain/Users"
	"Book-Rental-Service/services"
	"Book-Rental-Service/utils"
	error "Book-Rental-Service/utils"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"

	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
	reqBodyByte, err := utils.GetRequestBodyMap(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
	}
	user := users.UserInfo{}
	reqBodyMap := make(map[string]interface{}, 0)
	json.Unmarshal(reqBodyByte, &reqBodyMap)
	fmt.Println("user request body--", reqBodyMap)
	id := uuid.New().String()
	userName := reqBodyMap["UserName"].(string)
	password := reqBodyMap["Password"].(string)
	passByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
	}
	hashPassword := string(passByte)
	user.UserName = userName
	user.Credential = hashPassword
	user.Id = id
	fmt.Println("Hash to store:", hashPassword)
	user, err = services.CreateUser(&user)
	if err != nil {
		c.JSON(404, gin.H{"message": err.Error()})
		return

	}
	c.JSON(201, gin.H{"message": "user created"})

}

func UpdateUser(c *gin.Context) {

	var user *users.UserInfo

	requestBodyByte, err := utils.GetRequestBodyMap(c)
	if err != nil {
		restError := error.RestError{
			Message: "Error in binding json data",
			Status:  "bad_request",
		}
		c.JSON(http.StatusBadRequest, restError)
		return
	}
	fmt.Println("request body string --", string(requestBodyByte))

	json.Unmarshal(requestBodyByte, &user)
	fmt.Println("req body after marshelling--", user)
	err = services.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return

	}
	c.JSON(http.StatusOK, gin.H{"message": "user updated"})

}
func GetUser(c *gin.Context) {
	var user users.UserInfo
	userId := c.Query("id")
	fmt.Println("user id from params--", userId)
	users, exist := users.GetUser(user, userId)
	if !exist || len(users) == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"message": "user doesn't exist"})
		return
	}
	fmt.Println("is user present flag--", exist)
	c.JSON(http.StatusOK, users)

}

func DeleteUser(c *gin.Context) {
	userName := c.Query("id")
	fmt.Println("user id from params---", userName)
	err := users.DeleteUser(userName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "unable to delete user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})

}
func Login(c *gin.Context) {
	reqBodyByte, err := utils.GetRequestBodyMap(c)
	var user users.UserInfo
	usersArray := make([]users.UserInfo, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
	}
	var reqBody map[string]interface{}
	err = json.Unmarshal(reqBodyByte, &reqBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})

	}
	userName := reqBody["UserName"].(string)
	fmt.Println("username--", userName)
	password := reqBody["Password"].(string)

	usersArray, flag := users.GetUser(user, userName)
	if flag == false {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user is not present"})

	}
	fmt.Println(usersArray)
	userDbPassword := usersArray[0].Credential
	fmt.Println("userdb password---", userDbPassword)
	err = bcrypt.CompareHashAndPassword([]byte(userDbPassword), []byte(password))
	if err != nil {
		c.JSON(http.StatusAccepted, gin.H{"message": "user authenticated"})
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "forbidden"})

}
