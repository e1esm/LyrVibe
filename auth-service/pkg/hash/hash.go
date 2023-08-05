package hash

import "golang.org/x/crypto/bcrypt"

const (
	cost = bcrypt.DefaultCost
)

func GenerateHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
