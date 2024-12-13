package handlers

import (
	telebot "gopkg.in/telebot.v3"
)


// обработчик непредвиденной ошибки
func UnknownErrorHandler(err error, context telebot.Context) {
	settings.ErrorLog.Printf("BOT ERROR: %v", err)
	context.Send(
		"☠️ Возникла неизвестная ошибка. Попробуйте выйти в главное меню и попробовать ещё раз",
		// keyboards.BackToHomeInlineKeyboard,
	)
}
