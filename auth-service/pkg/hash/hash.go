package hash

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

func GenerateHash(password string) []byte {
	hash := sha256.New()
	hash.Write([]byte(password))
	return hash.Sum(nil)
}

func ComparePasswords(password string, hashedPassword []byte) bool {
	receivedHash := fmt.Sprintf("%x", GenerateHash(password))
	return bytes.Equal([]byte(receivedHash), hashedPassword)
}
