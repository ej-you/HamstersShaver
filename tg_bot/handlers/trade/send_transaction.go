package trade

import (
	"fmt"

	telebot "gopkg.in/telebot.v3"

	backgroundTrading "github.com/ej-you/HamstersShaver/tg_bot/background/trading"
	stateMachine "github.com/ej-you/HamstersShaver/tg_bot/state_machine"
	"github.com/ej-you/HamstersShaver/tg_bot/keyboards"
	"github.com/ej-you/HamstersShaver/tg_bot/services"
)


// отправка подготовленной транзакции
func SendTransaction(context telebot.Context) error {
	var err error

	// получение машины состояний текущего юзера
	userStateMachine := stateMachine.UserStateMachines.Get(services.GetUserID(context.Chat()))
	// установка нового состояния
	if err = userStateMachine.SetStatus("send_transaction"); err != nil {
		return fmt.Errorf("SendTransaction: %w", err)
	}

	// получение данных для новой транзакции
	newTrans, err := userStateMachine.GetNewTransactionPreparation()
	if err != nil {
		return fmt.Errorf("SendTransaction: %w", err)
	}

	// сообщение о начале обработки транзакции
	transMsg := "▶️ Начата обработка транзакции выше... 👆"
	sentTransMsg, err := context.Bot().Send(context.Recipient(), transMsg, keyboards.InlineKeyboardToHome)
	if err != nil {
		return fmt.Errorf("SendTransaction (start processing): %w", err)
	}

	// запуск обработки в фоне
	go backgroundTrading.ProcessTransaction(context, sentTransMsg, newTrans)
	return nil
}
