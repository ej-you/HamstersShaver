package helpers

import (
	"fmt"

	telebot "gopkg.in/telebot.v3"
	"github.com/pkg/errors"

	customErrors "github.com/ej-you/HamstersShaver/tg_bot/errors"
	"github.com/ej-you/HamstersShaver/tg_bot/keyboards"
	"github.com/ej-you/HamstersShaver/tg_bot/services"
	"github.com/ej-you/HamstersShaver/tg_bot/settings"
)


// обработчик всех ошибок
func MainErrorHandler(err error, context telebot.Context) {
	userId := services.GetUserID(context.Chat())

	restAPIErr := new(customErrors.RestAPIError)
	restAPITimeoutErr := new(customErrors.RestAPITimeoutError)
	redisErr := new(customErrors.RedisError)
	DBErr := new(customErrors.DBError)
	validateErr := new(customErrors.ValidateError)
	internalErr := new(customErrors.InternalError)
	accessErr := new(customErrors.AccessError)

	var msgText string
	switch {
		case errors.As(err, accessErr):
			settings.InfoLog.Printf("USER BLOCKED: %v", err)
			return
		case errors.As(err, restAPIErr):
			settings.ErrorLog.Printf("REST API ERROR (user %s): %v", userId, err)
			msgText = fmt.Sprintf("🛠 Возникла ошибка API:\n\n%v \n\nПопробуйте выйти в главное меню и попробовать ещё раз", restAPIErr)
		case errors.As(err, restAPITimeoutErr):
			settings.ErrorLog.Printf("REST API TIMEOUT ERROR (user %s): %v", userId, err)
			msgText = fmt.Sprintf("⌛️ Возникла ошибка ожидания API:\n\n%v \n\nПопробуйте выйти в главное меню и попробовать ещё раз", restAPITimeoutErr)
		
		case errors.As(err, redisErr):
			settings.ErrorLog.Printf("REDIS ERROR (user %s): %v", userId, err)
			msgText = fmt.Sprintf("💸 Возникла ошибка кэша:\n\n%v \n\nПопробуйте выйти в главное меню и попробовать ещё раз", redisErr)
		case errors.As(err, DBErr):
			settings.ErrorLog.Printf("DB ERROR (user %s): %v", userId, err)
			msgText = fmt.Sprintf("🗃 Возникла ошибка БД:\n\n%v \n\nПопробуйте выйти в главное меню и попробовать ещё раз", DBErr)
		
		case errors.As(err, validateErr):
			context.Send("😬 Упсс.. Введено некорректное значение! Попробуйте ещё раз 💪")
			return

		case errors.As(err, internalErr):
			settings.ErrorLog.Printf("INTERNAL ERROR (user %s): %v", userId, err)
			msgText = fmt.Sprintf("❌ Возникла внутренняя ошибка:\n\n%v \n\nПопробуйте выйти в главное меню и попробовать ещё раз", internalErr)
		default:
			settings.ErrorLog.Printf("UNKNOWN ERROR (user %s): %v", userId, err)
			msgText = "☠️ Возникла неизвестная ошибка. Попробуйте выйти в главное меню и попробовать ещё раз"
	}
	context.Send(msgText, keyboards.InlineKeyboardToHome)
}
