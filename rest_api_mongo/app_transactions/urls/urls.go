package urls

import (
	echo "github.com/labstack/echo/v4"

	"github.com/ej-you/HamstersShaver/rest_api_mongo/app_transactions/handlers"
)


func RouterGroup(group *echo.Group) {
	group.POST("/create-trade", handlers.CreateTrade)
	group.POST("/create-auto", handlers.CreateAuto)

	group.GET("/get-one", handlers.GetOne)
	group.GET("/get-many", handlers.GetMany)

	group.PATCH("", handlers.Update)
	group.PATCH("/commit-init-trans", handlers.CommitInitTrans)
}
