package services

import (
	users "Book-Rental-Service/domain/Users"
	"errors"
)

func CreateUserService(reqBody map[string]interface{}) error {

	valid, err := users.ValidateUser(reqBody)
	if err != nil || valid == false {
		return errors.New("invalid user")
	}
	err = users.CreateUser(reqBody)
	if err != nil {
		return err
	}
	return nil
}
