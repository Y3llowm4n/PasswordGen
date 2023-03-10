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
	//lengthPassword int
	Lower   int
	Upper   int
	Special int
)

func init() {
	flag.IntVar(&Lower, "l", 4, "Give number for amount of lowercase letters")
	flag.IntVar(&Upper, "u", 4, "Give number for amount of uppercase letters")
	flag.IntVar(&Special, "s", 4, "Give number for amount of special characters")
	//flag.IntVar(&lengthPassword, "p", 12, "Give a value for the length off your password")
	flag.Parse()

	// if lengthPassword < 8 || lengthPassword > 64 {
	// 	log.Fatal("Password length must be between 8 and 64 characters")
	// }
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

func genPassword() string {

	rand.Seed(time.Now().UnixNano())

	lowercaseLetters := "abcdefghijklmnopqrstuvwxyz"
	uppercaseLetters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharacters := "1234567890!@#$%^&*()"

	password := lowercaseLetters + uppercaseLetters + specialCharacters

	return password
}
