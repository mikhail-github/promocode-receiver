package config

import (
	"encoding/json"
	"log"
	"receiver/pkg/vars"

	"github.com/spf13/viper"
)

var (
	// Params variable provides access to configuration
	Params Configuration
)

// Configuration describes application configuration
type Configuration struct {
	LogLevel               string
	DynamoDBTable          string
	DynamoDBPrefix         string
	TelegramSenderQueueURL string
	VKPosterLambdaName     string
	AdidasRefLink          string
	ReebokRefLink          string
	BannedUsers            []int64
}

// New - initializes new configuration
func New() {
	initViper()

	var bannedUsers []int64
	if err := json.Unmarshal([]byte(viper.GetString(vars.ParamBannedUsersName)), &bannedUsers); err != nil {
		log.Panicf("Bot banned users list unmarshal error: %s", err.Error())
	}

	Params = Configuration{
		LogLevel:               viper.GetString(vars.ParamLogLevelName),
		DynamoDBTable:          viper.GetString(vars.ParamDynamoDBTableName),
		DynamoDBPrefix:         viper.GetString(vars.ParamDynamoDBPrefixName),
		TelegramSenderQueueURL: viper.GetString(vars.ParamTelegramSenderQueueURLName),
		VKPosterLambdaName:     viper.GetString(vars.ParamVKPosterLambdaName),
		BannedUsers:            bannedUsers,
		AdidasRefLink:          viper.GetString(vars.ParamAdidasRefLinkName),
		ReebokRefLink:          viper.GetString(vars.ParamReebokRefLinkName),
	}
}

func initViper() {
	viper.BindEnv(vars.ParamLogLevelName)
	viper.SetDefault(vars.ParamLogLevelName, vars.ParamLogLevelDefault)

	viper.BindEnv(vars.ParamDynamoDBTableName)
	viper.BindEnv(vars.ParamDynamoDBPrefixName)
	viper.SetDefault(vars.ParamDynamoDBPrefixName, vars.ParamDynamoDBPrefixDefault)
	viper.BindEnv(vars.ParamTelegramSenderQueueURLName)
	viper.BindEnv(vars.ParamVKPosterLambdaName)
	viper.BindEnv(vars.ParamAdidasRefLinkName)
	viper.BindEnv(vars.ParamReebokRefLinkName)
	viper.BindEnv(vars.ParamBannedUsersName)
	viper.SetDefault(vars.ParamBannedUsersName, vars.ParamBannedUsersDefault)

	for _, param := range vars.RequiredParams {
		if viper.GetString(param) == "" {
			log.Panicf("Parameter %s is required", param)
		}
	}
}
