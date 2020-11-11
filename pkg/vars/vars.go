package vars

const (
	// Common constants
	NewPromocodeMessageText = "Новый промокод %s %s"

	// Viper configuration
	ParamLogLevelName    = "LOG_LEVEL"
	ParamLogLevelDefault = "debug"

	ParamDynamoDBTableName     = "DYNAMODB_TABLE"
	ParamDynamoDBPrefixName    = "DYNAMODB_PREFIX"
	ParamDynamoDBPrefixDefault = "test-"

	ParamTelegramSenderQueueURLName = "SENDER_QUEUE_URL"
)

var (
	// RequiredParams - list of the required configuration parameters
	RequiredParams = []string{
		ParamDynamoDBTableName,
		ParamTelegramSenderQueueURLName,
	}
)
