package handlers

import (
	"fmt"

	telebot "gopkg.in/telebot.v3"

	"github.com/ej-you/HamstersShaver/tg_bot/keyboards"
	"github.com/ej-you/HamstersShaver/tg_bot/services"
	stateMachine "github.com/ej-you/HamstersShaver/tg_bot/state_machine"
)


// обработка функций, находящихся в разработке
func InDevelopmentHandler(context telebot.Context) error {
	userId := services.GetUserID(context.Chat())
	// получение машины состоянию текущего юзера
	userStateMachine := stateMachine.UserStateMachines.Get(userId)
	// установка нового состояния
	if err := userStateMachine.SetStatus("in_development"); err != nil {
		return fmt.Errorf("InDevelopmentHandler for user %s: %w", userId, err)
	}

	msgText := `Функция на данный момент находится в разработке ⚙️`
	return context.Send(msgText, keyboards.InlineKeyboardToHome)
}
