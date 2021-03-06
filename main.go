package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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

	executeAllHandleFuncFindUser()

	http.HandleFunc("/users/all", findUsers)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}

}

func index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "BIENVENIDOS :D")

}

func executeAllHandleFuncFindUser() {

	for i := range users {

		id := strconv.Itoa(users[i].ID)

		route := fmt.Sprintf("/users/%s", id)

		http.HandleFunc(route, findUser)
	}

}

func findUsers(w http.ResponseWriter, r *http.Request) {

	err := json.NewEncoder(w).Encode(users)

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

}

func findUser(w http.ResponseWriter, r *http.Request) {

	var u user

	idString := getIDFromURL(w, r)

	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for i := range users {

		if users[i].ID == id {

			u = users[i]
			break

		}

	}

	err = json.NewEncoder(w).Encode(u)

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

}

func getIDFromURL(w http.ResponseWriter, r *http.Request) (id string) {

	url := []byte(r.URL.Path)
	id = fmt.Sprintf("%s", url[7:])

	return

}
