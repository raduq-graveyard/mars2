package handler

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type planet struct {
	MaxX int
	MaxY int
}

type robot struct {
	AtX int
	AtY int
	Ori string
}

// NavHandler and stuff
func NavHandler(w http.ResponseWriter, r *http.Request) {
	commands := commands(r)
	planet := planet{5, 5}
	robot := robot{0, 0, "N"}
	pos, err := Walker(planet, robot, commands)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pos)
}

func commands(r *http.Request) []string {
	command := mux.Vars(r)["command"]
	return strings.Split(command, "")
}
