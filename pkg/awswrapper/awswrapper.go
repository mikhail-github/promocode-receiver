package awswrapper

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/lambda"
)

// Client provides AWS API client
type Client struct {
	DynamoDB *dynamodb.DynamoDB
	Lambda   *lambda.Lambda
}

// New provides AWS API client
func New() (*Client, error) {
	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}

	return &Client{
		DynamoDB: dynamodb.New(sess),
		Lambda:   lambda.New(sess),
	}, nil
}
