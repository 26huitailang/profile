package auth

import (
	"testing"
	"time"
)

func TestGenJWT(t *testing.T) {
	type args struct {
		name string
		key  []byte
		exp  time.Duration
		gap  time.Duration
	}
	type want struct {
		name   string
		expire bool
	}
	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name:    "token expire",
			args:    args{"admin", []byte("secret"), time.Millisecond * 100, 0},
			want:    want{"admin", true},
			wantErr: false,
		},
		{
			name:    "token out of expire",
			args:    args{"admin", []byte("secret"), time.Millisecond * 1, time.Millisecond * 1000},
			want:    want{"admin", false},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenStr, err := GenJWT(tt.args.name, tt.args.key, tt.args.exp)
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
				name   string
				expire bool
			}{
				name:   claims.Name,
				expire: expireValid,
			}
			if got != tt.want {
				t.Errorf("GenJWT() got = %v, want %v", got, tt.want)
			}
		})
	}
}
