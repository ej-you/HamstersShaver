package urls

import (
	echo "github.com/labstack/echo/v4"

	"github.com/ej-you/HamstersShaver/rest_api/app_account/handlers"
)


func RouterGroup(group *echo.Group) {
	group.GET("/get-ton", handlers.GetTon)
	group.GET("/get-jettons", handlers.GetJettons)
	group.GET("/get-jetton", handlers.GetJetton)
	group.GET("/get-seqno", handlers.GetSeqno)
}
