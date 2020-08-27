package hash

import (
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(data string) (result string, err error) {
	b := []byte(data)
	hash, err := bcrypt.GenerateFromPassword(b, bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
