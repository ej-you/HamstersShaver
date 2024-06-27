package ton_api

import (
	"github.com/xssnick/tonutils-go/ton"

	"github.com/Danil-114195722/HamstersShaver/settings"
)


var API ton.APIClientWrapped = settings.GetTonClient()
