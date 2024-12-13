package handlers

import (
	telebot "gopkg.in/telebot.v3"
)


// команда /start
func StartHandler(context telebot.Context) error {
	// newStatus := "home"
	// // установка состояния юзера
	// err := redis.SetStatus(redisClient, services.GetUserID(context), newStatus)
	// if err != nil {
	// 	return context.Send(errorMessage, keyboards.BackToHomeInlineKeyboard)
	// }

	msgText := `Привет 👋

❗️Для получения полной инструкции введите /help`

	return context.Send(msgText)
}
