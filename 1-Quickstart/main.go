package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Todo struct {
	Id      int
	Message string
}

func main() {
	fmt.Println("Hello World!!!")

	data := map[string][]Todo{
		"Todos": {
			Todo{Id: 1, Message: "Buy Milk"},
		},
	}

	todosHandler := func(w http.ResponseWriter, r *http.Request) {
		tpl := template.Must(template.ParseFiles("index.html"))
		tpl.Execute(w, data)
	}

	addTodoHandler := func(w http.ResponseWriter, r *http.Request) {
		msg := r.PostFormValue("message")
		tpl := template.Must(template.ParseFiles("index.html"))
		todo := Todo{Id: len(data["Todos"]) + 1, Message: msg}
		data["Todos"] = append(data["Todos"], todo)

		tpl.ExecuteTemplate(w, "todo-list-element", todo)
	}

	http.HandleFunc("/", todosHandler)
	http.HandleFunc("/add-todo", addTodoHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
