package handlers

import (
	telebot "gopkg.in/telebot.v3"

	"github.com/ej-you/HamstersShaver/tg_bot/keyboards"
)


// обработка функций, находящихся в разработке
func InDevelopmentHandler(context telebot.Context) error {
	msgText := `Функция на данный момент находится в разработке ⚙️`
	return context.Send(msgText, keyboards.InlineKeyboardToHome)
}
