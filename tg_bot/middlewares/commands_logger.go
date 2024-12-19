package middlewares

import (
	telebot "gopkg.in/telebot.v3"
	
	"github.com/ej-you/HamstersShaver/tg_bot/services"
	"github.com/ej-you/HamstersShaver/tg_bot/settings"
)


// перед каждой введённой юзером командой записывает её в лог
func CommandsLogger(nextHandler telebot.HandlerFunc) telebot.HandlerFunc {
	return func(context telebot.Context) error {
		// ТГ ID юзера
		userId := services.GetUserID(context.Chat())

		callback := context.Callback()
		// если был переход по кнопке
		if callback != nil {
			settings.InfoLog.Printf("User %s use button %q", userId, callback.Unique)
		// если была введена команда
		} else {
			settings.InfoLog.Printf("User %s use command %q", userId, context.Message().Text)
		}

		return nextHandler(context)
	}
}
