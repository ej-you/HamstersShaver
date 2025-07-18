package handlers

import (
	"context"
	"fmt"

	echo "github.com/labstack/echo/v4"

	myTonapiAccount "github.com/ej-you/HamstersShaver/rest_api/ton_api/tonapi/account"
	myTongoWallet "github.com/ej-you/HamstersShaver/rest_api/ton_api/tongo/wallet"

	"github.com/ej-you/HamstersShaver/rest_api/settings/constants"
	"github.com/ej-you/HamstersShaver/rest_api/settings"
)


// seqno аккаунта
type GetSeqnoOut struct {
	Seqno uint32 `json:"seqno" example:"105" description:"порядковый номер версии кошелька аккаунта"`
}


// эндпоинт получения Seqno аккаунта
// @Title Get account seqno
// @Description Get account seqno
// @Success 200 object GetSeqnoOut "Account seqno"
// @Tag account
// @Route /account/get-seqno [get]
func GetSeqno(ctx echo.Context) error {
	// создание API клиента TON для tongo
	tongoClient, err := settings.GetTonClientTongoWithTimeout("mainnet", constants.TongoClientTimeout)
	if err != nil {
		settings.ErrorLog.Println(fmt.Errorf("get account seqno: %w", err))
		return err
	}
	// получение данных о кошельке через tongo
	realWallet, err := myTongoWallet.GetWallet(tongoClient)
	if err != nil {
		settings.ErrorLog.Println(fmt.Errorf("get account seqno: %w", err))
		return err
	}

	// создание API клиента TON для tonapi-go
	tonapiClient, err := settings.GetTonClientTonapiWithTimeout("mainnet", constants.TonapiClientTimeout)
	if err != nil {
		settings.ErrorLog.Println(fmt.Errorf("get account seqno: %w", err))
		return err
	}
	// создание контекста с таймаутом
	getAccountSeqnoContext, cancel := context.WithTimeout(context.Background(), constants.GetAccountSeqnoContextTimeout)
	defer cancel()

	// получение значения Seqno
	seqno, err := myTonapiAccount.GetAccountSeqno(getAccountSeqnoContext, tonapiClient, realWallet)
	if err != nil {
		settings.ErrorLog.Println(err)
		return err
	}

	return ctx.JSON(200, GetSeqnoOut{Seqno: seqno})
}
