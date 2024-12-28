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
			// при любом статусе, кроме "start" и "home"
			case "cancel":
				accepted, err = userStateMachine.StatusEquals("start", "home")
				accepted = !accepted
			case "to_trade":
				accepted, err = userStateMachine.StatusEquals("home", "trade")
			case "to_auto":
				accepted, err = userStateMachine.StatusEquals("home", "auto")
			case "to_tokens":
				accepted, err = userStateMachine.StatusEquals("home", "tokens")
			case "to_buy":
				accepted, err = userStateMachine.StatusEquals("trade", "buy")
			case "to_cell":
				accepted, err = userStateMachine.StatusEquals("trade", "cell")
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
