package handlers

import (
	"time"
	"context"
	"net/http"

	echo "github.com/labstack/echo/v4"

	myTonapiAccount "github.com/Danil-114195722/HamstersShaver/ton_api/tonapi/account"
	myTongoWallet "github.com/Danil-114195722/HamstersShaver/ton_api/tongo/wallet"

	coreErrors "github.com/Danil-114195722/HamstersShaver/core/errors"
	"github.com/Danil-114195722/HamstersShaver/settings"
)


// эндпоинт получения Seqno аккаунта
func GetSeqno(ctx echo.Context) error {
	// создание API клиента TON для tongo с таймаутом в 3 секунд
	tongoClient, err := settings.GetTonClientTongoWithTimeout("mainnet", 3*time.Second)
	if err != nil {
		return coreErrors.GetTongoClientError
	}
	// получение данных о кошельке через tongo
	realWallet, err := myTongoWallet.GetWallet(tongoClient)
	if err != nil {
		return echo.NewHTTPError(500, map[string]string{"account": err.Error()})
	}

	// создание API клиента TON для tonapi-go с таймаутом в 3 секунд
	tonapiClient, err := settings.GetTonClientTonapiWithTimeout("mainnet", 3*time.Second)
	if err != nil {
		return coreErrors.GetTonapiClientError
	}
	// создание контекста с таймаутом в 5 секунд
	tonApiContext, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// получение значения Seqno
	seqno, err := myTonapiAccount.GetAccountSeqno(tonApiContext, tonapiClient, realWallet)
	if err != nil {
		return echo.NewHTTPError(500, map[string]string{"account": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]uint32{"seqno": seqno})
}
