package main

import (
	"fmt"
	"log"
	"net/http"
)


var db *Database
var rem *Database

func main() {
	db = New("nero")
	db.Init()

	rem = New("remaining")
	rem.Init()

	router := setupRouter(AllRoutes())

	go ResetRemainingCRON()

	fmt.Println("Running on port 3000")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}
}
