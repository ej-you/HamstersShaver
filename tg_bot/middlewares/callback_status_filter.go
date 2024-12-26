package middlewares

import (
	telebot "gopkg.in/telebot.v3"
	
	"github.com/ej-you/HamstersShaver/tg_bot/services"
	"github.com/ej-you/HamstersShaver/tg_bot/settings"
	stateMachine "github.com/ej-you/HamstersShaver/tg_bot/state_machine"
)


// фильтр по статусу и игнорирование нажатия инлайн-кнопок при несоответствии статуса
func GeneralCallbackStatusFilter(nextHandler telebot.HandlerFunc) telebot.HandlerFunc {
	return func(context telebot.Context) error {
		var err error
		var accepted bool

		telegramUserId := services.GetUserID(context.Chat())
		// получение машины состоянию текущего юзера
		userStateMachine := stateMachine.UserStateMachines.Get(telegramUserId)

		callback := context.Callback().Unique
		switch callback {
			// при любом статусе
			case "hide_help", "to_home":
				return nextHandler(context)
			case "to_trade", "to_auto", "to_tokens":
				accepted, err = userStateMachine.StatusEquals("home")
			case "to_buy", "to_cell":
				accepted, err = userStateMachine.StatusEquals("trade")
			default:
				return nextHandler(context)
		}

		if err != nil {
			return err
		}
		if !accepted {
			settings.InfoLog.Printf("Use button %q by user %s was failed (not accepted with current status)", callback, telegramUserId)
			return nil
		}

		return nextHandler(context)
	}
}
