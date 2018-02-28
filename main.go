package go_lambda_api

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"app/repository"
	"app/user"
)

func Create(request events.APIGatewayProxyRequest) Response {
	var myUser user.User

	err := json.Unmarshal([]byte(request.Body), &myUser)
	checkErr(err)

	repository.Store(myUser)

	return Response{Code:200}
}

func Get(request events.APIGatewayProxyRequest) Response {
	type Username struct {
		Username string
	}

	var username Username

	err := json.Unmarshal([]byte(request.Body), &username)
	checkErr(err)

	myUser := repository.Find(username.Username)
	jsonUser, err := json.Marshal(myUser)
	checkErr(err)

	return Response{jsonUser, 200}
}

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse) {
	var response Response

	switch request.HTTPMethod {
	case "POST":
		response = Create(request)
	case "GET":
		response = Get(request)
	}

	return events.APIGatewayProxyResponse{Body: string(response.Body), StatusCode: response.Code}
}

func main() {
	lambda.Start(HandleRequest)
}

type Response struct {
	Body []byte
	Code int
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
