package handlers

import (
	"context"

	echo "github.com/labstack/echo/v4"

	myTongoTransactions "github.com/ej-you/HamstersShaver/rest_api/ton_api/tongo/transactions"

	coreValidator "github.com/ej-you/HamstersShaver/rest_api/core/validator"
	"github.com/ej-you/HamstersShaver/rest_api/settings/constants"
)


// структура входных данных для отправки транзакции на покупку
type BuySendIn struct {
	JettonCA string `json:"jettonCA" validate:"required" example:"EQC47093oX5Xhb0xuk2lCr2RhS8rj-vul61u4W2UH5ORmG_O" description:"мастер-адрес покупаемой монеты (jetton_master)"`
	Amount float64 `json:"amount" validate:"required" example:"0.1" description:"кол-во используемых TON для покупки в формате, удобном для человека"` 
	Slippage int `json:"slippage" validate:"required,min=1,max=100" example:"20" description:"процент проскальзывания"`
}

// успешная отправка транзакции на покупку
type BuySendOut struct {
	Success bool `json:"success" example:"true", description:"успех"`
}


// эндпоинт отправки транзакции на покупку
// @Title Buy send
// @Description Send transaction to buy jettons using TON
// @Param BuySendIn body BuySendIn true "Cтруктура входных данных для отправки транзакции на покупку"
// @Success 202 object BuySendOut "Transaction was sent successfully"
// @Tag transactions
// @Route /transactions/buy/send [post]
func BuySend(ctx echo.Context) error {
	var err error
	var dataIn BuySendIn

	// парсинг JSON-body
	if err = ctx.Bind(&dataIn); err != nil {
		return err
	}
	// валидация полученной структуры
	if err = coreValidator.GetValidator().Struct(&dataIn); err != nil {
		return err
	}

	// создание контекста с таймаутом
	sendBuyJettonContext, cancel := context.WithTimeout(context.Background(), constants.SendBuyJettonContextTimeout)
	defer cancel()

	// отправка транзакции на покупку
	err = myTongoTransactions.BuyJetton(sendBuyJettonContext, dataIn.JettonCA, dataIn.Amount, dataIn.Slippage)
	if err != nil {
		return err
	}

	return ctx.JSON(202, BuySendOut{Success: true})
}
