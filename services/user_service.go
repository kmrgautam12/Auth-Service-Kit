package services

import (
	users "Book-Rental-Service/domain/Users"
	"errors"
	"fmt"
)

func CreateUser(user *users.UserInfo) (users.UserInfo, error) {

	valid, err := users.ValidateUser(user)
	if err != nil || valid == false {
		return users.UserInfo{}, errors.New("invalid user")
	}
	fmt.Println("Is user valid ", valid)
	if !valid {
		return users.UserInfo{}, errors.New("invalid user")
	}

	return *user, nil
}
func UpdateUser(user *users.UserInfo) error {

	err := users.UpdateUser(user)
	if err != nil {
		return errors.New("unable to update user")
	}

	return nil
}
