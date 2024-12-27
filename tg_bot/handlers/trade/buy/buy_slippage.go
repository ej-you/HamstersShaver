package buy
// номер в диалоге: 2

import (
	"fmt"
	"strings"

	telebot "gopkg.in/telebot.v3"

	apiClient "github.com/ej-you/HamstersShaver/tg_bot/api_client"
	stateMachine "github.com/ej-you/HamstersShaver/tg_bot/state_machine"
	"github.com/ej-you/HamstersShaver/tg_bot/keyboards"
	"github.com/ej-you/HamstersShaver/tg_bot/services"
)


func BuySlippageHandler(context telebot.Context) error {
	var err error

	// парсинг значения из строки, введённой юзером
	tonsAmount, err := services.ParseJettonsAmount(apiClient.TONMasterAddress, strings.TrimSpace(context.Message().Text))
	if err != nil {
		return fmt.Errorf("BuySlippageHandler: %w", err)
	}

	// получение машины состояний текущего юзера
	userStateMachine := stateMachine.UserStateMachines.Get(services.GetUserID(context.Chat()))
	// установка нового состояния
	if err = userStateMachine.SetStatus("buy_slippage"); err != nil {
		return fmt.Errorf("BuySlippageHandler: %w", err)
	}
	// установка значения количества TON
	if err = userStateMachine.SetJettonsAmount(tonsAmount); err != nil {
		return fmt.Errorf("BuySlippageHandler: %w", err)
	}

	msgText := fmt.Sprintf(`🫰 Количество используемых TON: %s

Теперь введите процент проскальзывания или выберите из предложенных вариантов 👇`, tonsAmount)

	return context.Send(msgText, keyboards.InlineKeyboardSlippageChoices)
}
