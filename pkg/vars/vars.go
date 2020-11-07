package vars

const (
	// Common constants

	// Viper configuration
	ParamLogLevelName    = "LOG_LEVEL"
	ParamLogLevelDefault = "debug"

	ParamDynamoDBTableName     = "DYNAMODB_TABLE"
	ParamDynamoDBPrefixName    = "DYNAMODB_PREFIX"
	ParamDynamoDBPrefixDefault = "test-"
)

var (
	// RequiredParams - list of the required configuration parameters
	RequiredParams = []string{
		ParamDynamoDBTableName,
	}
)
