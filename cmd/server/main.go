package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handleFunc).Methods("GET")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		panic(err)
	}
}
func handleFunc(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("rabotaet")
}
