package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// Make sure that every time we run the code, the generated values will be different.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Random int64 number between min and max.
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// Random string of length characters.
func RandomString(length int) string {
	var sb strings.Builder
	num := len(alphabet)

	for i := 0; i < length; i++ {
		char := alphabet[rand.Intn(num)]
		sb.WriteByte(char)
	}

	return sb.String()
}

// Generates a random item name.
func RandomName() string {
	return RandomString(8)
}

// Generates a random item price.
func RandomPrice() int64 {
	return RandomInt(0, 500)
}
