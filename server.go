package main

import (
	"fmt"
	"net/http"
)

func ProfileServer(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Path[len("/profiles/"):]

	fmt.Fprint(w, GetUserProfile(username))
}

func GetUserProfile(username string) string {
	if username == "Peter" {
		return "Peter's Profile"

	}
	if username == "Chris" {
		return "Chris's Profile"
	}

	return ""
}