package main

import (
	"net/http"

	"basic_crud/api"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/items", api.GetAllItems).Methods("GET")
	r.HandleFunc("/api/items/{id}", api.GetItem).Methods("GET")
	r.HandleFunc("/api/items", api.PostItem).Methods("POST")
	r.HandleFunc("/api/items/{id}", api.DeleteItem).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}
