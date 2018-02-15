package go_lambda_api

import (
	"encoding/json"
	"fmt"
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/apigatewayproxyevt"

	"app/repository"
	"app/user"
)

func Create(evt *apigatewayproxyevt.Event, ctx *runtime.Context) (interface{}, error) {
	var myUser user.User

	err := json.Unmarshal([]byte(evt.Body), &myUser)

	checkErr(err)

	repository.Store(myUser)

	return myUser, nil
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
