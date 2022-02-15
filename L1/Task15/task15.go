package main

import (
	"math/rand"
)

var justString string

func createHugeString(n int) string {
	var str string
	for i := 0; i < n; i++ {
		str+=string(byte(rand.Intn(26)+65))
	}
	return str
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
}