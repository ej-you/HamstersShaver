package handlers

import (
	"time"
	"context"
	"net/http"

	echo "github.com/labstack/echo/v4"

	myTonapiAccount "github.com/ej-you/HamstersShaver/rest_api/ton_api_rest/tonapi/account"
	coreErrors "github.com/ej-you/HamstersShaver/rest_api/core/errors"
	"github.com/ej-you/HamstersShaver/rest_api/settings"
)


// эндпоинт получения информации о всех монетах на аккаунте
//	@Summary		Get jettons balances on account
//	@Description	Get all non-null jettons balances on account and other info about jettons
//	@Router			/account/get-jettons [get]
//	@ID				get-jettons
//	@Tags			account
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Success		200	{array}	myTonapiAccount.AccountJetton
func GetJettons(ctx echo.Context) error {
	var err error
	var dataOut []myTonapiAccount.AccountJetton

	// создание API клиента TON для tonapi-go с таймаутом в 3 секунд
	tonapiClient, err := settings.GetTonClientTonapiWithTimeout("mainnet", 3*time.Second)
	if err != nil {
		return coreErrors.GetTonapiClientError
	}

	// создание контекста с таймаутом в 5 секунд
	tonApiContext, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// формирование структуры для ответа
	dataOut, err = myTonapiAccount.GetBalanceJettons(tonApiContext, tonapiClient)
	if err != nil {
		return echo.NewHTTPError(500, map[string]string{"account": err.Error()})
	}

	return ctx.JSON(http.StatusOK, dataOut)
}
