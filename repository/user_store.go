package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"

	"app/user"
)

type UserStore interface {
	Store(user user.User)
}

type Repository struct {
	db *sql.DB
}

func Store(user user.User) {
	db := getDb()

	// insert
	stmt, err := db.Prepare("INSERT INTO users(username) VALUES (username=?)")
	checkErr(err)

	_, err = stmt.Exec(user.Username)
	checkErr(err)
}

func Find(username string) user.User {
	db := getDb()

	// insert
	stmt, err := db.Prepare("SELECT * FROM users WHERE username=?")
	checkErr(err)

	_, err = stmt.Exec(username)
	checkErr(err)

	return user.User{Username: username}
}

func getDb() (*sql.DB) {
	conString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("USER"),
		os.Getenv("PASSWORD"),
		os.Getenv("HOST"),
		os.Getenv("PORT"),
		os.Getenv("DATABASE"),
	)
	db, err := sql.Open("mysql", conString)
	checkErr(err)

	return db
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
