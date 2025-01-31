package urls

import (
	echo "github.com/labstack/echo/v4"

	"github.com/ej-you/HamstersShaver/rest_api_mongo/app_transactions/handlers"
)


func RouterGroup(group *echo.Group) {
	// group.GET("/buy/pre-request", handlers.BuyPreRequest)
	// group.POST("/buy/send", handlers.BuySend)

	// group.GET("/cell/pre-request", handlers.CellPreRequest)
	// group.POST("/cell/send", handlers.CellSend)

	group.POST("/new", handlers.New)
}
