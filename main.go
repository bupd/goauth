package main

import (
	"fmt"
	"net/http"

	"github.com/bupd/goauth/users"
)

func main() {
	fmt.Printf("kumar")
	http.HandleFunc("/", userHandler)
	http.ListenAndServe("", nil)
}

func getSignInPage(w http.ResponseWriter, r *http.Request) {
}

func getSignUpPage(w http.ResponseWriter, r *http.Request) {
}

func logout(w http.ResponseWriter, r *http.Request) {
}

func getUser(r *http.Request) users.User {
	email := r.FormValue("email")
	password := r.FormValue("password")

	return users.User{
		Email:    email,
		Password: password,
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/signin":
		getSignInPage(w, r)
	case "/signup":
		getSignUpPage(w, r)
	case "/logout":
		logout(w, r)
	}
	fmt.Printf("write: %s, read: %v", w, r)
}
