package models

import (
	"Ayigya-Community-WebMap-go-and-go-template-geoserver-leaflet/inits/db"
	"database/sql"
	"strings"

	"fmt"
)

func CreateTable(querystring string) {

	tab, err := db.PG.Db.Exec(querystring)

	if err != nil {
		fmt.Println("Error creating table: ", err)
		return
	}

	fmt.Println("Table created successfully (if it didn't already exist).", tab)

}

// func Insert(db *sql.DB,tablename string, query string, args ...interface{}) (sql.Result, error) {
// 	// Execute the query with the provided arguments
// 	query := `INSERT INTO "%s" (firstname, surname, password1) VALUES ($1, $2, $3)`
// 	query = fmt.Sprintf(query, tablename)
// 	data, err := db.Exec(query, args...)
// 	if err != nil {
// 		return nil, fmt.Errorf("error saving data to the database: %v", err)
// 	}

// 	if err != nil {
// 		return nil, fmt.Errorf("error saving data to the database: %v", err)
// 	}

// 	// Extract the number of rows affected and the last insert ID (if applicable)
// 	rowsAffected, err := data.RowsAffected()
// 	if err != nil {
// 		return nil, fmt.Errorf("error getting rows affected: %v", err)
// 	}

// 	// Print meaningful details
// 	fmt.Printf("Rows affected: %d\n", rowsAffected)
// 	return data, nil
// }

func Insert(db *sql.DB, tablename string, columns []string, args ...interface{}) (sql.Result, error) {
	// Check if columns and arguments match
	if len(columns) != len(args) {
		return nil, fmt.Errorf("the number of columns and values must match")
	}

	// Dynamically build the column part for the query
	cols := strings.Join(columns, ", ")

	// Dynamically build the placeholders for the query (e.g. "$1, $2, $3")
	placeholders := make([]string, len(columns))
	for i := range placeholders {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}
	placeholderStr := strings.Join(placeholders, ", ")

	// Create the query string
	query := fmt.Sprintf(`INSERT INTO "%s" (%s) VALUES (%s)`, tablename, cols, placeholderStr)

	// Execute the query with the provided arguments
	data, err := db.Exec(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error saving data to the database: %v", err)
	}

	// Get the number of rows affected
	rowsAffected, err := data.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("error getting rows affected: %v", err)
	}

	// Print meaningful details
	fmt.Printf("Rows affected: %d\n", rowsAffected)
	return data, nil
}

// func FindOne(db *sql.DB, tableName string, queryitem string, queryfield string) (*sql.Rows, error) {

// 	querystring := "SELECT * FROM %s WHERE %s = '%s'"

// 	query := fmt.Sprintf(querystring, tableName, queryitem, queryfield)

// 	rows, err := db.Query(query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	return rows, fmt.Errorf("user not found")

// 	// for rows.Next() {
// 	// 	var email, password string
// 	// 	err := rows.Scan(&email, &password)
// 	// 	if err != nil {
// 	// 		return nil, err
// 	// 	}

// 	// 	return &Loginresult{
// 	// 		Email: email,
// 	// 		Password:      password,
// 	// 	}, nil
// 	// }

// 	// return nil, fmt.Errorf("User not found")
// }

func FindOne(db *sql.DB, tableName string, queryField string, queryValue string) (*sql.Rows, error) {

	// Correctly format the query string with parameterized query to avoid SQL injection
	querystring := fmt.Sprintf("SELECT * FROM \"%s\" WHERE %s = $1", tableName, queryField)

	// Execute the query with the parameterized query
	rows, err := db.Query(querystring, queryValue)
	if err != nil {
		return nil, fmt.Errorf("query execution failed: %v", err)
	}
	defer rows.Close()

	// Check if the query returned any rows
	if !rows.Next() {
		// If no rows found, return an error
		return nil, fmt.Errorf("user not found with %s = %s", queryField, queryValue)
	}

	// If rows are returned, you can scan the first row (for example, assuming the table has 'email' and 'password' columns)

	// Since we're only looking for one row, return rows as the result
	return rows, nil
}
