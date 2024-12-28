package middlewares

import (
	"strings"

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

		// если сообщение является командой
		if strings.HasPrefix(message, "/") {
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
					accepted, err = userStateMachine.StatusEquals("home", "trade", "buy", "cell")
				case "/cell":
					accepted, err = userStateMachine.StatusEquals("home", "trade", "cell", "buy")
			}
		// если сообщение - простой текст
		} else {
			// разрешаем вводить простой текст только при статусах, на которых нужно вручную вводить значения
			accepted, err = userStateMachine.StatusEquals(
				"buy_tons_amount", "buy_slippage", "buy_jetton_ca", // диалог покупки монет
				"cell_jettons_amount", "cell_slippage", // диалог продажи монет
			)
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
