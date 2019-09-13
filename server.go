package main

import (
	"fmt"
	"net/http"
)

func ProfileServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Peter's Profile")
}
