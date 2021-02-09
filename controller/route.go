package controller

import (
	"net/http"
)

func Register() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", status())
	mux.HandleFunc("/create", create())
	mux.HandleFunc("/readall", create())
	mux.HandleFunc("/name", SearchbyID())
	mux.HandleFunc("/delete", DeletebyID())
	return mux
}
