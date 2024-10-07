package urls

import (
	echo "github.com/labstack/echo/v4"

	"github.com/Danil-114195722/HamstersShaver/app_jettons/handlers"
)


func RouterGroup(group *echo.Group) {
	group.GET("/get-info", handlers.GetInfo)
}
