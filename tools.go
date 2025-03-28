package toolkit

import (
	mathrand "math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandomString generates a random string of the specified length.
func RandomString(length int) string {
	mathrand := mathrand.New(mathrand.NewSource(time.Now().UnixNano()))
	mathrand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[mathrand.Intn(len(charset))]
	}
	return string(b)
}
