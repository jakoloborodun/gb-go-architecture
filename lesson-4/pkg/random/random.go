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

// Random int number between min and max.
func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
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

// Generates a random integer slice.
func RandomIntSlice(length int) []int {
	var intSlice []int

	for i := 0; i < length; i++ {
		intSlice = append(intSlice, RandomInt(1, 500))
	}
	return intSlice
}
