package awssqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	log "github.com/sirupsen/logrus"
)

// Client type provides aws SQS client
type Client struct {
	SQS *sqs.SQS
}

// New provides aws api client
func New() (*Client, error) {
	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}

	log.Debug("AWS SQS client created successfully")
	return &Client{
		SQS: sqs.New(sess),
	}, nil
}

// Send provides message delivery to the queue
func (c *Client) SQSSend(message, queueURL string) error {
	log.Debugf("Sending SQS message to %s", queueURL)
	res, err := c.SQS.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(message),
		QueueUrl:    &queueURL,
	})
	if err != nil {
		return err
	}

	log.Debugf("Message sent successfully, MessageId: %s", *res.MessageId)
	return nil
}
