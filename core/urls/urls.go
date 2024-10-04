package urls

import (
	echo "github.com/labstack/echo/v4"

	accountUrls "github.com/Danil-114195722/HamstersShaver/account_app/urls"
)


// подгрузка urls каждого микроприложения и их общая настройка
func InitUrlRouters(echoApp *echo.Echo) {
	apiGroupAccount := echoApp.Group("/api/account")
	accountUrls.RouterGroup(apiGroupAccount)
}
