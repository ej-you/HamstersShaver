package urls

import (
	echo "github.com/labstack/echo/v4"

	"github.com/ej-you/HamstersShaver/rest_api/app_services/handlers"
)


func RouterGroup(group *echo.Group) {
	group.GET("/beauty-balance", handlers.BeautyBalance)
}
