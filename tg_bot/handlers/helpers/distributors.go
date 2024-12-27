package helpers

import (
	"fmt"

	telebot "gopkg.in/telebot.v3"

	handlersTradeCell "github.com/ej-you/HamstersShaver/tg_bot/handlers/trade/cell"
	handlersTradeBuy "github.com/ej-you/HamstersShaver/tg_bot/handlers/trade/buy"

	stateMachine "github.com/ej-you/HamstersShaver/tg_bot/state_machine"
	"github.com/ej-you/HamstersShaver/tg_bot/services"
)


// словарь статусов с их обработчиками для произвольных кнопок
var statusesToCheck = map[string]func(context telebot.Context)error{
	// в самом начале: выбор монеты
	"cell": handlersTradeCell.CellChooseDEXHandler, // выбор биржи
	"cell_dex": handlersTradeCell.CellJettonsAmountHandler, // кол-во монет
	"cell_jettons_amount": handlersTradeCell.CellSlippageHandler, // процент проскальзывания
	"cell_slippage": handlersTradeCell.CellConfirmTransactionHandler, // подтверждение транзакции

	// в самом начале: выбор биржи
	"buy": handlersTradeBuy.BuyTonsAmountHandler, // кол-во монет
	"buy_tons_amount": handlersTradeBuy.BuySlippageHandler, // процент проскальзывания
	"buy_slippage": handlersTradeBuy.BuyJettonCAHandler, // адрес монеты
	"buy_jetton_ca": handlersTradeBuy.BuyConfirmTransactionHandler, // подтверждение транзакции
}


// распределитель хэндлеров по статусу для ответов на произвольные кнопки и произвольный текст
func HandlersDistributor(context telebot.Context) error {
	var match bool
	var err error

	// получение машины состояний текущего юзера
	userStateMachine := stateMachine.UserStateMachines.Get(services.GetUserID(context.Chat()))

	// перебор в цикле всех статусов, для которых могут быть использованы произвольные кнопки
	for statusToCheck, handler := range statusesToCheck {
		match, err = userStateMachine.StatusEquals(statusToCheck)
		if err != nil {
			return fmt.Errorf("CallbackDistributor: %w", err)
		}
		if match {
			return handler(context)
		}
	}
	return nil
}
