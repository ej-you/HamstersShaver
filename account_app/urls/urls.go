package urls

import (
	echo "github.com/labstack/echo/v4"

	"github.com/Danil-114195722/HamstersShaver/account_app/handlers"
)


func RouterGroup(group *echo.Group) {
	group.GET("/get-ton", handlers.GetTon)
	group.GET("/get-jettons", handlers.GetJettons)
}
