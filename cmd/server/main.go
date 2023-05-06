package main

import (
	"encoding/json"
	"github.com/Onelvay/halyklife-test-task/pkg/domain"
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
	id := r.Header.Get("X-Request-Id")
	resp := domain.Response{id, http.StatusOK}
	jb, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jb)
	w.WriteHeader(http.StatusOK)

}
