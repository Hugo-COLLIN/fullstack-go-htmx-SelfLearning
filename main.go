package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		data := struct {
			Message string
		}{
			Message: "Hello World",
		}

		tmpl.Execute(w, data)
	})

	fmt.Println("Started http server on port :8080")
	http.ListenAndServe(":8080", nil)
}
