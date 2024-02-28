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

	userIdInserted := insertRandomUser(db)
	selectUser(userIdInserted, db)
	getAllUsers(db)
	deleteUser(userIdInserted, db)
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
func insertRandomUser(db *sql.DB) int64 {
	username := "johndoe"
	password := "secret"
	createdAt := time.Now()

	result, _ := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
	userID, _ := result.LastInsertId()
	fmt.Println("USER ID INSERTED:", userID)
	return userID
}

func selectUser(userId int64, db *sql.DB) {
	var (
		id        int
		username  string
		password  string
		createdAt time.Time
	)

	query := `SELECT id, username, password, created_at FROM users WHERE id = ?`
	err := db.QueryRow(query, userId).Scan(&id, &username, &password, &createdAt)
	fmt.Println("User id ", userId, " has values: ", username, password, createdAt)
	if err != nil {
		fmt.Println("error select user ", err)
	}
}

func getAllUsers(db *sql.DB) {
	type user struct {
		id        int
		username  string
		password  string
		createdAt time.Time
	}

	rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
	if err != nil {
		fmt.Println("error get all users ", err)
	}
	defer rows.Close()

	var users []user
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
		if err != nil {
			fmt.Println("error row scan ", err)
		}
		users = append(users, u)
	}
	errr := rows.Err()
	if errr != nil {
		fmt.Println("error in rows ", errr)
	}

	for index, user := range users {
		fmt.Println(index, user.username, user.password, user.createdAt)
	}
}

func deleteUser(userId int64, db *sql.DB) {
	delete_user_query := `DELETE FROM users WHERE id = ?`
	result, err := db.Exec(delete_user_query, userId)
	if err != nil {
		fmt.Println("Error deleting user ", err)

	}

	if result != nil {
		rowsAffected, _ := result.RowsAffected()
		fmt.Println("Number of user rows deleted: ", rowsAffected)
	}
}
