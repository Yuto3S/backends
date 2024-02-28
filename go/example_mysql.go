package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func initTables() {
	db, err := sql.Open("mysql", "root:test-go-mysql@tcp(127.0.0.1:3306)/testgo?parseTime=true")
	if err != nil {
		log.Fatal(err)
		defer db.Close()
	}

	err_maybe := db.Ping()
	fmt.Println("Ping: ", err, err_maybe)

	if !usersTableExists(db) {
		createUserTable(db)
	}

	insertRandomUser(db)
}

func usersTableExists(db *sql.DB) bool {
	exists, _ := db.Exec(`SELECT * FROM users`)
	return exists != nil
}

func createUserTable(db *sql.DB) {
	create_users_table_query := `
		CREATE TABLE users (
			id INT AUTO_INCREMENT,
			username TEXT NOT NULL,
			password TEXT NOT NULL,
			created_at DATETIME,
			PRIMARY KEY (id)
		)
	`

	_, err := db.Exec(create_users_table_query)
	if err != nil {
		log.Fatal(err)
		defer db.Close()
	}

}

func dropUsersTable(db *sql.DB) {
	delete_users_table_qeury := `DROP TABLE users`
	_, err := db.Exec(delete_users_table_qeury)
	if err != nil {
		log.Fatal(err)
		defer db.Close()
	}
}

// https://github.com/random-names/go
func insertRandomUser(db *sql.DB) {
	username := "johndoe"
	password := "secret"
	createdAt := time.Now()

	result, _ := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
	userID, _ := result.LastInsertId()
	fmt.Println("USER ID INSERTED:", userID)
}
