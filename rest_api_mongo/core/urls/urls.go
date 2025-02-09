package urls

import (
	echo "github.com/labstack/echo/v4"

	jettonsUrls "github.com/ej-you/HamstersShaver/rest_api_mongo/app_jettons/urls"
	transactionsUrls "github.com/ej-you/HamstersShaver/rest_api_mongo/app_transactions/urls"
)


// подгрузка urls каждого микроприложения и их общая настройка
func InitUrlRouters(echoGroup *echo.Group) {
	apiGroupJettons := echoGroup.Group("/jettons")
	jettonsUrls.RouterGroup(apiGroupJettons)

	apiGroupTransactions := echoGroup.Group("/transactions")
	transactionsUrls.RouterGroup(apiGroupTransactions)
}
