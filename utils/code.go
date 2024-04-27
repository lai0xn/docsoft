package utils

import (
	"math/rand"
)

const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateCode() string {
	// Seed the random number generator

	// Initialize a byte slice to hold the generated code
	code := make([]byte, 7)

	// Populate the byte slice with random characters from the charset
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}

	// Convert the byte slice to a string
	return string(code)
}
