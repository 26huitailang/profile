package auth

import (
	"testing"
	"time"
)

func TestGenJWT(t *testing.T) {
	type args struct {
		name    string
		isAdmin bool
		key     []byte
		exp     time.Duration
		gap time.Duration
	}
	type want struct {
		name    string
		isAdmin bool
		expire bool
	}
	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{name: "token expire", args: args{"admin", true, []byte("secret"), time.Millisecond * 100, 0}, want: want{"admin", true, true}, wantErr: false},
		{name: "token out of expire", args: args{"admin", false, []byte("secret"), time.Millisecond * 1, time.Millisecond * 1000}, want: want{"admin", false, false}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenStr, err := GenJWT(tt.args.name, tt.args.isAdmin, tt.args.key, tt.args.exp)
			token, err := ParseToken(tokenStr, tt.args.key)
			time.Sleep(tt.args.gap)
			claims := ParseClaims(token)
			expireErr := claims.Valid()
			expireValid := true
			if expireErr != nil {
				expireValid = false
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("GenJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got := struct {
				name string
				isAdmin bool
				expire bool
			}{
				name: claims.Name,
				isAdmin: claims.Admin,
				expire: expireValid,
			}
			if got != tt.want {
				t.Errorf("GenJWT() got = %v, want %v", got, tt.want)
			}
		})
	}
}
