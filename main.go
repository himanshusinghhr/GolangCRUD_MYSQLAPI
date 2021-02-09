package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/todoapi/controller"
	"github.com/todoapi/model"
)

func handleRequest() {
	mux := controller.Register()
	db := model.DBinstance()
	defer db.Close()
	fmt.Println("Serving")
	log.Fatal(http.ListenAndServe("localhost:3000", mux))

}

// func handleCalculator(){
// 	mux:=controller
// }

func main() {
	handleRequest()

}
