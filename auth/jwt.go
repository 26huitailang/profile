package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// JwtCustomClaims are custom claims extending default ones.
type JwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}


func GenJWT(name string, isAdmin bool, key []byte, exp time.Duration) (string, error) {

	// Set claims
	claims := &JwtCustomClaims{
		name,
		isAdmin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(exp).Unix(),
		},
	}
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString(key)

	return t, err
}

func ParseToken(tokenString string, key []byte) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("method not right, should be %v", t.Method)
		}
		return key, nil
	})
	return token, err
}

func ParseClaims(token *jwt.Token) *JwtCustomClaims {
	if claims, ok := token.Claims.(*JwtCustomClaims); ok && token.Valid {
		return claims
	}
	return nil
}