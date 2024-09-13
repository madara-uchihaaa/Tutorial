package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	greeting()
	r := mux.NewRouter()
	r.HandleFunc("/", server).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func greeting() string {
	return "Hello, World!"
}

func server(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
		<h1>` + greeting() + `</h1>
	`))
}
