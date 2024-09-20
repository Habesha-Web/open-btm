package utils

import (
	"math/rand"
)

const (
	charsetLen = 26
	charset    = "abcdefghijklmnopqrstuvwxyz"
)

func GenerateRandomString(length int) (string, error) {
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		index := rand.Intn(charsetLen)
		result[i] = charset[index]
	}

	return string(result), nil
}
