package buy
// номер в диалоге: 2

import (
	"fmt"
	"strings"

	telebot "gopkg.in/telebot.v3"

	apiClient "github.com/ej-you/HamstersShaver/tg_bot/api_client"
	stateMachine "github.com/ej-you/HamstersShaver/tg_bot/state_machine"
	"github.com/ej-you/HamstersShaver/tg_bot/services"
)


func BuySlippageHandler(context telebot.Context) error {
	var err error
	userId := services.GetUserID(context.Chat())

	// парсинг значения из строки, введённой юзером
	tonsAmount, err := services.ParseJettonsAmount(apiClient.TONMasterAddress, strings.TrimSpace(context.Message().Text))
	if err != nil {
		return fmt.Errorf("BuySlippageHandler for user %s: %w", userId, err)
	}

	// получение машины состояний текущего юзера
	userStateMachine := stateMachine.UserStateMachines.Get(userId)
	// установка нового состояния
	if err = userStateMachine.SetStatus("buy_slippage"); err != nil {
		return fmt.Errorf("BuySlippageHandler for user %s: %w", userId, err)
	}
	// установка значения количества TON
	if err = userStateMachine.SetJettonsAmount(tonsAmount); err != nil {
		return fmt.Errorf("BuySlippageHandler for user %s: %w", userId, err)
	}

	msgText := fmt.Sprintf(`🫰 Количество используемых монет: %s

Теперь введите процент проскальзывания (число от 1 до 100)`, tonsAmount)

	return context.Send(msgText)
}
