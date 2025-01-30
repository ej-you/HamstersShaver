package handlers

import (
	"context"

	echo "github.com/labstack/echo/v4"

	myTongoTransactions "github.com/ej-you/HamstersShaver/rest_api/ton_api/tongo/transactions"
	
	coreValidator "github.com/ej-you/HamstersShaver/rest_api/core/validator"
	"github.com/ej-you/HamstersShaver/rest_api/settings/constants"
	"github.com/ej-you/HamstersShaver/rest_api/settings"
)


// структура входных данных для отправки транзакции на продажу
type CellSendIn struct {
	JettonCA string `json:"jettonCA" validate:"required" example:"EQC47093oX5Xhb0xuk2lCr2RhS8rj-vul61u4W2UH5ORmG_O" description:"мастер-адрес продаваемой монеты (jetton_master)"`
	Amount float64 `json:"amount" validate:"required" example:"200" description:"кол-во используемых монет на продажу в формате, удобном для человека"`
	Slippage int `json:"slippage" validate:"required,min=0,max=100" example:"20" description:"процент проскальзывания"`
}

// успешная отправка транзакции на продажу
type CellSendOut struct {
	Success bool `json:"success" example:"true" description:"успех"`
}


// эндпоинт отправки транзакции на продажу
// @Title Cell send
// @Description Send transaction to cell jettons to TON
// @Param CellSendIn body CellSendIn true "Cтруктура входных данных для отправки транзакции на продажу"
// @Success 202 object CellSendOut "Transaction was sent successfully"
// @Tag transactions
// @Route /transactions/cell/send [post]
func CellSend(ctx echo.Context) error {
	var err error
	var dataIn CellSendIn

	// парсинг JSON-body
	if err = ctx.Bind(&dataIn); err != nil {
		return err
	}
	// валидация полученной структуры
	if err = coreValidator.GetValidator().Struct(&dataIn); err != nil {
		return err
	}

	// создание контекста с таймаутом
	sendCellJettonContext, cancel := context.WithTimeout(context.Background(), constants.SendCellJettonContextTimeout)
	defer cancel()

	// отправка транзакции на продажу
	err = myTongoTransactions.CellJetton(sendCellJettonContext, dataIn.JettonCA, dataIn.Amount, dataIn.Slippage)
	if err != nil {
		settings.ErrorLog.Println(err)
		return err
	}

	return ctx.JSON(202, CellSendOut{Success: true})
}
