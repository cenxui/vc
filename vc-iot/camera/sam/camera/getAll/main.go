package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"logitech.com/vc/lib/repository"
	"fmt"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	r, err := repository.Scan()

	if err != nil {
		fmt.Println(err.Error())
		return events.APIGatewayProxyResponse{Body: "error", StatusCode: 200}, nil
	}

	j, _ := json.Marshal(r)


	return events.APIGatewayProxyResponse{Body: string(j), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}