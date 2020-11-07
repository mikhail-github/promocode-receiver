package vars

const (
	// Common constants

	// dynamoDB

	// Viper configuration
	ParamLogLevelName    = "LOG_LEVEL"
	ParamLogLevelDefault = "debug"

	ParamDynamoDBTableName  = "DYNAMODB_TABLE"
	ParamDynamoDBItemIDName = "DYNAMODB_ITEM_ID"
)

var (
	// RequiredParams - list of the required configuration parameters
	RequiredParams = []string{
		ParamDynamoDBTableName,
		ParamDynamoDBItemIDName,
	}
)
