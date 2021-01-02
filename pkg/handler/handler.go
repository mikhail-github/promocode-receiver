package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"receiver/pkg/awssqs"
	"receiver/pkg/awswrapper"
	"receiver/pkg/config"
	"receiver/pkg/vars"

	"github.com/aws/aws-lambda-go/events"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	common "github.com/mikhail-github/promocode-common"
	subscribers "github.com/mikhail-github/promocode-subscribers"
	senderCommon "github.com/mikhail-github/telegram-sender-common"
	log "github.com/sirupsen/logrus"
)

var (
	SubscribersDB subscribers.DB
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
		Prefix:    config.Params.DynamoDBPrefix,
	}
	SubscribersDB = subscribers.DB{
		Client:    awsclient.DynamoDB,
		TableName: config.Params.DynamoDBTable,
		Prefix:    config.Params.DynamoDBPrefix,
	}
	for _, p := range promocodes {
		if err := db.Add(&p); err != nil {
			log.Errorf("can not add promocode to database: %s", err.Error())
			return err
		}
		if err := sendUpdateToSubscribers(&p); err != nil {
			log.Errorf("can not send promocode to subscribers: %s", err.Error())
			return err
		}
	}

	log.Info("Lambda Function finished successfully")
	return nil
}

func sendUpdateToSubscribers(promo *common.Promocode) error {
	subscribers, err := SubscribersDB.GetAll()
	if err != nil {
		return err
	}

	var message senderCommon.Message
	msg := tgbotapi.NewMessage(0, promoUpdateMessage(promo.ShopID))
	message.Msg.MessageConfig = &msg

	for _, s := range subscribers {
		if s.NewPromoSubscription && !isIntInSlice(config.Params.BannedUsers, s.ID) {
			message.Recipients = append(message.Recipients, s.ID)
		}
	}

	log.Debugf("message: %+v", message)

	sqs, err := awssqs.New()
	if err != nil {
		log.Panicf("SQS client creation error: %s", err.Error())
	}
	messages := []senderCommon.Message{message}
	j, err := json.Marshal(messages)
	if err != nil {
		log.Panicf("can not marshal message: %+v error: %s", messages, err.Error())
	}

	if err := sqs.SQSSend(string(j), config.Params.TelegramSenderQueueURL); err != nil {
		log.Panicf("bot.Send() can not send message to telegram sender sqs queue: %s", err.Error())
	} else {
		log.Debugf("Message send to %s successfully", config.Params.TelegramSenderQueueURL)
	}

	return nil
}

func promoUpdateMessage(shopID common.PromocodeShopID) string {
	var link string
	switch shopID {
	case common.AdidasShopID:
		link = config.Params.AdidasRefLink
	case common.ReebokShopID:
		link = config.Params.ReebokRefLink
	default:
		log.Panicf("Unknown shopID: %s", shopID)
	}

	return fmt.Sprintf(vars.NewPromocodeMessageText, shopID, link)
}

func isIntInSlice(s []int64, i int64) bool {
	for _, v := range s {
		if i == v {
			return true
		}
	}

	return false
}
