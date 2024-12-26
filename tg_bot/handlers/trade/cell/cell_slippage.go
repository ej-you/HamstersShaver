package cell
// номер в диалоге: 3

import (
	"fmt"
	"strings"

	telebot "gopkg.in/telebot.v3"

	stateMachine "github.com/ej-you/HamstersShaver/tg_bot/state_machine"
	"github.com/ej-you/HamstersShaver/tg_bot/services"
)


func CellSlippageHandler(context telebot.Context) error {
	var err error
	userId := services.GetUserID(context.Chat())

	// получение машины состояний текущего юзера
	userStateMachine := stateMachine.UserStateMachines.Get(userId)
	// получение CA монеты из кэша
	jettonCA, err := userStateMachine.GetJettonCA()
	if err != nil {
		return fmt.Errorf("CellSlippageHandler for user %s: %w", jettonCA)
	}

	// парсинг значения из строки, введённой юзером
	jettonsAmount, err := services.ParseJettonsAmount(jettonCA, strings.TrimSpace(context.Message().Text))
	if err != nil {
		return fmt.Errorf("CellSlippageHandler for user %s: %w", userId, err)
	}

	// установка нового состояния
	if err = userStateMachine.SetStatus("cell_slippage"); err != nil {
		return fmt.Errorf("CellSlippageHandler for user %s: %w", userId, err)
	}
	// установка значения количества монет
	if err = userStateMachine.SetJettonsAmount(jettonsAmount); err != nil {
		return fmt.Errorf("CellSlippageHandler for user %s: %w", userId, err)
	}

	msgText := fmt.Sprintf(`🫰 Количество используемых монет: %s

Теперь введите процент проскальзывания (число от 1 до 100)`, jettonsAmount)

	return context.Send(msgText)
}
