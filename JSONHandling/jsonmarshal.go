package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	BookName string `json:"bookname"`
	Author   Author `json:"author"`
}
type Author struct {
	Name   string `json:"name"`
	Age    int    `json:"number"`
	Gender string `json":"gender"`
}

func main() {
	author := Author{Name: "Himanshu Singh", Age: 22, Gender: "male"}
	book := Book{BookName: "abcd", Author: author}

	byteArray, err := json.MarshalIndent(book, "", "   ")
	if err != nil {
		fmt.Println("Error occured")
	}
	fmt.Println(string(byteArray))

}
