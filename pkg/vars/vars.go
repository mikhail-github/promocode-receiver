package vars

const (
	VKPosterLambdaInvokationType = "Event"
	// Common constants
	NewPromocodeMessageText = `Появился новый промокод %s
	Добавьте товары в корзину на сайте магазина, 
	затем скопируйте и вставьте в адресную строку браузера ссылку 
	%s 
	скопируйте и пришлите боту ссылку, которая получилась в результате 
	(должна открыться Ваша корзина с товарами)`

	// Viper configuration
	ParamLogLevelName    = "LOG_LEVEL"
	ParamLogLevelDefault = "debug"

	ParamDynamoDBTableName     = "DYNAMODB_TABLE"
	ParamDynamoDBPrefixName    = "DYNAMODB_PREFIX"
	ParamDynamoDBPrefixDefault = "test-"

	ParamTelegramSenderQueueURLName = "SENDER_QUEUE_URL"
	ParamVKPosterLambdaName         = "VKPOSTER_LAMBDA_NAME"

	ParamAdidasRefLinkName = "ADIDAS_REFLINK"
	ParamReebokRefLinkName = "REEBOK_REFLINK"

	ParamBannedUsersName    = "BANNED_USERS"
	ParamBannedUsersDefault = "[994595835,787082678,1179042696,1340947680,1204519029]"
)

var (
	// RequiredParams - list of the required configuration parameters
	RequiredParams = []string{
		ParamDynamoDBTableName,
		ParamTelegramSenderQueueURLName,
		ParamAdidasRefLinkName,
		ParamReebokRefLinkName,
		ParamVKPosterLambdaName,
	}
)
