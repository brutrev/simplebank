package util

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// RandomInt returns, as an int64, a non-negative pseudo-random number between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString returns a string with length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// GenerateOwner returns, as a string, a random Owner name
func GenerateOwner() string {
	return RandomString(6)
}

// GenerateMoney returns, as an int64, a random amount of money
func GenerateMoney() int64 {
	return RandomInt(0, 1000)
}

func GenerateCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
