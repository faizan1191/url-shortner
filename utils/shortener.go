package utils

import (
	"math/rand"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// GenerateCode func generates a random code of given length

func GenerateCode(n int) string {
	b := make([]byte, n) // create byte slice of given length
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	// convert byte slice to string
	return string(b)
}
