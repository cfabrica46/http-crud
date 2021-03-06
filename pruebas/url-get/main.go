package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", home)

	http.ListenAndServe(":8080", nil)

}

func home(w http.ResponseWriter, r *http.Request) {

	fmt.Print(r.URL.Path)

	for i, v := range r.URL.Query() {

		fmt.Fprintf(w, "%s-%-s\n", i, v)

		fmt.Fprintf(w, "%T-%-T\n", i, v)

	}

}
