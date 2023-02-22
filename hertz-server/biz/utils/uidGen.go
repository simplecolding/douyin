package utils

import (
	"math/rand"
)

func GenToken() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 50)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string((b))
}