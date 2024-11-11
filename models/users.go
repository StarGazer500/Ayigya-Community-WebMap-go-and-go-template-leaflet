package models

// const UserSQLModel = `
// CREATE TABLE IF NOT EXISTS user (
// 	id            		 INTEGER NOT NULL PRIMARY KEY,
// 	firstname         		 TEXT NOT NULL UNIQUE,
// 	surname 	  		 TEXT NOT NULL,
// 	password1  TEXT NOT NULL

// );`
// // CREATE INDEX IF NOT EXISTS user_email_idx ON user(email);
// CREATE INDEX IF NOT EXISTS user_active_idx ON user(active);`

type UserTable struct {
	TableName string
}

var UserModel *UserTable

func CreateUserTable() {
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS "UserSQLModel" (
			id SERIAL PRIMARY KEY,
			firstname VARCHAR(100),
			surname VARCHAR(100),
			password1 VARCHAR(100),
			email VARCHAR(100)
		);
	`

	UserModel = &UserTable{TableName: "UserSQLModel"}
	CreateTable(createTableQuery)

}
