package main

import (
	"flag"
	"math/rand"
	"time"

	"github.com/atotto/clipboard"
)

var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var specs = []rune("!@#$%^&*(){}[]|+")
var nums = []rune("0123456789")

func main() {
	lengthPtr := flag.Int("len", 32, "String length")
	specialsPtr := flag.Bool("spec", false, "String has special characters=")
	numsPtr := flag.Bool("num", true, "String has numbers")

	flag.Parse()

	var output = make([]rune, *lengthPtr)

	if *numsPtr {
		chars = append(chars, nums...)
	}

	if *specialsPtr {
		chars = append(chars, specs...)
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < *lengthPtr; i++ {
		output[i] = chars[rand.Intn(len(chars))]
	}
	clipboard.WriteAll(string(output))
}
