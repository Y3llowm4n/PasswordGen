// input wachtwoord geven
// functie wachtwoord genereren
// functie hashing

// username toevoegen
// username duplicate error
// database connecten met file
// error schrijven naar logfile
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

// global variable which includes all the letters, numbers and tokens.
var (
	letters        = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()")
	lengthPassword int
)

func init() {
	flag.IntVar(&lengthPassword, "p", 12, "Give a value for the length off your password")
	flag.Parse()
}

func main() {
	username := getUser()
	fmt.Println("Name:", username)
	genPassword()
}

func getUser() string {
	var username string
	fmt.Println("Please give a username")
	fmt.Scanln(&username)
	return username
}
func genPassword() {

	rand.Seed(time.Now().UnixNano())

	str1 := "abcdefghijklmnopqrstuvwxyz"
	str2 := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	str3 := "1234567890!@#$%^&*()"

	str := str1 + str2 + str3

	p := make([]rune, lengthPassword)

	for i := range p {
		p[i] = letters[rand.Intn(len(str))]
	}

	fmt.Println("Password:", string(p))
}
