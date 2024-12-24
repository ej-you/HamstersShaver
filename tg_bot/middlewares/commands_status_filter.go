package middlewares

import (
	telebot "gopkg.in/telebot.v3"
	
	"github.com/ej-you/HamstersShaver/tg_bot/services"
	"github.com/ej-you/HamstersShaver/tg_bot/settings"
	stateMachine "github.com/ej-you/HamstersShaver/tg_bot/state_machine"
)


// фильтр по статусу и игнорирование команд при несоответствии статуса
func GeneralCommandsStatusFilter(nextHandler telebot.HandlerFunc) telebot.HandlerFunc {
	return func(context telebot.Context) error {
		var err error
		var accepted bool

		telegramUserId := services.GetUserID(context.Chat())
		// получение машины состоянию текущего юзера
		userStateMachine := stateMachine.UserStateMachines.Get(telegramUserId)

		message := context.Message().Text
		switch message {
			// при любом статусе
			case "/start", "/help", "/home":
				return nextHandler(context)
			// при любом статусе, кроме "start" и "home"
			case "/cancel":
				accepted, err = userStateMachine.StatusEquals("start", "home")
				accepted = !accepted
			case "/trade":
				accepted, err = userStateMachine.StatusEquals("home", "trade")
			case "/auto":
				accepted, err = userStateMachine.StatusEquals("home", "auto")
			case "/tokens":
				accepted, err = userStateMachine.StatusEquals("home", "tokens")
			case "/buy":
				accepted, err = userStateMachine.StatusEquals("home", "trade", "buy")
			case "/cell":
				accepted, err = userStateMachine.StatusEquals("home", "trade", "cell")
		}

		if err != nil {
			return err
		}
		if !accepted {
			settings.InfoLog.Printf("Use command %q by user %s was failed (not accepted with current status)", message, telegramUserId)
			// также удаляем игнорируемое сообщение от юзера
			context.Delete()
			return nil
		}

		return nextHandler(context)
	}
}
