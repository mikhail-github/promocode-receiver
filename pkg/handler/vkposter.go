package handler

import (
	"encoding/json"
	"receiver/pkg/config"
	"receiver/pkg/vars"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/service/lambda"
	common "github.com/mikhail-github/promocode-common"
	log "github.com/sirupsen/logrus"
)

func callVKPoster(l *lambda.Lambda, promocodes []common.Promocode) error {
	log.Info("Calling vkposter")
	payload, err := json.Marshal(promocodes)
	if err != nil {
		return err
	}
	input := &lambda.InvokeInput{
		FunctionName:   aws.String(config.Params.VKPosterLambdaName),
		InvocationType: aws.String(vars.VKPosterLambdaInvokationType),
		Payload:        payload,
	}
	output, err := l.Invoke(input)
	if err != nil {
		return err
	}

	log.Debugf("invoke payload: %+v", promocodes)
	log.Debugf("invoke status code: %v", *output.StatusCode)

	return nil
}
