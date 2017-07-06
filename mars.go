package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/rest/mars/{command}", NavHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}

type planet struct {
	MaxX int
	MaxY int
}

type robot struct {
	AtX int
	AtY int
	Ori string
}
