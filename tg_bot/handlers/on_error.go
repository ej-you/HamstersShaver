package handlers

import (
	telebot "gopkg.in/telebot.v3"

	"github.com/ej-you/HamstersShaver/tg_bot/keyboards"
	"github.com/ej-you/HamstersShaver/tg_bot/settings"
)


// обработчик непредвиденной ошибки
func UnknownErrorHandler(err error, context telebot.Context) {
	settings.ErrorLog.Printf("BOT CRITICAL ERROR: %v", err)
	context.Send(
		"☠️ Возникла неизвестная ошибка. Попробуйте выйти в главное меню и попробовать ещё раз",
		keyboards.InlineKeyboardToHome,
	)
}
