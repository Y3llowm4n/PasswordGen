// connect with database
func ConnectDB() {
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(192.168.180.17:3306)/"+dbname)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Succesfully connected!")
}