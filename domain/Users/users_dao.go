package users

import (
	userdb_test "Book-Rental-Service/databases/mysql/user_db"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

const (
	insertUser       = "Insert into Users (UserName,FirstName,LastName,Email,PhoneNumber,ZipCode,State,Country,AddressLine1,AddressLine2) values (?,?,?,?,?,?,?,?,?,?) "
	updateUser       = "update Users Set FirstName=?,LastName=?, Email =?,  Zipcode = ?, PhoneNumber=? where UserName=?"
	isUserNameUnique = "select * from Users where  UserName = ? "
	isUserPresent    = "select * from Users where UserName= ?"
	deleteuser       = "delete from Users where  UserName=? "
	deleteCredential = "delete from UserCredential where UserName=?"
	insertCredential = "Insert into UserCredential (UserName,Password) values (?,?) "
)

func StoreUserCredentials(username string, password string) error {
	fmt.Println("username is ", username, " password is ", password)
	stmt, err := userdb_test.DBClient.Prepare(insertCredential)
	if err != nil {
		return err
	}
	stmtexec, err := stmt.Exec(username, password)
	if err != nil {
		return err
	}
	fmt.Println("statemt after executing query ", stmtexec)
	return nil
}

func GetUser(username string) ([]UserInfo, bool) {
	fmt.Println("username is ", username)
	rows, err := userdb_test.DBClient.Query(isUserPresent, username)
	fmt.Println("sql rows returned ---", rows)

	if err != nil {
		fmt.Println("unable to query ", err.Error())
	}
	defer rows.Close()
	users := make([]UserInfo, 0)
	var user UserInfo

	fmt.Println("total colums", rows.Scan(user.FirstName, user.LastName, user.Email,
		user.ZipCode, user.PhoneNumber, user.UserName, user.Country, user.AddressLine1, user.AddressLine2))

	for rows.Next() {
		err := rows.Scan(&user.FirstName, &user.LastName, &user.Email,
			&user.ZipCode, &user.PhoneNumber, &user.UserName, &user.Country, &user.AddressLine1, &user.AddressLine2, &user.State)

		log.Println(user)
		switch err {
		case sql.ErrNoRows:
			return users, false

		}
		users = append(users, user)
	}
	return users, true

}

func CreateUser(reqBody map[string]interface{}) error {

	if err := userdb_test.DBClient.Ping(); err != nil {
		fmt.Println("db not configured correctly--")

		panic(err)
	}

	queryStatement, err := userdb_test.DBClient.Prepare(insertUser)
	if err != nil {
		fmt.Println("error block for save user is executed")
		return errors.New("Error in preparing the query statement")

	}
	fmt.Println("Quer statement ----", queryStatement)
	userName := reqBody["UserName"].(string)
	firstName := reqBody["FirstName"].(string)
	lastName := reqBody["LastName"].(string)
	email := reqBody["Email"].(string)
	phoneNumber := reqBody["PhoneNumber"].(string)
	state := reqBody["State"].(string)
	country := reqBody["Country"].(string)
	addressline1 := reqBody["AddressLine1"].(string)
	addressLine2 := reqBody["AddressLine2"].(string)
	zipCode := reqBody["ZipCode"].(string)

	execQueryStatement, err := queryStatement.Exec(userName, firstName, lastName,
		email, phoneNumber, zipCode, state, country, addressline1, addressLine2)

	if err != nil {
		fmt.Println("error in executing sql query---")
		return errors.New("Error in executing the sql query :---")
	}
	fmt.Println("sql result after executing query--", execQueryStatement)

	return nil

}

func DeleteUser(user_name string) error {
	queryStmt, err := userdb_test.DBClient.Prepare(deleteuser)
	if err != nil {
		fmt.Println("error in preparing statement")
		return err

	}
	fmt.Println("inside delete user")
	sqlResult, err := queryStmt.Exec(user_name)
	if err != nil {
		fmt.Println("sql result err--", err.Error())
		return err
	}
	rowsAffected, err := sqlResult.RowsAffected()
	if err != nil {
		fmt.Println("error getting row affected--", err.Error())
	}
	fmt.Println("rows affected--", rowsAffected)

	// removing credential of user
	_, err = userdb_test.DBClient.Exec(deleteCredential, user_name)
	if err != nil {
		return errors.New("unable to delete user")
	}
	fmt.Println("result from sql after delete query execution--", sqlResult)
	return nil

}

func isUserNameExistAlready(username string) (bool, error) {
	if err := userdb_test.DBClient.Ping(); err != nil {
		panic(err)
	}
	queryExec := userdb_test.DBClient.QueryRow(isUserNameUnique, username)
	err := queryExec.Scan(&username)
	switch err {
	case sql.ErrNoRows:
		return false, nil
	default:
		return true, nil

	}
}
