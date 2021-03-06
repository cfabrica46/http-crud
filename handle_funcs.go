package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "BIENVENIDOS :D")

}

func findUser(w http.ResponseWriter, r *http.Request) {

	u := getUser(w, r, "find")

	err := json.NewEncoder(w).Encode(u)

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

}

func findUsers(w http.ResponseWriter, r *http.Request) {

	err := json.NewEncoder(w).Encode(users)

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

}

func createUser(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		t, err := template.ParseFiles("create_user.html")

		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		err = t.Execute(w, nil)

		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

	case "POST":

		id64 := time.Now().UnixNano()

		idString := strconv.FormatInt(id64, 10)

		id, err := strconv.Atoi(idString)

		err = r.ParseForm()

		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		username := r.Form.Get("username")
		password := r.Form.Get("password")

		if username == "" || password == "" {
			http.Error(w, http.StatusText(400), 400)
		}

		for i := range users {

			if users[i].ID != id {
				break
			}
			id++
		}

		users = append(users, user{ID: id, Username: username, Password: password})

		fmt.Fprintf(w, "SE AGREGO UN NUEVO USUARIO:\n")

		for i := range users {

			if id == users[i].ID {

				json.NewEncoder(w).Encode(users[i])
			}
		}

		go addHandleFuncs(id)

	}

}

func deleteUser(w http.ResponseWriter, r *http.Request) {

	u := getUser(w, r, "delete")

	for i := range users {

		if u.ID == users[i].ID {

			users = append(users[:i], users[i+1:]...)
			break
		}

	}
	fmt.Fprint(w, "SE ELIMINO EL USUARIO")

}

func updateUser(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		t, err := template.ParseFiles("update_user.html")

		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		err = t.Execute(w, nil)

		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

	case "POST":

		var u user

		url := []byte(r.URL.Path)
		idString := fmt.Sprintf("%s", url[13:])

		id, err := strconv.Atoi(idString)

		err = r.ParseForm()

		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		username := r.Form.Get("username")
		password := r.Form.Get("password")

		if username == "" || password == "" {
			http.Error(w, http.StatusText(400), 400)
		}

		for i := range users {

			if users[i].ID == id {

				users[i].Username = username
				users[i].Password = password
				u = users[i]
				break
			}

		}

		fmt.Fprintf(w, "SE ACTUALIZO TUS DATOS:\n")

		json.NewEncoder(w).Encode(u)

	}

}
