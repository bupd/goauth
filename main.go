package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("kumar")
	http.HandleFunc("/", userHandler)
	http.ListenAndServe("", nil)
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
