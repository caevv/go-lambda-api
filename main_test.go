package go_lambda_api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/aws/aws-lambda-go/events"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"app/repository"
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
	myUser := repository.Find(username)

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
		os.Setenv("HOST", "database")
		os.Setenv("PORT", "3306")
		os.Setenv("DATABASE", "go")
		os.Setenv("USER", "user")
		os.Setenv("PASSWORD", "pass")

		createTable()
	})

	s.AfterScenario(func(interface{}, error) {
		removeTable()
	})
}

func createTable() {
	db := getDb()

	// insert
	_, err := db.Exec("CREATE table users (username VARCHAR(100))")
	checkErr(err)
}

func removeTable() {
	db := getDb()

	// insert
	_, err := db.Exec("DROP TABLE users;")
	checkErr(err)
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
