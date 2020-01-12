package auth

import (
	"bytes"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

var TooShortPasswordError = errors.New("password too short")
var UnsupportedPasswordCodeError = errors.New("unsupported code")
var TooSimplePasswordError = errors.New("too simple password")

// PasswordCheck to check password include at least three of them:
//   - capital letter
//   - lower letter
//   - one of {'!', '@', '#', '$', '%', '^', '&', '*', '_', '-'}
//   - number
func PasswordCheck(password string) error {
	indNum := [4]int{0, 0, 0, 0}
	spCode := []byte{'!', '@', '#', '$', '%', '^', '&', '*', '_', '-'}

	if len(password) < 6 {
		return TooShortPasswordError
	}

	passwdByte := []byte(password)

	for _, i := range passwdByte {

		if i >= 'A' && i <= 'Z' {
			indNum[0] = 1
			continue
		}

		if i >= 'a' && i <= 'z' {
			indNum[1] = 1
			continue
		}

		if i >= '0' && i <= '9' {
			indNum[2] = 1
			continue
		}

		notEnd := 0
		for _, s := range spCode {
			if i == s {
				indNum[3] = 1
				notEnd = 1
				break
			}
		}

		if notEnd != 1 {
			return UnsupportedPasswordCodeError
		}

	}

	codeCount := 0

	for _, i := range indNum {
		codeCount += i
	}

	if codeCount < 3 {
		return TooSimplePasswordError
	}

	return nil
}

func EncryptPassword(password string) string {
	var buf bytes.Buffer
	salt := []byte("$hello$")
	pwd := []byte(password)
	buf.Write(salt)
	buf.Write(pwd)
	buf.Write(salt)
	pwd = buf.Bytes()
	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}

func Login(username, password string) bool {
	return true
}
