package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":8080", r)
}
