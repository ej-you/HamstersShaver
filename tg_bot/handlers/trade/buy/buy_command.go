package buy
// номер в диалоге: 0

import (
	"fmt"

	telebot "gopkg.in/telebot.v3"

	"github.com/ej-you/HamstersShaver/tg_bot/keyboards"
	"github.com/ej-you/HamstersShaver/tg_bot/services"
	stateMachine "github.com/ej-you/HamstersShaver/tg_bot/state_machine"
)


// команды: /buy
func BuyHandlerCommand(context telebot.Context) error {
	var err error

	// получение машины состояний текущего юзера
	userStateMachine := stateMachine.UserStateMachines.Get(services.GetUserID(context.Chat()))
	// установка нового состояния
	if err = userStateMachine.SetStatus("buy"); err != nil {
		return fmt.Errorf("BuyHandlerCommand: %w", err)
	}
	// установка значения действия
	if err = userStateMachine.SetAction("buy"); err != nil {
		return fmt.Errorf("BuyHandlerCommand: %w", err)
	}

	msgText := `📈 Активирован диалог покупки монет. Для отмены всех действий и выхода в главное меню используйте /cancel

Выберите DEX-биржу 👇`

	return context.Send(msgText, keyboards.InlineKeyboardChooseDEX)
}


// кнопки: to_buy
func BuyHandlerCallback(context telebot.Context) error {
	var err error

	// получение машины состоянию текущего юзера
	userStateMachine := stateMachine.UserStateMachines.Get(services.GetUserID(context.Chat()))
	// установка нового состояния
	if err = userStateMachine.SetStatus("buy"); err != nil {
		return fmt.Errorf("BuyHandlerCallback: %w", err)
	}
	// установка значения действия
	if err = userStateMachine.SetAction("buy"); err != nil {
		return fmt.Errorf("BuyHandlerCallback: %w", err)
	}

	msgText := `Хорошо. Выбрано действие покупки монет 📈

Выберите DEX-биржу 👇`

	return context.Send(msgText, keyboards.InlineKeyboardChooseDEX)
}
