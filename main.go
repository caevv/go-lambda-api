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

func main() {
	lambda.Start(Create)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
