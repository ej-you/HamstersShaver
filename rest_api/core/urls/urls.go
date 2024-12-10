package urls

import (
	echo "github.com/labstack/echo/v4"

	accountUrls "github.com/ej-you/HamstersShaver/rest_api/app_account/urls"
	jettonsUrls "github.com/ej-you/HamstersShaver/rest_api/app_jettons/urls"
	transactionsUrls "github.com/ej-you/HamstersShaver/rest_api/app_transactions/urls"
	servicesUrls "github.com/ej-you/HamstersShaver/rest_api/app_services/urls"
)


// подгрузка urls каждого микроприложения и их общая настройка
func InitUrlRouters(echoGroup *echo.Group) {
	apiGroupAccount := echoGroup.Group("/account")
	accountUrls.RouterGroup(apiGroupAccount)

	apiGroupJettons := echoGroup.Group("/jettons")
	jettonsUrls.RouterGroup(apiGroupJettons)

	apiGroupTransactions := echoGroup.Group("/transactions")
	transactionsUrls.RouterGroup(apiGroupTransactions)

	apiGroupServices := echoGroup.Group("/services")
	servicesUrls.RouterGroup(apiGroupServices)
}
