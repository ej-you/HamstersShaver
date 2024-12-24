package handlers

import (
	"fmt"

	telebot "gopkg.in/telebot.v3"

	"github.com/ej-you/HamstersShaver/tg_bot/keyboards"
	"github.com/ej-you/HamstersShaver/tg_bot/services"
	stateMachine "github.com/ej-you/HamstersShaver/tg_bot/state_machine"
)


// команды: /trade
// кнопки: to_trade
func TradeHandler(context telebot.Context) error {
	var err error
	userId := services.GetUserID(context.Chat())

	// получение машины состоянию текущего юзера
	userStateMachine := stateMachine.UserStateMachines.Get(userId)
	// установка нового состояния
	if err = userStateMachine.SetStatus("trade"); err != nil {
		return fmt.Errorf("TradeHandler for user %s: %w", userId, err)
	}

	msgText := `Активирован диалог трейдинга. Для отмены всех действий и выхода в главное меню используйте /cancel

Выберите действие 👇`

	return context.Send(msgText, keyboards.InlineKeyboardTrade)
}
