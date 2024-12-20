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
			case "hide_help":
				// при любом статусе
				return nextHandler(context)
			case "to_buy", "to_cell":
				// статус "trade"
				accepted, err = userStateMachine.StatusEquals("trade")
			case "to_trade", "to_auto", "to_tokens":
				// статус "home"
				accepted, err = userStateMachine.StatusEquals("home")
			case "to_home":
				// статус "start", "in_development" и "home"
				accepted, err = toHomeButtonStatusFilter(userStateMachine)
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

// возвращает true, если при текущем статусе можно использовать данную команду
func toHomeButtonStatusFilter(userStateMachine stateMachine.UserStateMachine) (bool, error) {
	equals, err := userStateMachine.StatusEquals("start")
	if err != nil {
		return false, err
	}
	if equals {
		return true, nil
	}

	equals, err = userStateMachine.StatusEquals("in_development")
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