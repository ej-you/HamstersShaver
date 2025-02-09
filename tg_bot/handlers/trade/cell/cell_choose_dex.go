package cell
// номер в диалоге: 1

import (
	"fmt"

	telebot "gopkg.in/telebot.v3"

	stateMachine "github.com/ej-you/HamstersShaver/tg_bot/state_machine"
	"github.com/ej-you/HamstersShaver/tg_bot/keyboards"
	"github.com/ej-you/HamstersShaver/tg_bot/services"
)


// кнопки выбора монеты для продажи в диалоге команды /cell
func CellChooseDEXHandler(context telebot.Context) error {
	var err error

	// получение машины состояний текущего юзера
	userStateMachine := stateMachine.UserStateMachines.Get(services.GetUserID(context.Chat()))
	// установка нового состояния
	if err = userStateMachine.SetStatus("cell_dex"); err != nil {
		return fmt.Errorf("CellChooseDEXHandler: %w", err)
	}
	// установка значения CA монеты
	if err = userStateMachine.SetJettonCA(context.Callback().Data); err != nil {
		return fmt.Errorf("CellHandlerCommand: %w", err)
	}

	return context.Send("🪙 Отлично! Монета выбрана. Выберите DEX-биржу 👇", keyboards.InlineKeyboardChooseDEX)
}
