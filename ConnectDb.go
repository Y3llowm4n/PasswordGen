package ConnectDb

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type config struct {
	dbname string
	dbuser string
	dbpass string
	host   int
}

// connect with database
func ConnectDB(dbname string, dbuser string, dbpass string, host int) {
	jsonFile, err := os.Open("/passwordgen/config.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened config.json")
	defer jsonFile.Close()
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(192.168.180.17:3306)/"+dbname)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Succesfully connected!")
}
