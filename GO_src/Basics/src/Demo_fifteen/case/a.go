package main

import (
	"fmt"
	"math/rand"
	"time"
)

var defaultLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandomString returns a random string with a fixed length
func RandomString(n int, allowedChars string) string {

	b := make([]byte,n)
	for i := range b {
		b[i] = allowedChars[rand.Intn(len(allowedChars))]
	}

	return string(b)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(RandomString(6, defaultLetters))
}
