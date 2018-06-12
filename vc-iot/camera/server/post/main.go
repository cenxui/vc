package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"fmt"
	"time"
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

	t := time.Now().Second();

	for {

		if time.Now().Second() - t > 3 {
			t = time.Now().Second()
			send(svc)
		}
	}



}

func send(svc *sqs.SQS)  {
	url := "https://sqs.ap-northeast-1.amazonaws.com/895233118661/Camera_01"

	result, err := svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(0),
		MessageBody: aws.String("open"),
		QueueUrl:  &url,
	})

	if err != nil {
		fmt.Println("send error", err)
		return
	}

	fmt.Println("Success", *result.MessageId)
}