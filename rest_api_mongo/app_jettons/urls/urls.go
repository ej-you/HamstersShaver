package urls

import (
	echo "github.com/labstack/echo/v4"

	"github.com/ej-you/HamstersShaver/rest_api_mongo/app_jettons/handlers"
)


func RouterGroup(group *echo.Group) {
	group.GET("/get-many", handlers.GetMany)
	group.GET("/get-one", handlers.GetOne)

	group.POST("", handlers.Create)
	group.DELETE("", handlers.Delete)
}
