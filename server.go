package main

import (
	"fmt"
	"net/http"
)

func ProfileServer(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Path[len("/profiles/"):]

	if username == "Peter" {
		fmt.Fprint(w, "Peter's Profile")
	}
	if username == "Chris" {
		fmt.Fprint(w, "Chris's Profile")
	}
}
