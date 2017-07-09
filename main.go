package main

import (
	"log"
	"net/http"
)

func main() {

	socialApi := CreateApiHandlers()
	router := CreateNewRouter(socialApi)

	log.Fatal(http.ListenAndServe(":8080", router))

}
