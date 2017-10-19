package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"fmt"
)

var db *sql.DB

var (
	SQL_USERNAME = os.Getenv("HBB_MYSQL_USER")
	SQL_PASSWORD = os.Getenv("HBB_MYSQL_PASS")
	SQL_HOST= os.Getenv("HBB_DBHOST")+":3306"
	DB_NAME = os.Getenv("HBB_DBNAME")
)

func connectToMysql(){
	fmt.Println("Connecting to MySQL Driver", SQL_USERNAME, SQL_PASSWORD, DB_NAME)
	db1, err := sql.Open("mysql", SQL_USERNAME + ":" + SQL_PASSWORD + "@tcp(" + SQL_HOST + ")/" + DB_NAME)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	db = db1

	_,err = db.Exec("USE " + DB_NAME)
	if err != nil {
		fmt.Errorf(err.Error())
	}
}
