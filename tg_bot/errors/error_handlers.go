package errors

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/google/uuid"
	telebot "gopkg.in/telebot.v3"

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

	accessErr := new(AccessError)
	validateErr := new(ValidateError)
	lastTransNotFinishedErr := new(LastTransNotFinishedError)
	internalErr := new(InternalError)
	restAPIErr := new(RestAPIError)
	restAPITimeoutErr := new(RestAPITimeoutError)
	redisErr := new(RedisError)
	dbErr := new(DBError)
	dbNotFoundErr := new(DBNotFoundError)

	var msgText string
	switch {
		case errors.As(err, accessErr):
			settings.InfoLog.Printf("USER BLOCKED: %v", err)
			return

		case errors.As(err, validateErr):
			context.Send("😬 Упсс.. Введено некорректное значение! Попробуйте ещё раз 💪")
			return
		case errors.As(err, lastTransNotFinishedErr):
			context.Send("🫷 Стойте. Нельзя начать новую транзакцию. Подождите завершения предыдущей 😉")
			return

		case errors.As(err, internalErr):
			settings.ErrorLog.Printf("INTERNAL ERROR (user %d): %v", userId, err)
			msgText = fmt.Sprintf("❌ Возникла внутренняя ошибка:\n\n%v \n\nПопробуйте выйти в главное меню и попробовать ещё раз", internalErr)

		case errors.As(err, restAPIErr):
			settings.ErrorLog.Printf("REST API ERROR (user %d): %v", userId, err)
			msgText = fmt.Sprintf("🛠 Возникла ошибка API:\n\n%v \n\nПопробуйте выйти в главное меню и попробовать ещё раз", restAPIErr)
		case errors.As(err, restAPITimeoutErr):
			settings.ErrorLog.Printf("REST API TIMEOUT ERROR (user %d): %v", userId, err)
			msgText = fmt.Sprintf("⌛️ Возникла ошибка ожидания API:\n\n%v \n\nПопробуйте выйти в главное меню и попробовать ещё раз", restAPITimeoutErr)

		case errors.As(err, redisErr):
			settings.ErrorLog.Printf("REDIS ERROR (user %d): %v", userId, err)
			msgText = fmt.Sprintf("☁️ Возникла ошибка кэша:\n\n%v \n\nПопробуйте выйти в главное меню и попробовать ещё раз", redisErr)
		case errors.As(err, dbErr):
			settings.ErrorLog.Printf("DB ERROR (user %d): %v", userId, err)
			msgText = fmt.Sprintf("🗃 Возникла ошибка БД:\n\n%v \n\nПопробуйте выйти в главное меню и попробовать ещё раз", dbErr)
		case errors.As(err, dbNotFoundErr):
			settings.ErrorLog.Printf("DB ERROR (user %d): %v", userId, err)
			msgText = fmt.Sprintf("🗃 Возникла ошибка БД:\n\n%v \n\nПопробуйте выйти в главное меню и попробовать ещё раз", dbNotFoundErr)

		default:
			settings.ErrorLog.Printf("UNKNOWN ERROR (user %d): %v", userId, err)
			msgText = "☠️ Возникла неизвестная ошибка. Попробуйте выйти в главное меню и попробовать ещё раз"
	}
	context.Send(msgText, inlineKeyboardToHome)
}


// обработчик ошибок в фоновых функциях
func BackgroundErrorHandler(action string, actionId uuid.UUID, err error, context *telebot.Context) {
	userId := (*context).Chat().ID

	// создание начала текста сообщения с использованием фонового действия и его uuid
	var msgText string
	switch action {
		case "transaction":
			msgText = fmt.Sprintf("Обработка транзакции %s\n\n", actionId)
		default:
			msgText = fmt.Sprintf("Обработка фонового действия %s\n\n", actionId)
	}
	logPrefix := fmt.Sprintf("(Background task: %s | UUID: )", action, actionId.String())

	internalErr := new(InternalError)
	restAPIErr := new(RestAPIError)
	restAPITimeoutErr := new(RestAPITimeoutError)
	redisErr := new(RedisError)
	dbErr := new(DBError)
	dbNotFoundErr := new(DBNotFoundError)

	switch {
		case errors.As(err, internalErr):
			settings.ErrorLog.Printf("%s INTERNAL ERROR (user %d): %v", logPrefix, userId, err)
			msgText += fmt.Sprintf("❌ Возникла внутренняя ошибка:\n\n%v", internalErr)
		
		case errors.As(err, restAPIErr):
			settings.ErrorLog.Printf("%s REST API ERROR (user %d): %v", logPrefix, userId, err)
			msgText += fmt.Sprintf("🛠 Возникла ошибка API:\n\n%v", restAPIErr)
		case errors.As(err, restAPITimeoutErr):
			settings.ErrorLog.Printf("%s REST API TIMEOUT ERROR (user %d): %v", logPrefix, userId, err)
			msgText += fmt.Sprintf("⌛️ Возникла ошибка ожидания API:\n\n%v", restAPITimeoutErr)

		case errors.As(err, redisErr):
			settings.ErrorLog.Printf("%s REDIS ERROR (user %d): %v", logPrefix, userId, err)
			msgText += fmt.Sprintf("☁️ Возникла ошибка кэша:\n\n%v", redisErr)
		case errors.As(err, dbErr):
			settings.ErrorLog.Printf("%s DB ERROR (user %d): %v", logPrefix, userId, err)
			msgText += fmt.Sprintf("🗃 Возникла ошибка БД:\n\n%v", dbErr)
		case errors.As(err, dbNotFoundErr):
			settings.ErrorLog.Printf("%s DB ERROR (user %d): %v", logPrefix, userId, err)
			msgText += fmt.Sprintf("🗃 Возникла ошибка БД:\n\n%v", dbNotFoundErr)

		default:
			settings.ErrorLog.Printf("%s UNKNOWN ERROR (user %d): %v", logPrefix, userId, err)
			msgText += "☠️ Возникла неизвестная ошибка."
	}
	(*context).Send(msgText)
}
