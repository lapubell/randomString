package main

import (
	"flag"
	"math/rand"
	"time"

	"github.com/atotto/clipboard"
)

type charRune []rune

var chars = charRune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var specs = charRune("!@#$%^&*(){}[]|+")
var nums = charRune("01234567890123456789")

func main() {
	lengthPtr := flag.Int("len", 32, "String length")
	specialsPtr := flag.Bool("spec", false, "String has special characters=")
	numsPtr := flag.Bool("num", true, "String has numbers")

	flag.Parse()

	var output = make(charRune, *lengthPtr)

	if *numsPtr {
		chars = append(chars, nums...)
	}

	if *specialsPtr {
		chars = append(chars, specs...)
	}

	chars.shuffle()

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < *lengthPtr; i++ {
		output[i] = chars[rand.Intn(len(chars)-1)]
	}
	clipboard.WriteAll(string(output))
}

func (c charRune) shuffle() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := range c {
		rand := r.Intn(len(c) - 1)
		c[i], c[rand] = c[rand], c[i]
	}
}
