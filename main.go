package go_lambda_api

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"app/repository"
	"app/user"
)

func Create(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println(request.Body)

	var myUser user.User

	err := json.Unmarshal([]byte(request.Body), &myUser)
	checkErr(err)

	repository.Store(myUser)

	return events.APIGatewayProxyResponse{Body: "OK", StatusCode: 200}, nil
}

//func Get(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
//	// Get user
//	fmt.Print(evt)
//
//	return nil, nil
//}
//
//func List(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
//	// List users
//	fmt.Print(evt)
//
//	return nil, nil
//}
//
//func Update(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
//	// Update user
//	fmt.Print(evt)
//
//	return nil, nil
//}
//
//func Delete(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
//	// Delete user
//	fmt.Print(evt)
//
//	return nil, nil
//}

func main() {
	lambda.Start(Create)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
