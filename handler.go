package main

import (
	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/apigatewayproxyevt"
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
	"github.com/BurntSushi/toml"
	"fmt"
	"encoding/json"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
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

func store(user User) {
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

type User struct {
	Username string
}

/// Create
func Create(evt *apigatewayproxyevt.Event, ctx *runtime.Context) (interface{}, error) {
	// Create user
	var user User

	err := json.Unmarshal([]byte(evt.Body), &user)

	checkErr(err)

	store(user)

	return user, nil
}

func Get(evt *apigatewayproxyevt.Event, _ *runtime.Context) (interface{}, error) {
	// Get user
	fmt.Print(evt)

	return nil, nil
}

func List(evt *apigatewayproxyevt.Event, _ *runtime.Context) (interface{}, error) {
	// List users
	fmt.Print(evt)

	return nil, nil
}

func Update(evt *apigatewayproxyevt.Event, _ *runtime.Context) (interface{}, error) {
	// Update user
	fmt.Print(evt)

	return nil, nil
}

func Delete(evt *apigatewayproxyevt.Event, _ *runtime.Context) (interface{}, error) {
	// Delete user
	fmt.Print(evt)

	return nil, nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
