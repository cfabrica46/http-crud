package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		t, err := template.ParseFiles("./templates/index.html")

		if err != nil {
			log.Fatal(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), 500)
			return
		}

		err = t.Execute(w, nil)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), 500)
			return
		}
	}
}

func findUser(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		u, err := getUser(r.URL.Query())

		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), 400)
			return
		}

		err = json.NewEncoder(w).Encode(u)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), 500)
			return
		}
	}
}

func findUsers(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		err := json.NewEncoder(w).Encode(users)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), 500)
			return
		}
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		t, err := template.ParseFiles("./templates/create_user.html")

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), 500)
			return
		}

		err = t.Execute(w, nil)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), 500)
			return
		}

	case "POST":

		id64 := time.Now().UnixNano()

		idString := strconv.FormatInt(id64, 10)

		id, err := strconv.Atoi(idString)

		err = r.ParseForm()

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), 500)
			return
		}

		username := r.Form.Get("username")
		password := r.Form.Get("password")

		if username == "" || password == "" {
			http.Error(w, http.StatusText(http.StatusNotAcceptable), 400)
			return
		}

		for i := range users {

			if users[i].ID != id {
				break
			}
			id++
		}

		users = append(users, user{ID: id, Username: username, Password: password})

		for i := range users {

			if id == users[i].ID {

				json.NewEncoder(w).Encode(users[i])
			}
		}

	}

}

func deleteUser(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		u, err := getUser(r.URL.Query())

		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), 400)
			return
		}

		if len(users) < 1 {

			users = []user{}

		} else {

			for i := range users {
				if users[i] == u {
					users = append(users[:i], users[i+1:]...)
					break
				}
			}
		}

		err = json.NewEncoder(w).Encode(u)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), 500)
			return
		}
	}

}

func updateUser(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":

		u, err := getUser(r.URL.Query())

		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), 400)
			return
		}

		t, err := template.ParseFiles("./templates/update_user.html")

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), 500)
			return
		}

		err = t.Execute(w, u.ID)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), 500)
			return
		}
	case "POST":

		var u user

		err := r.ParseForm()

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), 500)
			return
		}

		id := r.Form.Get("id")
		username := r.Form.Get("username")
		password := r.Form.Get("password")

		if username == "" || password == "" {
			http.Error(w, http.StatusText(http.StatusNotAcceptable), 400)
			return
		}

		iD, err := strconv.Atoi(id)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), 500)
			return
		}

		u.ID = iD
		u.Username = username
		u.Password = password

		for i := range users {

			if users[i].ID == u.ID {
				users[i] = u
			}

		}

		err = json.NewEncoder(w).Encode(u)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), 500)
			return
		}
	}

}
