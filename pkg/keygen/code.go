package keygen

import (
	"math/rand"
)

func CodeGenerator(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ)(*$!&")
	if n <= 0 {
		n = rand.Int()
	}
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
	// return rand.Seed(time.Now().UnixNano())
}
