package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETProfile(t *testing.T) {
	t.Run("returns Peter's profile", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/profiles/Peter", nil)
		response := httptest.NewRecorder()

		ProfileServer(response, request)
		got := response.Body.String()
		want := "Peter's Profile"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns Chris's profile", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/profiles/Chris", nil)
		response := httptest.NewRecorder()

		ProfileServer(response, request)
		got := response.Body.String()
		want := "Chris's Profile"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
