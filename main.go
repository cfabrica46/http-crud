package main

import (
	"log"
	"net/http"
)

type user struct {
	ID                 int
	Username, Password string
}

type p int

const (
	find   p = 1
	delete p = 2
)

var users = []user{
	{
		ID:       1,
		Username: "cfabrica46",
		Password: "01234",
	},
	{
		ID:       2,
		Username: "arthuronavah",
		Password: "456456",
	},
}

func main() {

	http.HandleFunc("/", index)

	executeAllHandleFuncs()

	http.HandleFunc("/users/all", findUsers)

	http.HandleFunc("/user/create", createUser)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}

}
