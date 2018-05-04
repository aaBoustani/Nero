package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := setupRouter(AllRoutes())

	fmt.Println("Running on port 3000")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}
}