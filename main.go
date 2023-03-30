// add max length of password
// randomize position of characters
// error logging
// username duplicate error
// database connecten met file
// error schrijven naar logfile
package main

import (
	"bytes"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

// global variable which includes all the letters, numbers and tokens.
var (
	lowerCount     int
	upperCount     int
	specialCount   int
	lengthPassword int
)

const (
	dbname = "data"
	dbuser = "student"
	dbpass = "admin01"
)

// here's all the input that the user can specify.
func init() {
	flag.IntVar(&lowerCount, "l", 2, "Number of lowercase letters in password")
	flag.IntVar(&upperCount, "u", 3, "Number of uppercase letters in password")
	flag.IntVar(&specialCount, "s", 4, "Number of special characters in password")
	flag.IntVar(&lengthPassword, "p", 12, "Give a value for the length off your password")
	flag.Parse()
}

func main() {
	fmt.Println(dbpass)
	username := getUser()
	fmt.Println("Username:", username)
	password := genPassword(lowerCount, upperCount, specialCount)
	fmt.Println("Password:", password)
	ConnectDB()
}

// connect with database
func ConnectDB() {
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(192.168.180.17:3306)/"+dbname)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Succesfully connected!")
}

// Give username for log in.
func getUser() string {
	var username string
	fmt.Println("Please give a username")
	fmt.Scanln(&username)
	return username
}

// Generate customized password with given inputs by user.
func genPassword(lowerCount, upperCount, specialCount int) string {

	lowercase := []rune("abcdefghijklmnopqrstuvwxyz")
	uppercase := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	special := []rune("1234567890~!@#$%^&*()")

	rand.Seed(time.Now().UnixNano())

	//save the password inside the buffer
	var buffer bytes.Buffer

	//add the given amount of lowercase characters.
	for i := 0; i < lowerCount; i++ {
		buffer.WriteRune(lowercase[rand.Intn(len(lowercase))])
	}

	//add the given amount of uppercase characters.
	for i := 0; i < upperCount; i++ {
		buffer.WriteRune(uppercase[rand.Intn(len(uppercase))])
	}

	//add the given amount of special characters.
	for i := 0; i < specialCount; i++ {
		buffer.WriteRune(special[rand.Intn(len(special))])
	}

	// adds random character from the three options to meet the required length.
	for buffer.Len() < lengthPassword {
		switch rand.Intn(3) {
		case 0:
			buffer.WriteRune(lowercase[rand.Intn(len(lowercase))])
		case 1:
			buffer.WriteRune(uppercase[rand.Intn(len(uppercase))])
		case 2:
			buffer.WriteRune(special[rand.Intn(len(special))])
		}
	}

	return buffer.String()
}
