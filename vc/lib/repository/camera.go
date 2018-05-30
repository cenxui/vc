package repository

import (
	"errors"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Camera struct {
	MAC string `json:"mac"`
	Camera string `json:"camera"`
}

var svc *dynamodb.DynamoDB

func init() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1")},
	)

	if err != nil {
		println(err.Error())
	}

	// Create DynamoDB client
	svc = dynamodb.New(sess)
}

func Put(item *Camera) error {
	av, err := dynamodbattribute.MarshalMap(item)

	input := &dynamodb.PutItemInput{
		Item: av,
		TableName: aws.String("camera"),
	}
	_, err = svc.PutItem(input)

	return err
}

func Query(mac string) (*Camera, error) {
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("camera"),
		Key: map[string]*dynamodb.AttributeValue{
			"mac": {
				S: aws.String(mac),
			},
		},
	})

	if err != nil {
		return nil, err
	}

	item := Camera{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)

	if err != nil {
		return nil, err
	}

	if item.MAC == "" {
		return nil, errors.New("Item not found")
	}
	return &item, nil
}