package urls

import (
	echo "github.com/labstack/echo/v4"

	accountUrls "github.com/Danil-114195722/HamstersShaver/rest_api/app_account/urls"
	jettonsUrls "github.com/Danil-114195722/HamstersShaver/rest_api/app_jettons/urls"
	transactionsUrls "github.com/Danil-114195722/HamstersShaver/rest_api/app_transactions/urls"
)


// подгрузка urls каждого микроприложения и их общая настройка
func InitUrlRouters(echoApp *echo.Echo) {
	apiGroupAccount := echoApp.Group("/api/account")
	accountUrls.RouterGroup(apiGroupAccount)

	apiGroupJettons := echoApp.Group("/api/jettons")
	jettonsUrls.RouterGroup(apiGroupJettons)

	apiGroupTransactions := echoApp.Group("/api/transactions")
	transactionsUrls.RouterGroup(apiGroupTransactions)
}
