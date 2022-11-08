package users

import (
	"fmt"
)

type UserInfo struct {
	UserName    string
	Credential  string
	Id          string
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	ZipCode     string
}

func ValidateUser(user *UserInfo) (bool, error) {
	fmt.Println("Validate User Pointer--", user)
	exist, err := isUserNameExistAlready(user)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	if !exist {
		fmt.Println("user already exist--")
		return false, nil
	}
	fmt.Println("Email does not exit in db")
	err = CreateUser(user)
	if err == nil {
		fmt.Println("user has been saved in database")
		return true, nil

	}
	return false, nil
}
