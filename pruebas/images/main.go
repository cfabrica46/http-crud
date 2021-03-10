package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", index)

	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("index.html")

	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(w, nil)

	if err != nil {
		log.Fatal(err)
	}

}
