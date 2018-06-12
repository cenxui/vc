package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"fmt"
	"github.com/aws/aws-sdk-go/service/sqs"
)

/**
todo
 */
func main() {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-1")},
	)

	if err != nil {
		fmt.Println("session :", err)
	}

	// Create a SQS service client.
	svc := sqs.New(sess)


	create(svc)



}

func create(svc *sqs.SQS) {

	queue := "Camera_01"

	result, err := svc.CreateQueue(&sqs.CreateQueueInput{
		QueueName: aws.String(queue),
		Attributes: map[string]*string{
			"DelaySeconds":           aws.String("0"),
			"MessageRetentionPeriod": aws.String("60"),
		},
	})
	if err != nil {
		fmt.Println("create :", err)
		return
	}

	fmt.Println("Success", *result.QueueUrl)

}