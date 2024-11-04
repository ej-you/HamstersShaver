package urls

import (
	echo "github.com/labstack/echo/v4"

	"github.com/ej-you/HamstersShaver/rest_api/app_jettons/handlers"
)


func RouterGroup(group *echo.Group) {
	group.GET("/get-info", handlers.GetInfo)
}
