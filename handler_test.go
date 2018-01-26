package main

import (
	"github.com/DATA-DOG/godog"
	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/apigatewayproxyevt"
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/BurntSushi/toml"
	"fmt"

	"./repository"
	"./user"
)

func iHaveANewClient(user string) error {
	return nil
}

func iAskToCreateANewUser(user string) error {
	Get(&apigatewayproxyevt.Event{
		HTTPMethod: "POST",
		QueryStringParameters:map[string]string{"username": "John"},
	},
		&runtime.Context{},
	)

	return nil
}

func theUserShouldHaveBeenCreated(username string) error {
	myUser := find(username)

	if myUser.Username != username {
		panic("username not found")
	}

	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I have a new client "([^"]*)"$`, iHaveANewClient)
	s.Step(`^I ask to create a new user "([^"]*)"$`, iAskToCreateANewUser)
	s.Step(`^the user "([^"]*)" should have been created$`, theUserShouldHaveBeenCreated)

	s.BeforeScenario(func(interface{}) {
		// TODO: create database
	})

	s.AfterScenario(func(interface{}, error) {
		// TODO: clean database
	})
}

func find(username string) user.User {
	var conf repository.Config
	_, err := toml.DecodeFile("./config.toml", &conf)
	checkErr(err)

	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.Database.User, conf.Database.Password, conf.Database.Server, conf.Database.Port, conf.Database.Database)

	db, err := sql.Open("mysql", connString)
	checkErr(err)

	// insert
	stmt, err := db.Prepare("SELECT * FROM users WHERE username=?")
	checkErr(err)

	_, err = stmt.Exec(username)
	checkErr(err)

	return user.User{Username:username}
}
