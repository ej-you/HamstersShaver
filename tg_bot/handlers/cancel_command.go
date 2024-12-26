package handlers

import (
	telebot "gopkg.in/telebot.v3"
)


// команда: /cancel
func CancelHandler(context telebot.Context) error {
	context.Send("❌ Отмена всех действий. Возврат в главное меню...")
	return HomeHandler(context)
}
