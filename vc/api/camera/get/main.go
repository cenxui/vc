package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"logitech.com/vc/lib/repository"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	mac := request.PathParameters["mac"]
	c, err := repository.Query(mac)

	if err != nil{
		return events.APIGatewayProxyResponse{Body: "camera not found", StatusCode: 200}, nil
	}
	b, err := json.Marshal(c);

	return events.APIGatewayProxyResponse{Body: "camera :" + string(b), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
