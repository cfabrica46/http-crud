package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func executeAllHandleFuncs() {

	for i := range users {

		id := strconv.Itoa(users[i].ID)

		routeFindUser := fmt.Sprintf("/users/%s", id)

		http.HandleFunc(routeFindUser, findUser)

		routeDeleteUser := fmt.Sprintf("/user/delete/%s", id)

		http.HandleFunc(routeDeleteUser, deleteUser)

		routeUpdateUser := fmt.Sprintf("/user/update/%s", id)

		http.HandleFunc(routeUpdateUser, updateUser)

	}

}

func getIDFromURLFindUser(w http.ResponseWriter, r *http.Request) (id string) {

	url := []byte(r.URL.Path)
	id = fmt.Sprintf("%s", url[7:])

	return

}

func getIDFromURLDeleteUser(w http.ResponseWriter, r *http.Request) (id string) {

	url := []byte(r.URL.Path)
	id = fmt.Sprintf("%s", url[13:])

	return

}

func getIDFromURLUpdateUser(w http.ResponseWriter, r *http.Request) (id string) {

	url := []byte(r.URL.Path)
	id = fmt.Sprintf("%s", url[13:])

	return

}

func addHandleFuncs(id int) {

	idString := strconv.Itoa(id)

	routeFindUser := fmt.Sprintf("/users/%s", idString)

	http.HandleFunc(routeFindUser, findUser)

	routeDeleteUser := fmt.Sprintf("/user/delete/%s", idString)

	http.HandleFunc(routeDeleteUser, deleteUser)

	routeUpdateUser := fmt.Sprintf("/user/update/%s", idString)

	http.HandleFunc(routeUpdateUser, updateUser)

}

func getUser(w http.ResponseWriter, r *http.Request, handle p) (u user) {

	var check bool
	switch handle {
	case find:
		idString := getIDFromURLFindUser(w, r)

		id, err := strconv.Atoi(idString)

		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		for i := range users {

			if users[i].ID == id {

				u = users[i]
				check = true
				break

			}

		}

		if check == false {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		return

	case delete:

		idString := getIDFromURLDeleteUser(w, r)

		id, err := strconv.Atoi(idString)

		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		for i := range users {

			if users[i].ID == id {

				u = users[i]
				check = true
				break

			}

		}

		if check == false {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		return
	}
	return
}
