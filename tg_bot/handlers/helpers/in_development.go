package helpers

import (
	"fmt"

	telebot "gopkg.in/telebot.v3"

	"github.com/ej-you/HamstersShaver/tg_bot/keyboards"
	"github.com/ej-you/HamstersShaver/tg_bot/services"
	stateMachine "github.com/ej-you/HamstersShaver/tg_bot/state_machine"
)


// обработка функций, находящихся в разработке
func InDevelopmentHandler(context telebot.Context) error {
	// получение машины состояний текущего юзера
	userStateMachine := stateMachine.UserStateMachines.Get(services.GetUserID(context.Chat()))
	// установка нового состояния
	if err := userStateMachine.SetStatus("in_development"); err != nil {
		return fmt.Errorf("InDevelopmentHandler: %w", err)
	}

	return context.Send("Функция на данный момент находится в разработке ⚙️", keyboards.InlineKeyboardToHome)
}
