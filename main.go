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
	http.HandleFunc("/", todosHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
