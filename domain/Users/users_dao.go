package users

import (
	userdb_test "Book-Rental-Service/databases/mysql/user_db"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

const (
	insertUser       = "Insert into Users (UserName,Credential,ID) values (?,?,?) "
	updateUser       = "update Users Set FirstName=?,LastName=?, Email =?,  Zipcode = ?, PhoneNumber=? where UserName=?"
	isUserNameUnique = "select Email from Users where  UserName = ? "
	isUserPresent    = "select UserName,credential from Users where UserName= ?"
	deleteuser       = "delete from Users where  UserName=?"
)

func GetUser(user UserInfo, id string) ([]UserInfo, bool) {
	fmt.Println("is user id pointer nill--", &user.Id == nil)
	rows, err := userdb_test.DBClient.Query(isUserPresent, id)
	fmt.Println("sql rows returned ---", rows)

	if err != nil {
		fmt.Println("unable to query ", err.Error())
	}
	defer rows.Close()
	users := make([]UserInfo, 0)

	for rows.Next() {
		err := rows.Scan(&user.UserName, &user.Credential)
		log.Println(user)
		switch err {
		case sql.ErrNoRows:
			return users, false

		}
		users = append(users, user)
	}
	return users, true

}

func CreateUser(user *UserInfo) error {

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

	execQueryStatement, err := queryStatement.Exec(&user.UserName, &user.Credential, &user.Id)
	if err != nil {
		fmt.Println("error in executing sql query---")
		return errors.New("Error in executing the sql query :---")
	}
	fmt.Println("sql result after executing query--", execQueryStatement)

	return nil

}
func UpdateUser(user *UserInfo) error {

	if err := userdb_test.DBClient.Ping(); err != nil {
		fmt.Println("db not configured correctly--")

		panic(err)
	}

	queryStatement, err := userdb_test.DBClient.Prepare(updateUser)
	if err != nil {
		fmt.Println("error block for save user is executed")
		return errors.New("Error in preparing the query statement")

	}
	fmt.Println("Quer statement ----", queryStatement)

	execQueryStatement, err := queryStatement.Exec(&user.FirstName, &user.LastName, &user.Email, &user.ZipCode, &user.PhoneNumber, &user.UserName)
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
		return err

	}
	defer queryStmt.Close()
	sqlResult, err := queryStmt.Exec(user_name)
	if err != nil {
		fmt.Println("sql result err--", err.Error())
		return err
	}
	rowsAffected, err := sqlResult.RowsAffected()
	fmt.Println("rows affected--", rowsAffected)
	if rowsAffected == 0 {
		return errors.New("unable to delete user")
	}
	fmt.Println("result from sql after delete query execution--", sqlResult)
	return nil

}

func isUserNameExistAlready(user *UserInfo) (bool, error) {
	if err := userdb_test.DBClient.Ping(); err != nil {
		panic(err)
	}
	username := user.UserName
	queryExec := userdb_test.DBClient.QueryRow(isUserNameUnique, username)
	err := queryExec.Scan(&username)
	switch err {

	case sql.ErrNoRows:
		return true, nil
	default:
		return false, nil

	}
}
