package buy
// номер в диалоге: 1

import (
	"fmt"

	telebot "gopkg.in/telebot.v3"

	stateMachine "github.com/ej-you/HamstersShaver/tg_bot/state_machine"
	"github.com/ej-you/HamstersShaver/tg_bot/services"
)


func BuyTonsAmountHandler(context telebot.Context) error {
	var err error

	// получение машины состояний текущего юзера
	userStateMachine := stateMachine.UserStateMachines.Get(services.GetUserID(context.Chat()))
	// установка нового состояния
	if err = userStateMachine.SetStatus("buy_tons_amount"); err != nil {
		return fmt.Errorf("BuyTonsAmountHandler: %w", err)
	}

	// получение выбранной DEX-биржи
	chosenDex := services.GetCallbackData(context.Callback())
	// установка значения DEX-биржи
	if err = userStateMachine.SetDEX(chosenDex); err != nil {
		return fmt.Errorf("BuyTonsAmountHandler: %w", err)
	}

	msgText := fmt.Sprintf(`💹 Выбранная биржа - %s

Теперь введите количество используемых TON с кошелька или их процент`, chosenDex)

	return context.Send(msgText)
}
