package handlers

import (
	"fmt"

	telebot "gopkg.in/telebot.v3"
	"github.com/pkg/errors"

	customErrors "github.com/ej-you/HamstersShaver/tg_bot/errors"
	"github.com/ej-you/HamstersShaver/tg_bot/keyboards"
	"github.com/ej-you/HamstersShaver/tg_bot/settings"
)


// обработчик всех ошибок
func UnknownErrorHandler(err error, context telebot.Context) {
	restAPIErr := new(customErrors.RestAPIError)
	redisErr := new(customErrors.RedisError)
	DBErr := new(customErrors.DBError)
	validateErr := new(customErrors.ValidateError)
	internalErr := new(customErrors.InternalError)

	var msgText string
	switch {
		case errors.As(err, restAPIErr) :
			settings.ErrorLog.Printf("REST API ERROR: %v", err)
			msgText = fmt.Sprintf("🛠 Возникла ошибка API:\n\n%v \n\nПопробуйте выйти в главное меню и попробовать ещё раз", restAPIErr)
		case errors.As(err, redisErr) :
			settings.ErrorLog.Printf("REDIS ERROR: %v", err)
			msgText = fmt.Sprintf("💸 Возникла ошибка кэша:\n\n%v \n\nПопробуйте выйти в главное меню и попробовать ещё раз", redisErr)
		case errors.As(err, DBErr) :
			settings.ErrorLog.Printf("DB ERROR: %v", err)
			msgText = fmt.Sprintf("🗃 Возникла ошибка БД:\n\n%v \n\nПопробуйте выйти в главное меню и попробовать ещё раз", DBErr)
		case errors.As(err, validateErr) :
			settings.ErrorLog.Printf("VALIDATE ERROR: %v", err)
			msgText = fmt.Sprintf("🗑 Возникла ошибка валидации:\n\n%v \n\nПопробуйте выйти в главное меню и попробовать ещё раз", validateErr)
		case errors.As(err, internalErr) :
			settings.ErrorLog.Printf("INTERNAL ERROR: %v", err)
			msgText = fmt.Sprintf("❌ Возникла внутренняя ошибка:\n\n%v \n\nПопробуйте выйти в главное меню и попробовать ещё раз", internalErr)
		default:
			settings.ErrorLog.Printf("UNKNOWN ERROR: %v", err)
			msgText = "☠️ Возникла неизвестная ошибка. Попробуйте выйти в главное меню и попробовать ещё раз"
	}
	context.Send(msgText, keyboards.InlineKeyboardToHome)
}
