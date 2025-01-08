package errors

import (
	"fmt"

	telebot "gopkg.in/telebot.v3"
	"github.com/pkg/errors"

	"github.com/ej-you/HamstersShaver/tg_bot/settings"
)


// создаём здесь отдельную клавиатуру для избежания кругового импорта
var inlineKeyboardToHome = func() *telebot.ReplyMarkup {
	inlineKeyboard := &telebot.ReplyMarkup{}
	var btn = inlineKeyboard.Data("главное меню", "to_home")

	inlineKeyboard.Inline(
		inlineKeyboard.Row(btn),
	)
	return inlineKeyboard
}()


// обработчик всех ошибок
func MainErrorHandler(err error, context telebot.Context) {
	userId := context.Chat().ID

	restAPIErr := new(RestAPIError)
	restAPITimeoutErr := new(RestAPITimeoutError)
	redisErr := new(RedisError)
	DBErr := new(DBError)
	validateErr := new(ValidateError)
	internalErr := new(InternalError)
	accessErr := new(AccessError)

	var msgText string
	switch {
		case errors.As(err, accessErr):
			settings.InfoLog.Printf("USER BLOCKED: %v", err)
			return
		case errors.As(err, restAPIErr):
			settings.ErrorLog.Printf("REST API ERROR (user %d): %v", userId, err)
			msgText = fmt.Sprintf("🛠 Возникла ошибка API:\n\n%v \n\nПопробуйте выйти в главное меню и попробовать ещё раз", restAPIErr)
		case errors.As(err, restAPITimeoutErr):
			settings.ErrorLog.Printf("REST API TIMEOUT ERROR (user %d): %v", userId, err)
			msgText = fmt.Sprintf("⌛️ Возникла ошибка ожидания API:\n\n%v \n\nПопробуйте выйти в главное меню и попробовать ещё раз", restAPITimeoutErr)

		case errors.As(err, redisErr):
			settings.ErrorLog.Printf("REDIS ERROR (user %d): %v", userId, err)
			msgText = fmt.Sprintf("☁️ Возникла ошибка кэша:\n\n%v \n\nПопробуйте выйти в главное меню и попробовать ещё раз", redisErr)
		case errors.As(err, DBErr):
			settings.ErrorLog.Printf("DB ERROR (user %d): %v", userId, err)
			msgText = fmt.Sprintf("🗃 Возникла ошибка БД:\n\n%v \n\nПопробуйте выйти в главное меню и попробовать ещё раз", DBErr)
		
		case errors.As(err, validateErr):
			context.Send("😬 Упсс.. Введено некорректное значение! Попробуйте ещё раз 💪")
			return

		case errors.As(err, internalErr):
			settings.ErrorLog.Printf("INTERNAL ERROR (user %d): %v", userId, err)
			msgText = fmt.Sprintf("❌ Возникла внутренняя ошибка:\n\n%v \n\nПопробуйте выйти в главное меню и попробовать ещё раз", internalErr)
		default:
			settings.ErrorLog.Printf("UNKNOWN ERROR (user %d): %v", userId, err)
			msgText = "☠️ Возникла неизвестная ошибка. Попробуйте выйти в главное меню и попробовать ещё раз"
	}
	context.Send(msgText, inlineKeyboardToHome)
}
