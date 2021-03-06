package main

import (
	"log"
	"net/http"
)

type user struct {
	ID                 int
	Username, Password string
}

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

	fileServer := http.FileServer(http.Dir("./images"))

	http.Handle("/images/", http.StripPrefix("/images", fileServer))

	http.HandleFunc("/users/one", findUser)

	http.HandleFunc("/users/all", findUsers)

	http.HandleFunc("/user/create", createUser)

	http.HandleFunc("/user/delete", deleteUser)

	http.HandleFunc("/user/update", updateUser)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}

}
