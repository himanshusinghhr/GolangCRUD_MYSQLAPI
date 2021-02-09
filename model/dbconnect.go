package model

import (
	"database/sql"
	"fmt"

	//blank import
	_ "github.com/go-sql-driver/mysql"
)

var con *sql.DB

//DBinstance is
func DBinstance() *sql.DB {

	db, err := sql.Open("mysql", "root:1234@/himanshu")

	if err != nil {
		panic(err)
	}
	fmt.Println("DB Connected")
	con = db
	return db
}
