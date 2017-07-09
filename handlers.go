package main

import (
	"fmt"
	"net/http"
)

type Handlers struct {
	dbConnection *DBConnection
	uController  *UsersController
}

func CreateApiHandlers() *Handlers {
	API := &Handlers{
		dbConnection: OpenConnectionSession(),
		uController:  &UsersController{},
	}
	API.uController.dbConnection = API.dbConnection

	return API
}

func (c *Handlers) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}
