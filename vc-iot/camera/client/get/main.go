package main

import (
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/aws"
	"log"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1")},
	)

	if err != nil {
		fmt.Println("session :", err)
	}

	// Create a SQS service client.
	svc := sqs.New(sess)

	for {
		receive(svc)
	}

}

func receive(svc *sqs.SQS)  {
	url := "https://sqs.ap-northeast-1.amazonaws.com/895233118661/Camera_01"

	result, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl: &url,
		MaxNumberOfMessages: aws.Int64(1),
	})
	if err != nil {
		log.Fatal("Unable to receive message from queue %v.", err)
	}

	fmt.Printf("Received %d messages.\n", len(result.Messages))
	if len(result.Messages) > 0 {
		fmt.Println(result.Messages)
	}
}
