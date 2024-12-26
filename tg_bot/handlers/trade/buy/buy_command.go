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
	userId := services.GetUserID(context.Chat())

	// получение машины состояний текущего юзера
	userStateMachine := stateMachine.UserStateMachines.Get(userId)
	// установка нового состояния
	if err = userStateMachine.SetStatus("buy"); err != nil {
		return fmt.Errorf("BuyHandlerCommand for user %s: %w", userId, err)
	}
	// установка значения действия
	if err = userStateMachine.SetAction("buy"); err != nil {
		return fmt.Errorf("BuyHandlerCommand for user %s: %w", userId, err)
	}

	msgText := `📈 Активирован диалог покупки монет. Для отмены всех действий и выхода в главное меню используйте /cancel

Выберите DEX-биржу 👇`

	return context.Send(msgText, keyboards.InlineKeyboardChooseDEX)
}


// кнопки: to_buy
func BuyHandlerCallback(context telebot.Context) error {
	var err error
	userId := services.GetUserID(context.Chat())

	// получение машины состоянию текущего юзера
	userStateMachine := stateMachine.UserStateMachines.Get(userId)
	// установка нового состояния
	if err = userStateMachine.SetStatus("buy"); err != nil {
		return fmt.Errorf("BuyHandlerCallback for user %s: %w", userId, err)
	}
	// установка значения действия
	if err = userStateMachine.SetAction("buy"); err != nil {
		return fmt.Errorf("BuyHandlerCallback for user %s: %w", userId, err)
	}

	msgText := `Хорошо. Выбрано действие покупки монет 📈

Выберите DEX-биржу 👇`

	return context.Send(msgText, keyboards.InlineKeyboardChooseDEX)
}
