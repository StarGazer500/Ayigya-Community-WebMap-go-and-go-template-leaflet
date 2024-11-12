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

func InsertOne(db *sql.DB, tablename string, columns []string, args ...interface{}) (sql.Result, error) {
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

// InsertMany inserts multiple rows into the specified table.
func InsertMany(db *sql.DB, tablename string, columns []string, rows [][]interface{}) (sql.Result, error) {
	// Check if columns are provided
	if len(columns) == 0 {
		return nil, fmt.Errorf("columns must not be empty")
	}

	// Ensure that each row has the same number of values as there are columns
	for _, row := range rows {
		if len(row) != len(columns) {
			return nil, fmt.Errorf("each row must have the same number of values as the number of columns")
		}
	}

	// Dynamically build the column part for the query
	cols := strings.Join(columns, ", ")

	// Build the placeholders part for multiple rows (e.g. "($1, $2, $3), ($4, $5, $6), ...")
	var placeholders []string
	for i := range rows {
		// Generate the placeholders for each row
		rowPlaceholders := make([]string, len(columns))
		for j := range rowPlaceholders {
			rowPlaceholders[j] = fmt.Sprintf("$%d", i*len(columns)+j+1)
		}
		placeholders = append(placeholders, fmt.Sprintf("(%s)", strings.Join(rowPlaceholders, ", ")))
	}

	// Create the query string
	placeholderStr := strings.Join(placeholders, ", ")
	query := fmt.Sprintf(`INSERT INTO "%s" (%s) VALUES %s`, tablename, cols, placeholderStr)

	// Flatten the rows into a single slice of arguments
	var args []interface{}
	for _, row := range rows {
		args = append(args, row...)
	}

	// Execute the query with the provided arguments
	result, err := db.Exec(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error inserting data into the database: %v", err)
	}

	// Get the number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("error getting rows affected: %v", err)
	}

	// Print meaningful details
	fmt.Printf("Rows affected: %d\n", rowsAffected)
	return result, nil
}

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

// AddColumnIfNotExists checks if the column exists in the table, and if not, adds it.
func AddColumnIfNotExists(db *sql.DB, tableName, columnName, columnType string) (string, error) {
	// Check if the column already exists in the table
	query := `
		SELECT column_name
		FROM information_schema.columns
		WHERE table_name = $1 AND column_name = $2;
	`

	var existingColumn string
	err := db.QueryRow(query, tableName, columnName).Scan(&existingColumn)
	if err != nil && err != sql.ErrNoRows {
		// If there's an error other than no rows found, log and return the error
		return "", fmt.Errorf("error checking column existence: %w", err)
	}

	if existingColumn == columnName {
		// If the column exists, no need to add it
		fmt.Printf("Column '%s' already exists in table '%s'.\n", columnName, tableName)
	} else {
		// Add the new column if it doesn't exist
		// Here we have to directly interpolate table and column names into the query
		// but make sure the values are sanitized to avoid SQL injection
		// *** Be cautious with interpolating user inputs directly into SQL ***
		// In this case, we are assuming `tableName` and `columnName` are safe.

		// Build the ALTER TABLE query
		addColumnQuery := fmt.Sprintf(`
			ALTER TABLE "%s"
			ADD COLUMN "%s" %s;
		`, tableName, columnName, columnType)

		// Execute the ALTER TABLE query
		_, err := db.Exec(addColumnQuery)
		if err != nil {
			return "", fmt.Errorf("error adding column '%s' to table '%s': %w", columnName, tableName, err)
		}

		fmt.Printf("Column '%s' added successfully to table '%s'.\n", columnName, tableName)
	}

	// Generate the updated CREATE TABLE SQL syntax with all columns
	createTableSQL, err := generateCreateTableSQL(db, tableName)
	if err != nil {
		return "", fmt.Errorf("error generating CREATE TABLE SQL: %w", err)
	}

	// Return the generated CREATE TABLE SQL
	return createTableSQL, nil
}

// DeleteColumnIfExists checks if the column exists in the table, and if so, deletes it.
func DeleteColumnIfExists(db *sql.DB, tableName, columnName string) (string, error) {
	// Check if the column already exists in the table
	query := `
		SELECT column_name
		FROM information_schema.columns
		WHERE table_name = $1 AND column_name = $2;
	`

	var existingColumn string
	err := db.QueryRow(query, tableName, columnName).Scan(&existingColumn)
	if err != nil && err != sql.ErrNoRows {
		// If there's an error other than no rows found, log and return the error
		return "", fmt.Errorf("error checking column existence: %w", err)
	}

	if existingColumn == columnName {
		// If the column exists, delete it
		deleteColumnQuery := fmt.Sprintf(`
			ALTER TABLE "%s"
			DROP COLUMN "%s";
		`, tableName, columnName)

		// Execute the ALTER TABLE query to drop the column
		_, err := db.Exec(deleteColumnQuery)
		if err != nil {
			return "", fmt.Errorf("error deleting column '%s' from table '%s': %w", columnName, tableName, err)
		}

		fmt.Printf("Column '%s' deleted successfully from table '%s'.\n", columnName, tableName)
	} else {
		// If the column does not exist, no action needed
		fmt.Printf("Column '%s' does not exist in table '%s'. No action taken.\n", columnName, tableName)
	}

	// Generate the updated CREATE TABLE SQL syntax with all remaining columns
	createTableSQL, err := generateCreateTableSQL(db, tableName)
	if err != nil {
		return "", fmt.Errorf("error generating CREATE TABLE SQL: %w", err)
	}

	// Return the generated CREATE TABLE SQL
	return createTableSQL, nil
}

// generateCreateTableSQL generates the SQL query Syntz to create a table
// based on the columns from an existing table in the database.
func generateCreateTableSQL(db *sql.DB, tableName string) (string, error) {
	// Query to get all columns and their types
	query := `
		SELECT column_name, data_type, character_maximum_length
		FROM information_schema.columns
		WHERE table_name = $1
		ORDER BY ordinal_position;
	`

	rows, err := db.Query(query, tableName)
	if err != nil {
		return "", fmt.Errorf("error retrieving columns from information_schema: %w", err)
	}
	defer rows.Close()

	// Start building the CREATE TABLE SQL query
	var createTableSQL = fmt.Sprintf("CREATE TABLE \"%s\" (\n", tableName)

	// Loop through the columns and construct the SQL for each
	firstColumn := true
	for rows.Next() {
		var columnName, dataType string
		var maxLength *int

		err := rows.Scan(&columnName, &dataType, &maxLength)
		if err != nil {
			return "", fmt.Errorf("error scanning row: %w", err)
		}

		// Add column definition to the CREATE TABLE SQL
		if !firstColumn {
			createTableSQL += ",\n"
		}

		// Handle data types (e.g., VARCHAR with length)
		if dataType == "character varying" && maxLength != nil {
			createTableSQL += fmt.Sprintf("\t%s %s(%d)", columnName, dataType, *maxLength)
		} else {
			createTableSQL += fmt.Sprintf("\t%s %s", columnName, dataType)
		}

		firstColumn = false
	}

	// Check for any errors after the loop
	if err := rows.Err(); err != nil {
		return "", fmt.Errorf("error iterating over rows: %w", err)
	}

	// Close the CREATE TABLE statement
	createTableSQL += "\n);\n"

	// Return the generated SQL
	return createTableSQL, nil
}

func DeleteRowByColumn(db *sql.DB, tableName, columnName, columnValue string) error {
	// Step 1: Check if the row exists by querying the table for the column and value
	query := fmt.Sprintf(`
		SELECT COUNT(*)
		FROM "%s"
		WHERE "%s" = $1;
	`, tableName, columnName)

	var count int
	err := db.QueryRow(query, columnValue).Scan(&count)
	if err != nil {
		return fmt.Errorf("error checking row existence: %w", err)
	}

	if count == 0 {
		// If the row does not exist, return an error message
		return fmt.Errorf("row with %s = '%s' not found in table '%s'", columnName, columnValue, tableName)
	}

	// Step 2: If the row exists, delete it using the DELETE statement
	deleteQuery := fmt.Sprintf(`
		DELETE FROM "%s"
		WHERE "%s" = $1;
	`, tableName, columnName)

	// Execute the DELETE query
	_, err = db.Exec(deleteQuery, columnValue)
	if err != nil {
		return fmt.Errorf("error deleting row where %s = '%s' from table '%s': %w", columnName, columnValue, tableName, err)
	}

	fmt.Printf("Row with %s = '%s' deleted successfully from table '%s'.\n", columnName, columnValue, tableName)
	return nil
}

func DeleteRowByID(db *sql.DB, tableName string, id int) error {
	// Step 1: Check if the row exists by querying the table for the id
	query := fmt.Sprintf(`
		SELECT COUNT(*)
		FROM "%s"
		WHERE id = $1;
	`, tableName)

	var count int
	err := db.QueryRow(query, id).Scan(&count)
	if err != nil {
		return fmt.Errorf("error checking row existence: %w", err)
	}

	if count == 0 {
		// If the row does not exist, return a message
		return fmt.Errorf("row with id %d not found in table '%s'", id, tableName)
	}

	// Step 2: If the row exists, delete it using the DELETE statement
	deleteQuery := fmt.Sprintf(`
		DELETE FROM "%s"
		WHERE id = $1;
	`, tableName)

	// Execute the DELETE query
	_, err = db.Exec(deleteQuery, id)
	if err != nil {
		return fmt.Errorf("error deleting row with id %d from table '%s': %w", id, tableName, err)
	}

	fmt.Printf("Row with id %d deleted successfully from table '%s'.\n", id, tableName)
	return nil
}

// UpdateOne updates a single row in the specified table with the provided column values.
func UpdateOne(db *sql.DB, tablename string, columns []string, args []interface{}, whereColumn string, whereValue interface{}) (sql.Result, error) {
	// Check if columns and arguments match
	if len(columns) != len(args) {
		return nil, fmt.Errorf("the number of columns and values must match")
	}

	// Dynamically build the column=value part for the query (e.g. "col1 = $1, col2 = $2")
	setClauses := make([]string, len(columns))
	for i := range columns {
		setClauses[i] = fmt.Sprintf("\"%s\" = $%d", columns[i], i+1)
	}
	setClauseStr := strings.Join(setClauses, ", ")

	// Create the query string with the WHERE clause for the row to update
	query := fmt.Sprintf(`UPDATE "%s" SET %s WHERE "%s" = $%d`, tablename, setClauseStr, whereColumn, len(columns)+1)

	// Combine the arguments (the column values and the WHERE condition)
	args = append(args, whereValue)

	// Execute the query with the provided arguments
	result, err := db.Exec(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error updating data in the database: %v", err)
	}

	// Get the number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("error getting rows affected: %v", err)
	}

	// Print meaningful details
	fmt.Printf("Rows affected: %d\n", rowsAffected)
	return result, nil
}
