package users

import (
	"errors"
	"fmt"
)

type UserInfo struct {
	UserName     string `json:"user_name"`
	Password     string `json:"password,omitempty"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phone_number"`
	Country      string `json:"country"`
	State        string `json:"state"`
	ZipCode      string `json:"zip_code"`
	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2"`
}

func ValidateUser(reqBody map[string]interface{}) (bool, error) {
	fmt.Println("Inside validateuser  ", reqBody)
	userName := reqBody["UserName"].(string)
	exist, err := isUserNameExistAlready(userName)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	if exist {
		fmt.Println("user already exist--")
		return false, nil
	}
	if len(userName) > 255 || len(userName) == 0 {
		return false, errors.New("username is not valid")
	}

	return true, nil

}
