package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	echo "github.com/labstack/echo/v4"

	myTonapiAccount "github.com/ej-you/HamstersShaver/rest_api/ton_api_rest/tonapi/account"
	myTongoWallet "github.com/ej-you/HamstersShaver/rest_api/ton_api_rest/tongo/wallet"

	"github.com/ej-you/HamstersShaver/rest_api/app_account/serializers"

	coreErrors "github.com/ej-you/HamstersShaver/rest_api/core/errors"
	"github.com/ej-you/HamstersShaver/rest_api/settings"
)

// эндпоинт получения Seqno аккаунта
// @Title Get account seqno
// @Description Get account seqno
// @Success 200 object serializers.GetSeqnoOut "Account seqno"
// @Tag account
// @Route /account/get-seqno [get]
func GetSeqno(ctx echo.Context) error {
	// создание API клиента TON для tongo с таймаутом в 3 секунд
	tongoClient, err := settings.GetTonClientTongoWithTimeout("mainnet", 3*time.Second)
	if err != nil {
		settings.ErrorLog.Println(fmt.Errorf("get account seqno: %w", err))
		return coreErrors.AssertAPIError(err).GetHTTPError()
	}
	// получение данных о кошельке через tongo
	realWallet, err := myTongoWallet.GetWallet(tongoClient)
	if err != nil {
		settings.ErrorLog.Println(fmt.Errorf("get account seqno: %w", err))
		return coreErrors.AssertAPIError(err).GetHTTPError()
	}

	// создание API клиента TON для tonapi-go с таймаутом в 3 секунд
	tonapiClient, err := settings.GetTonClientTonapiWithTimeout("mainnet", 3*time.Second)
	if err != nil {
		settings.ErrorLog.Println(fmt.Errorf("get account seqno: %w", err))
		return coreErrors.AssertAPIError(err).GetHTTPError()
	}
	// создание контекста с таймаутом в 5 секунд
	tonApiContext, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// получение значения Seqno
	seqno, err := myTonapiAccount.GetAccountSeqno(tonApiContext, tonapiClient, realWallet)
	if err != nil {
		settings.ErrorLog.Println(err)
		return coreErrors.AssertAPIError(err).GetHTTPError()
	}

	return ctx.JSON(http.StatusOK, serializers.GetSeqnoOut{Seqno: seqno})
}
