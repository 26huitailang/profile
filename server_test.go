package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETProfile(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	ProfileServer(response, request)

	t.Run("returns Peter's profile", func(t *testing.T) {
		got := response.Body.String()
		want := "Peter's Profile"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
