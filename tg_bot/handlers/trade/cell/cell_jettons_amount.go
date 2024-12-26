package cell
// номер в диалоге: 2

import (
	"fmt"

	telebot "gopkg.in/telebot.v3"

	stateMachine "github.com/ej-you/HamstersShaver/tg_bot/state_machine"
	"github.com/ej-you/HamstersShaver/tg_bot/services"
)


func CellJettonsAmountHandler(context telebot.Context) error {
	var err error
	userId := services.GetUserID(context.Chat())

	// получение машины состояний текущего юзера
	userStateMachine := stateMachine.UserStateMachines.Get(userId)
	// установка нового состояния
	if err = userStateMachine.SetStatus("cell_jettons_amount"); err != nil {
		return fmt.Errorf("CellJettonsAmountHandler for user %s: %w", userId, err)
	}

	// получение выбранной DEX-биржи
	chosenDex := services.GetCallbackData(context.Callback())
	// установка значения DEX-биржи
	if err = userStateMachine.SetDEX(chosenDex); err != nil {
		return fmt.Errorf("CellJettonsAmountHandler for user %s: %w", userId, err)
	}

	// корректировка для вывода
	if chosenDex == "stonfi" {
		chosenDex = "Ston.fi"
	} else if chosenDex == "dedust" {
		chosenDex = "Dedust.io"
	}

	msgText := fmt.Sprintf(`💹 Выбранная биржа - %s

Теперь введите количество используемых токенов с кошелька (число больше 0) или их процент (число от 1 до 100 с %% на конце)`, chosenDex)

	return context.Send(msgText)
}
