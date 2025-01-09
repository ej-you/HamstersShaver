package cell
// номер в диалоге: 3

import (
	"fmt"
	"strings"

	telebot "gopkg.in/telebot.v3"

	stateMachine "github.com/ej-you/HamstersShaver/tg_bot/state_machine"
	"github.com/ej-you/HamstersShaver/tg_bot/keyboards"
	"github.com/ej-you/HamstersShaver/tg_bot/services"
)


func CellSlippageHandler(context telebot.Context) error {
	var err error

	// получение машины состояний текущего юзера
	userStateMachine := stateMachine.UserStateMachines.Get(services.GetUserID(context.Chat()))
	// получение CA монеты из кэша
	jettonCA, err := userStateMachine.GetJettonCA()
	if err != nil {
		return fmt.Errorf("CellSlippageHandler: %w", err)
	}

	var jettonsAmount string
	callback := context.Callback()
	// парсинг значения кол-ва монет из текста или нажатой кнопки
	if callback == nil {
		jettonsAmount, err = services.ParseJettonsAmount(jettonCA, strings.TrimSpace(context.Message().Text))
	} else {
		callbackData := services.GetCallbackData(context.Callback())
		// если нажата левая кнопка (не с выбором процента кол-ва монет на продажу)		
		if !strings.HasPrefix(callbackData, "jettons_amount_choice") {
			return nil
		}
		jettonsAmount, err = services.GetJettonAmountFromPercentFromCallback(jettonCA, callbackData)
	}
	if err != nil {
		return fmt.Errorf("CellSlippageHandler: %w", err)
	}

	// установка нового состояния
	if err = userStateMachine.SetStatus("cell_slippage"); err != nil {
		return fmt.Errorf("CellSlippageHandler: %w", err)
	}
	// установка значения количества монет
	if err = userStateMachine.SetJettonsAmount(jettonsAmount); err != nil {
		return fmt.Errorf("CellSlippageHandler: %w", err)
	}

	msgText := fmt.Sprintf(`🫰 Количество используемых монет: %s

Теперь введите процент проскальзывания или выберите из предложенных вариантов 👇`, jettonsAmount)

	return context.Send(msgText, keyboards.InlineKeyboardSlippageChoices)
}
