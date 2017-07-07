package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/raduq/mars2/handler"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/rest/mars/{command}", handler.NavHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
