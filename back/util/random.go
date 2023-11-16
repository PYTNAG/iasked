package util

import (
	"fmt"
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz_()!$"

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomUsername generates a random string of length 24
func RandomUsername() string {
	return RandomString(24)
}

// RandomPassword generates a random string of length 16
func RandomPassword() string {
	return RandomString(16)
}

// RandomID generates a random positive int32
func RandomID() int32 {
	res := rand.Int31()
	if res == int32(0) {
		res = 1
	}

	return res
}

// RandomEmail generates a random email
func RandomEmail() string {
	mailbox := RandomString(10)
	mailServer := RandomString(10)
	domain := RandomString(3)

	return fmt.Sprintf("%s@%s.%s", mailbox, mailServer, domain)
}
