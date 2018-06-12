package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"log"
	"github.com/aws/aws-sdk-go/service/sqs"
	"fmt"
)

func main() {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1")},
	)

	if err != nil {
		log.Fatal("error")
	}

	svc := sqs.New(sess)
	queue := "CAMERA_01"
	//create(svc, queue)
	send(svc, queue)
	receive(svc, queue)
}

func create(svc *sqs.SQS, queue string)  {

	result, err := svc.CreateQueue(&sqs.CreateQueueInput{
		QueueName: aws.String(queue),
	})
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Success", *result.QueueUrl)
}

func receive(svc *sqs.SQS, queue string)  {
	// Receive a message from the SQS queue with long polling enabled.

	resultURL, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(queue),
	})

	result, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl: resultURL.QueueUrl,
		AttributeNames: aws.StringSlice([]string{
			"SentTimestamp",
		}),
		MaxNumberOfMessages: aws.Int64(2),
		MessageAttributeNames: aws.StringSlice([]string{
			"All",
		}),
		WaitTimeSeconds: aws.Int64(20),
	})
	if err != nil {
		log.Fatal("Unable to receive message from queue %q, %v.", queue, err)
	}

	fmt.Printf("Received %d messages.\n", len(result.Messages))
	if len(result.Messages) > 0 {
		fmt.Println(result.Messages)
	}


}

func send(svc *sqs.SQS, queue string)  {
	resultURL, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(queue),
	})


	result, err := svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageBody: aws.String("open"),
		QueueUrl:  resultURL.QueueUrl,
	})

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Success", *result.MessageId)
}