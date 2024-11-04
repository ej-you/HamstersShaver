package urls

import (
	echo "github.com/labstack/echo/v4"

	accountUrls "github.com/ej-you/HamstersShaver/rest_api/app_account/urls"
	jettonsUrls "github.com/ej-you/HamstersShaver/rest_api/app_jettons/urls"
	transactionsUrls "github.com/ej-you/HamstersShaver/rest_api/app_transactions/urls"
)


// подгрузка urls каждого микроприложения и их общая настройка
func InitUrlRouters(echoApp *echo.Group) {
	apiGroupAccount := echoApp.Group("/account")
	accountUrls.RouterGroup(apiGroupAccount)

	apiGroupJettons := echoApp.Group("/jettons")
	jettonsUrls.RouterGroup(apiGroupJettons)

	apiGroupTransactions := echoApp.Group("/transactions")
	transactionsUrls.RouterGroup(apiGroupTransactions)
}
