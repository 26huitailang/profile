package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETProfile(t *testing.T) {
	t.Run("returns Peter's profile", func(t *testing.T) {
		request := newGetProfileRequest("Peter")
		response := httptest.NewRecorder()

		ProfileServer(response, request)
		got := response.Body.String()
		want := "Peter's Profile"

		assertResponseBody(t, got, want)
	})

	t.Run("returns Chris's profile", func(t *testing.T) {
		request := newGetProfileRequest("Chris")
		response := httptest.NewRecorder()

		ProfileServer(response, request)
		got := response.Body.String()
		want := "Chris's Profile"

		assertResponseBody(t, got, want)
	})
}

func newGetProfileRequest(username string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/profiles/%s", username), nil)
	return request
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q, want %q", got, want)
	}
}