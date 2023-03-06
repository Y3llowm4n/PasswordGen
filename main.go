// input wachtwoord geven
// functie wachtwoord genereren
// functie hashing
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// global variable which includes all the letters, numbers and tokens.
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()")

func main() {
	genPassword()
}

func genPassword() {

	rand.Seed(time.Now().UnixNano())

	str1 := "abcdefghijklmnopqrstuvwxyz"
	str2 := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	str3 := "1234567890!@#$%^&*()"

	str := str1 + str2 + str3

	n := 13

	p := make([]rune, n)

	for i := range p {
		p[i] = letters[rand.Intn(len(str))]
	}

	fmt.Println(string(p))
}
