package repository

import (
	"github.com/BurntSushi/toml"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"../user"
)

type Config struct {
	Database database
}

type database struct {
	Server   string
	Port     string
	Database string
	User     string
	Password string
}

func Store(user user.User) {
	var conf Config
	_, err := toml.DecodeFile("./config.toml", &conf)
	checkErr(err)

	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.Database.User, conf.Database.Password, conf.Database.Server, conf.Database.Port, conf.Database.Database)

	db, err := sql.Open("mysql", connString)
	checkErr(err)

	// insert
	stmt, err := db.Prepare("INSERT users SET username=?")
	checkErr(err)

	_, err = stmt.Exec(user.Username)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
