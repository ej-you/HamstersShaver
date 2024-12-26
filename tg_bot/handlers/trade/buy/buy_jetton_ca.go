package buy
// номер в диалоге: 3

import (
	"fmt"
	"strings"

	telebot "gopkg.in/telebot.v3"

	stateMachine "github.com/ej-you/HamstersShaver/tg_bot/state_machine"
	"github.com/ej-you/HamstersShaver/tg_bot/services"
)


// кнопки выбора монеты для продажи в диалоге команды /cell
func BuyJettonCAHandler(context telebot.Context) error {
	var err error
	userId := services.GetUserID(context.Chat())

	// парсим значение процента проскальзывания из строки (для проверки, что введено корректное число)
	stringSlippage, err := services.ParseSlippageAmount(strings.TrimSpace(context.Message().Text))
	if err != nil {
		return fmt.Errorf("BuyJettonCAHandler for user %s: %w", userId, err)
	}

	// получение машины состояний текущего юзера
	userStateMachine := stateMachine.UserStateMachines.Get(userId)
	// установка нового состояния
	if err = userStateMachine.SetStatus("buy_jetton_ca"); err != nil {
		return fmt.Errorf("BuyJettonCAHandler for user %s: %w", userId, err)
	}
	// установка значения процента проскальзывания
	if err = userStateMachine.SetSlippage(stringSlippage); err != nil {
		return fmt.Errorf("BuyJettonCAHandler for user %s: %w", userId, err)
	}

	return context.Send(fmt.Sprintf("🙂‍↔️ Процент проскальзывания: %s%% \n\nТеперь введите адрес покупаемой монеты", stringSlippage))
}
