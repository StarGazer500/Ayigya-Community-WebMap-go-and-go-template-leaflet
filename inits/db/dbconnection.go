package db

import (
	"database/sql"
	"fmt"

	// "reflect"

	_ "github.com/lib/pq"
)

type DbInstance struct {
	Db *sql.DB
}

var PG *DbInstance

func ConnectTODb() error {
	connStr := "user=postgres password=0549martin dbname=ayigyadb sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Something happpened", err)
		return err
	}

	PG = &DbInstance{Db: db}

	// fmt.Println("This is  the db instance%", PG.Db)

	return nil
}
