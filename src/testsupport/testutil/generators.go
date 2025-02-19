package testutil

import (
	"math/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RandSeqPointer(n int) *string {
	random := RandSeq(n)
	return &random
}

func RandListStringPointer(n int) []string {
	b := make([]string, n)
	for i := range b {
		b[i] = RandSeq(10)
	}
	return b
}

func RanNumberPointer() *int32 {
	i := rand.Int31()
	return &i
}
