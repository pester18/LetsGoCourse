package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"../person"
)

func main() {
	personCache := person.NewPersonCache()
	handler := createMainHandler(personCache)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8090", nil)
}

func createMainHandler(personCache *person.PersonCache) func(http.ResponseWriter, *http.Request) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			ids, ok := r.URL.Query()["id"]

			if !ok || len(ids[0]) < 1 {
				http.Error(w, "Param 'id' is missing", http.StatusBadRequest)
				return
			}

			stringId := ids[0]
			personId, err := strconv.Atoi(stringId)

			if err != nil {
				http.Error(w, "Invalid 'id'", http.StatusBadRequest)
				return
			}

			p := personCache.GetPerson(personId)

			js, err := json.Marshal(p)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
		case "POST":
			var p person.Person

			err := json.NewDecoder(r.Body).Decode(&p)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			personId := personCache.AddPerson(p)

			resp := struct{ Id int }{personId}

			js, err := json.Marshal(resp)
			if err != nil {
				fmt.Println(err.Error())
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
		default:
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		}
	}
	return handler
}
