package config

import (
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
	LogLevel       string
	DynamoDBTable  string
	DynamoDBItemID string
}

// New - initializes new configuration
func New() {
	initViper()

	Params = Configuration{
		LogLevel:       viper.GetString(vars.ParamLogLevelName),
		DynamoDBTable:  viper.GetString(vars.ParamDynamoDBTableName),
		DynamoDBItemID: viper.GetString(vars.ParamDynamoDBItemIDName),
	}
}

func initViper() {
	viper.BindEnv(vars.ParamLogLevelName)
	viper.SetDefault(vars.ParamLogLevelName, vars.ParamLogLevelDefault)

	viper.BindEnv(vars.ParamDynamoDBTableName)
	viper.BindEnv(vars.ParamDynamoDBItemIDName)

	for _, param := range vars.RequiredParams {
		if viper.GetString(param) == "" {
			log.Panicf("Parameter %s is required", param)
		}
	}
}
