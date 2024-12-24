package handlers

import (
	"fmt"

	telebot "gopkg.in/telebot.v3"

	"github.com/ej-you/HamstersShaver/tg_bot/keyboards"
	"github.com/ej-you/HamstersShaver/tg_bot/services"
	stateMachine "github.com/ej-you/HamstersShaver/tg_bot/state_machine"
)


// команды: /cell
func CellHandlerCommand(context telebot.Context) error {
	var err error
	userId := services.GetUserID(context.Chat())

	// получение машины состоянию текущего юзера
	userStateMachine := stateMachine.UserStateMachines.Get(userId)
	// установка нового состояния
	if err = userStateMachine.SetStatus("cell"); err != nil {
		return fmt.Errorf("CellHandlerCommand for user %s: %w", userId, err)
	}
	// установка значения действия
	if err = userStateMachine.SetAction("cell"); err != nil {
		return fmt.Errorf("CellHandlerCommand for user %s: %w", userId, err)
	}

	msgText := `📉 Активирован диалог продажи монет. Для отмены всех действий и выхода в главное меню используйте /cancel

Выберите из имеющихся у вас на аккаунте монет ту, которую хотите продать 👇`

	// обновление кнопок клавиатуры в соответствии с текущим списком монет на кошельке аккаунта
	err = keyboards.SetWalletJettonsButtons()
	if err != nil {
		return fmt.Errorf("CellHandlerCommand for user %s: %w", userId, err)
	}
	return context.Send(msgText, keyboards.InlineKeyboardWalletJettons)
}


// кнопки: to_cell
func CellHandlerCallback(context telebot.Context) error {
	var err error
	userId := services.GetUserID(context.Chat())

	// получение машины состоянию текущего юзера
	userStateMachine := stateMachine.UserStateMachines.Get(userId)
	// установка нового состояния
	if err = userStateMachine.SetStatus("cell"); err != nil {
		return fmt.Errorf("CellHandlerCommand for user %s: %w", userId, err)
	}
	// установка значения действия
	if err = userStateMachine.SetAction("cell"); err != nil {
		return fmt.Errorf("CellHandlerCommand for user %s: %w", userId, err)
	}

	msgText := `Отлично! Выбрано действие продажи монет 📉

Выберите из имеющихся у вас на аккаунте монет ту, которую хотите продать 👇`

	// обновление кнопок клавиатуры в соответствии с текущим списком монет на кошельке аккаунта
	err = keyboards.SetWalletJettonsButtons()
	if err != nil {
		return fmt.Errorf("CellHandlerCommand for user %s: %w", userId, err)
	}
	return context.Send(msgText, keyboards.InlineKeyboardWalletJettons)
}
