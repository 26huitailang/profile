package v1_test

import (
	"net/http"
	"net/http/httptest"
	v1 "profile/api/v1"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestProfiles(t *testing.T) {
	e := echo.New()

	cases := []struct {
		name     string
		username string
		want     string
	}{
		{name: "returns Peter's Profile", username: "Peter", want: "Peter's Profile"},
		{name: "returns Chris's Profile", username: "Chris", want: "Chris's Profile"},
		{name: "returns None Profile", username: "None", want: ""},
	}

	for _, tt := range cases {
		t.Run("returns one profile", func(t *testing.T) {
			request := httptest.NewRequest(echo.GET, "/profiles/:username", nil)
			response := httptest.NewRecorder()
			h := &v1.ViewHandler{}

			c := e.NewContext(request, response)
			c.SetParamNames("username")
			c.SetParamValues(tt.username)
			h.Profiles(c)

			got := response.Body.String()
			assert.Equal(t, http.StatusOK, response.Code)
			assert.Equal(t, tt.want, got)
		})
	}
}
