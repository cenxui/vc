package main

import (
	"fmt"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"logitech.com/vc/lib/repository"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	c := new(repository.Camera)
	err:= json.Unmarshal([]byte(request.Body), c)

	if err != nil {
		panic(err)
		fmt.Println(err.Error())
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 200}, nil
	}

	repository.Put(c);

	return events.APIGatewayProxyResponse{Body: "mac :" + c.Mac + "camera :" + c.Camera  , StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
