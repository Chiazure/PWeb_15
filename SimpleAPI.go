package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type student struct {
	ID   string
	NIM  string
	NAME string
}

var data = []student{
	student{"1", "5042138", "Nayu"},
	student{"2", "5042167", "Izmi"},
	student{"3", "5142153", "Vieri"},
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var requestData map[string]interface{}

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&requestData); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id := requestData["id"].(string)

		var result []byte
		var err error
		for _, each := range data {
			if each.ID == id {
				result, err = json.Marshal(each)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.Write(result)
				return
			}
		}
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/mahasiswas", users)
	http.HandleFunc("/mahasiswa", user)
	fmt.Println("starting web server at http://localhost:6060/")
	http.ListenAndServe(":6060", nil)
}
