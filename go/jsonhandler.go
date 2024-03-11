package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func decodeUserFromRequest(writer http.ResponseWriter, req *http.Request) {
	var user User
	json.NewDecoder(req.Body).Decode(&user)
	fmt.Fprintf(writer, "%s %s is %d years old!", user.Firstname, user.Lastname, user.Age)
}

func encodeUser(writer http.ResponseWriter, req *http.Request) {
	peter := User{
		Firstname: "Peter",
		Lastname:  "Pan",
		Age:       15,
	}

	json.NewEncoder(writer).Encode(peter)
}
