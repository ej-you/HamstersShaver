package cell
// номер в диалоге: 3

import (
	"fmt"
	"strconv"
	"strings"

	telebot "gopkg.in/telebot.v3"

	customErrors "github.com/ej-you/HamstersShaver/tg_bot/errors"
	stateMachine "github.com/ej-you/HamstersShaver/tg_bot/state_machine"
	"github.com/ej-you/HamstersShaver/tg_bot/keyboards"
	"github.com/ej-you/HamstersShaver/tg_bot/services"
)


func CellSlippageHandler(context telebot.Context) error {
	var err error
	var jettonsAmount string

	// получение машины состояний текущего юзера
	userStateMachine := stateMachine.UserStateMachines.Get(services.GetUserID(context.Chat()))
	// получение CA монеты из кэша
	jettonCA, err := userStateMachine.GetJettonCA()
	if err != nil {
		return fmt.Errorf("CellSlippageHandler: %w", err)
	}


	// если нажата кнопка для выбора процента кол-ва монет на продажу
	if context.Callback() != nil {
		callbackData := services.GetCallbackData(context.Callback())
		// если нажата левая кнопка (не с выбором процента кол-ва монет на продажу)		
		if !strings.HasPrefix(callbackData, "jettons_amount_choice") {
			return nil
		}

		// достаём процент из данных кнопки и переводим его в int
		intPercent, err := strconv.Atoi(strings.TrimPrefix(callbackData, "jettons_amount_choice|"))
		if err != nil {
			internalErr := customErrors.InternalError("failed to parse percent value")
			return fmt.Errorf("CellSlippageHandler: %w", fmt.Errorf("parse int from jettons_amount_choice button: %v: %w", err, internalErr))
		}
		// переводим процент в кол-во монет
		jettonsAmount, err = services.GetJettonsAmountFromPercent(jettonCA, intPercent)
		if err != nil {
			return fmt.Errorf("CellSlippageHandler: %w", err)
		}

	// если введён текст
	} else {
		// парсим значение из строки, введённой юзером
		jettonsAmount, err = services.ParseJettonsAmount(jettonCA, strings.TrimSpace(context.Message().Text))
		if err != nil {
			return fmt.Errorf("CellSlippageHandler: %w", err)
		}
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
