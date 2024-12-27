package handlers

import (
	"fmt"

	telebot "gopkg.in/telebot.v3"

	"github.com/ej-you/HamstersShaver/tg_bot/keyboards"
	"github.com/ej-you/HamstersShaver/tg_bot/services"
	stateMachine "github.com/ej-you/HamstersShaver/tg_bot/state_machine"
)


// команды: /start
func StartHandler(context telebot.Context) error {
	var err error

	// получение машины состояний текущего юзера
	userStateMachine := stateMachine.UserStateMachines.Get(services.GetUserID(context.Chat()))
	// установка нового состояния
	if err = userStateMachine.SetStatus("start"); err != nil {
		return fmt.Errorf("StartHandler: %w", err)
	}
	// очистка кэша с информацией для новой транзакции
	if err = userStateMachine.ClearNewTransactionPreparation(); err != nil {
		return fmt.Errorf("StartHandler: %w", err)
	}

	msgText := `Привет 👋

🥳 Это бот для быстрых транзакций на TON! 🥳
С его помощью можно проводить транзакции покупки и продажи монет быстро и максимально эффективно - без посредников 🤑

❗️Для получения полной инструкции введите /help`

	return context.Send(msgText, keyboards.InlineKeyboardToHome)
}
