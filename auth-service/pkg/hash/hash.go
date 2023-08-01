package hash

import "golang.org/x/crypto/bcrypt"

const (
	cost = 20
)

func GenerateHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func ValidateHash(inputPassword, hashedPassword string) bool {
	password, err := GenerateHash(inputPassword)
	if err != nil {
		return false
	}
	return password == hashedPassword
}
