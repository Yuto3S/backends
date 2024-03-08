package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func handleSubmitDetails(writer http.ResponseWriter, req *http.Request, tmpl *template.Template) {
	if req.Method != http.MethodPost {
		tmpl.Execute(writer, nil)
		return
	}

	details := ContactDetails{
		Email:   req.FormValue("email"),
		Subject: req.FormValue("subject"),
		Message: req.FormValue("message"),
	}

	fmt.Println(details)

	tmpl.Execute(writer, struct{ Success bool }{true})
}
