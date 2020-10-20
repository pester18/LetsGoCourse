package api

import (
	"net/http"

	"encoding/json"
	"fmt"

	"basic_crud/db"

	"github.com/gorilla/mux"
)

func handleError(err error, message string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(fmt.Sprintf(message, err)))
}

func GetAllItems(w http.ResponseWriter, req *http.Request) {
	rs, err := db.GetAllItems()
	if err != nil {
		handleError(err, "Failed to load database items: %v", w)
		return
	}

	bs, err := json.Marshal(rs)
	if err != nil {
		handleError(err, "Failed to load marshal data: %v", w)
		return
	}

	w.Write(bs)
}

func GetItem(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	rs, err := db.GetItem(id)
	if err != nil {
		handleError(err, "Failed to read database: %v", w)
		return
	}

	bs, err := json.Marshal(rs)
	if err != nil {
		handleError(err, "Failed to marshal data: %v", w)
		return
	}

	w.Write(bs)
}

func PostItem(w http.ResponseWriter, req *http.Request) {
	var item db.Item

	err := json.NewDecoder(req.Body).Decode(&item)
	if err != nil {
		handleError(err, "Failed to parse input data: %v", w)
		return
	}

	if err = db.SaveItem(item); err != nil {
		handleError(err, "Failed to save data: %v", w)
		return
	}

	w.Write([]byte("OK"))
}

func DeleteItem(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	if err := db.RemoveItem(id); err != nil {
		handleError(err, "Failed to remove item: %v", w)
		return
	}

	w.Write([]byte("OK"))
}
