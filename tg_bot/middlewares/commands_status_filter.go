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
			case "/start", "/help", "/home", "/cancel":
				// при любом статусе
				return nextHandler(context)
			case "/trade", "/auto", "/tokens":
				// статус "home"
				accepted, err = userStateMachine.StatusEquals("home")
			case "/buy", "/cell":
				// статус "home" или "trade"
				accepted, err = tradeSubfuncsCommandsStatusFilter(userStateMachine)
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

// возвращает true, если при текущем статусе можно использовать данную команду
func tradeSubfuncsCommandsStatusFilter(userStateMachine stateMachine.UserStateMachine) (bool, error) {
	equals, err := userStateMachine.StatusEquals("trade")
	if err != nil {
		return false, err
	}
	if equals {
		return true, nil
	}

	equals, err = userStateMachine.StatusEquals("home")
	if err != nil {
		return false, err
	}
	if equals {
		return true, nil
	}

	return false, nil
}
