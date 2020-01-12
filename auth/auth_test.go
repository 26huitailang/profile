package auth_test

import (
	"github.com/stretchr/testify/assert"
	"profile/auth"
	"testing"
)

func TestPasswordCheck(t *testing.T) {
	cases := []struct {
		name     string
		password string
		want     error
	}{
		{name: "ok", password: "123Aa&", want: nil},
		{name: "invalid code", password: "123Aa&)", want: auth.UnsupportedPasswordCodeError},
		{name: "too short", password: "1Aa&)", want: auth.TooShortPasswordError},
		{name: "too simple", password: "123123123aa", want: auth.TooSimplePasswordError},
	}
	for _, tt := range cases {
		got := auth.PasswordCheck(tt.password)
		assert.Equal(t, tt.want, got)
	}
}
