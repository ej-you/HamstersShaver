package middlewares

import (
	telebot "gopkg.in/telebot.v3"
	
	"github.com/ej-you/HamstersShaver/tg_bot/services"
	"github.com/ej-you/HamstersShaver/tg_bot/settings"
)


// перед каждым введённой юзером командой записываем её в лог
func GeneralCommandsLogger(nextHandler telebot.HandlerFunc) telebot.HandlerFunc {
	return func(context telebot.Context) error {
		telegramUserId := services.GetUserID(context.Chat())
		settings.InfoLog.Printf("User %s use command %q", telegramUserId, context.Message().Text)

		return nextHandler(context)
	}
}

// перед каждой нажатой юзером инлайн-кнопкой записываем её в лог
func GeneralCallbackLogger(nextHandler telebot.HandlerFunc) telebot.HandlerFunc {
	return func(context telebot.Context) error {
		telegramUserId := services.GetUserID(context.Chat())
		settings.InfoLog.Printf("User %s use button %q", telegramUserId, context.Callback().Unique)

		return nextHandler(context)
	}
}
