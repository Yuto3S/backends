package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func getData() TodoPageData {
	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
			{Title: "Do the laundry", Done: false},
		},
	}
	return data
}

func renderTemplate(writer http.ResponseWriter, tmpl *template.Template) {
	data := getData()
	fmt.Println(data)
	fmt.Println(tmpl)
	tmpl.Execute(writer, data)
}
