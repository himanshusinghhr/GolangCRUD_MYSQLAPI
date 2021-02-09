package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/todoapi/model"
	"github.com/todoapi/view"
)

func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			data := view.PostRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			fmt.Println(data)
			json.NewEncoder(w).Encode(data)

			if err := model.CreateTodoo(data.Name, data.Todo); err != nil {
				w.Write([]byte("Some error occured"))
				fmt.Println("Error")
				return
			}
			w.WriteHeader(http.StatusCreated) //Printing 201
		} else if r.Method == http.MethodGet {

			data, err := model.ReadAll()
			if err != nil {
				w.Write([]byte("Error Occured"))
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)

		}
	}

}
func SearchbyID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			name := r.URL.Query().Get("name")
			data, err := model.ReadByName(name)
			if err != nil {
				w.Write([]byte("Error Occured"))
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		}
	}
}
func DeletebyID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			name := r.URL.Query().Get("name")
			err := model.Delete(name)
			if err != nil {
				w.Write([]byte("Error Occured"))
			}
			data, err := model.ReadAll()
			if err != nil {
				w.Write([]byte("Error Occured"))
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)

			// w.WriteHeader(http.StatusOK)
			// json.NewEncoder(w).Encode(data)
		}
	}
}
