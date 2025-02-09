package cell
// номер в диалоге: 2

import (
	"fmt"

	telebot "gopkg.in/telebot.v3"

	stateMachine "github.com/ej-you/HamstersShaver/tg_bot/state_machine"
	"github.com/ej-you/HamstersShaver/tg_bot/keyboards"
	"github.com/ej-you/HamstersShaver/tg_bot/services"
)


func CellJettonsAmountHandler(context telebot.Context) error {
	var err error

	// получение машины состояний текущего юзера
	userStateMachine := stateMachine.UserStateMachines.Get(services.GetUserID(context.Chat()))
	// установка нового состояния
	if err = userStateMachine.SetStatus("cell_jettons_amount"); err != nil {
		return fmt.Errorf("CellJettonsAmountHandler: %w", err)
	}

	// получение выбранной DEX-биржи
	chosenDex := services.GetCallbackData(context.Callback())
	// установка значения DEX-биржи
	if err = userStateMachine.SetDEX(chosenDex); err != nil {
		return fmt.Errorf("CellJettonsAmountHandler: %w", err)
	}

	msgText := fmt.Sprintf(`💹 Выбранная биржа - %s

Теперь введите количество используемых монет с кошелька или их процент или выберите процент из предложенных вариантов 👇`, chosenDex)

	return context.Send(msgText, keyboards.InlineKeyboardJettonsAmountChoices)
}
