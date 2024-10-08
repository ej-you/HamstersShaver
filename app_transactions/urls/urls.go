package urls

import (
	echo "github.com/labstack/echo/v4"

	"github.com/Danil-114195722/HamstersShaver/app_transactions/handlers"
)


func RouterGroup(group *echo.Group) {
	group.GET("/buy/pre-request", handlers.BuyPreRequest)
	group.POST("/buy/send", handlers.BuySend)
}
