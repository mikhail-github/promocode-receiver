package handler

import (
	"context"
	"encoding/json"
	"receiver/pkg/awswrapper"
	"receiver/pkg/config"

	"github.com/aws/aws-lambda-go/events"
	common "github.com/mikhail-github/promocode-common"
	log "github.com/sirupsen/logrus"
)

// HandleRequest - lambda call handler function
func HandleRequest(ctx context.Context, e events.SQSEvent) error {
	log.Debugf("Lambda Funstion started")

	var promocodes []common.Promocode
	for _, record := range e.Records {
		var p []common.Promocode
		if err := json.Unmarshal([]byte(record.Body), &p); err != nil {
			log.Errorf("Promocodes SQS string: %s unmarshal error: %s", record.Body, err.Error())
		}
		promocodes = append(promocodes, p...)
	}

	awsclient, err := awswrapper.New()
	if err != nil {
		log.Errorf("can't create aws api client: %s", err.Error())
		return err
	}

	db := common.DB{
		Client:    awsclient.DynamoDB,
		TableName: config.Params.DynamoDBTable,
		Prefix:    "test",
	}
	for _, p := range promocodes {
		if err := db.Add(&p); err != nil {
			log.Errorf("can not add promocode to database: %s", err.Error())
			return err
		}
	}

	log.Info("Lambda Function finished successfully")
	return nil
}
