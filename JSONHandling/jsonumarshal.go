package main

import (
	"encoding/json"
	"fmt"
)

type Author struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json":"gender"`
}

func main() {

	// var jsonString = `{"name":"Himansuh","age":22,"gender":"male"}`
	// var author Author
	// err := json.Unmarshal([]byte(jsonString), &author)
	// if err != nil {
	// 	fmt.Println("Error Occured")
	// }
	// fmt.Println(author)

	//If we do not know the struct before hand we can do

	var jsonString = `{"name":"Himansuh","age":22,"gender":"male"}`
	var author1 map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &author1)
	if err != nil {
		fmt.Println("Error Occured")
	}
	fmt.Println(author1)
	for i, k := range author1 {
		fmt.Println(i)
		fmt.Println(k)
	}

}
