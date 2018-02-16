package go_lambda_api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/aws/aws-lambda-go/events"

	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"

	"app/repository"
	"app/user"
)

func iHaveANewClient(user string) error {
	return nil
}

func iAskToCreateANewUser(username string) error {
	body, err := json.Marshal(map[string]string{"username": username})
	checkErr(err)

	Create(
		events.APIGatewayProxyRequest{
			HTTPMethod: "POST",
			Body:       string(body),
		},
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

	return user.User{Username: username}
}
