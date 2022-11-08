package userdb_test

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DBClient *sql.DB
var err error

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",

		"root",
		"Rahul@#$987",
		"127.0.0.1:3306",
		"UserDB")
	fmt.Println("datasource name :-", dataSourceName)

	DBClient, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	fmt.Println("database configured correctly---")

}
