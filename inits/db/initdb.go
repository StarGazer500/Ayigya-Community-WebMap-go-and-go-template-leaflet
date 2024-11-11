package db

import (
	"fmt"

	_ "github.com/lib/pq"
)

func InitpgDb() {

	err := ConnectTODb()
	if err != nil {
		// If there was an error, log it and exit
		fmt.Println("Error connecting to the database:", err)
		return
	}

	// fmt.Println("This is the PG assessing", PG.Db)

	// If no error, use the db connection (example)

	// fmt.Println("Connected to the database successfully!",value.Db)

	// Don't forget to close the database when don

}
