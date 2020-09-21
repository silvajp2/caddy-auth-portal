package utils

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seed *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()),
)

func gen(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seed.Intn(len(charset))]
	}
	return string(b)
}

// GetRandomString returns X character long random string.
func GetRandomString(i int) string {
	if i < 1 {
		i = 40
	}
	return gen(i, charset)
}

// GetRandomStringFromRange generates random string of a random length. The
// random lenght is bounded by a and b.
func GetRandomStringFromRange(a, b int) string {
	var i int
	if a > b {
		i = rand.Intn(a-b) + b
	} else {
		i = rand.Intn(b-a) + a
	}
	return gen(i, charset)
}