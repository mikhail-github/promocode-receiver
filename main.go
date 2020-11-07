package main

import (
	"receiver/pkg/config"
	"receiver/pkg/handler"
	"receiver/pkg/logger"

	"github.com/aws/aws-lambda-go/lambda"
)

func init() {
	config.New()
	logger.Init(config.Params.LogLevel)
}

func main() {
	lambda.Start(handler.HandleRequest)
}
