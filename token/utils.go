package token

import (
	"math/rand"
	"strings"
)

// RandomString generates random string.
func RandomString(length int) string {
	var sb strings.Builder
	k := len(englishAlphabet)

	for i := 0; i < length; i++ {
		c := englishAlphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}
