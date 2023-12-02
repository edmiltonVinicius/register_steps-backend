package utils

import "golang.org/x/crypto/bcrypt"

func GenerateHashString(data string, cost int) (h string, e error) {
	b, e := bcrypt.GenerateFromPassword([]byte(data), cost)
	h = string(b)
	return h, e
}

func CheckHash(hash, data string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(data))
	return err == nil
}
