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

	db, err := sql.Open("mysql", "root:<root_password>@/<database_name>")

	if err != nil {
		panic(err)
	}
	fmt.Println("DB Connected")
	con = db
	return db
}
