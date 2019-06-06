package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/frankienicoletti/fullstack-exercise/src/server/adapters"
	"github.com/frankienicoletti/fullstack-exercise/src/server/routers"
)

func main() {
	// Connection.
	mySQL := adapters.NewMySQLStore()
	peopleAPI := routers.NewPeopleAPI(mySQL)

	fmt.Println("connected")

	// Serve
	log.Fatal(http.ListenAndServe(":4000", peopleAPI.Router))
}
