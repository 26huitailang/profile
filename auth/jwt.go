package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"time"
)

// JwtCustomClaims are custom claims extending default ones.
type JwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func GetUserInfo(c echo.Context) JwtCustomClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(JwtCustomClaims)

	return claims
}

func GenJWT(name string, isAdmin bool, key []byte) (string, error) {

	// Set claims
	claims := &JwtCustomClaims{
		Name:  name,
		Admin: isAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 6).Unix(),
		},
	}
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString(key)

	return t, err
}
