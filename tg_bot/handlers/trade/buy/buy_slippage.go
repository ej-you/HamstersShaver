package buy
// номер в диалоге: 2

import (
	"fmt"
	"strings"

	telebot "gopkg.in/telebot.v3"

	stateMachine "github.com/ej-you/HamstersShaver/tg_bot/state_machine"
	"github.com/ej-you/HamstersShaver/tg_bot/keyboards"
	"github.com/ej-you/HamstersShaver/tg_bot/services"
)


func BuySlippageHandler(context telebot.Context) error {
	var err error

	var tonsAmount string
	callback := context.Callback()

	// парсинг значения кол-ва TON из текста или нажатой кнопки
	if callback == nil {
		tonsAmount, err = services.ParseTonAmount(strings.TrimSpace(context.Message().Text))
	} else {
		callbackData := services.GetCallbackData(context.Callback())
		// если нажата левая кнопка (не с выбором процента кол-ва монет на продажу)		
		if !strings.HasPrefix(callbackData, "jettons_amount_choice") {
			return nil
		}
		tonsAmount, err = services.GetTonAmountFromPercentFromCallback(callbackData)
	}
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
