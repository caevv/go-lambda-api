package main

import (
	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/apigatewayproxyevt"
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
	"fmt"
	"encoding/json"

	"./infrastructure"
)

func Create(evt *apigatewayproxyevt.Event, ctx *runtime.Context) (interface{}, error) {
	var user infrastructure.User

	err := json.Unmarshal([]byte(evt.Body), &user)

	checkErr(err)

	infrastructure.Store(user)

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
