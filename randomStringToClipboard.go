package main

import (
	"flag"
	"math/rand"
	"time"

	"github.com/atotto/clipboard"
)

var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var specs = []rune("!@#$%^&*(){}[]|+")

func main() {
	lengthPtr := flag.Int("len", 32, "Password length")
	specialsPtr := flag.Bool("spec", false, "Password has special characters or just letters and numbers.")

	flag.Parse()

	var output = make([]rune, *lengthPtr)

	if *specialsPtr {
		chars = append(chars, specs...)
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < *lengthPtr; i++ {
		output[i] = chars[rand.Intn(len(chars))]
	}
	clipboard.WriteAll(string(output))
}
