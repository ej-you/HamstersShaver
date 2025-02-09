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
	var stringSlippage string

	// если нажата кнопка для выбора процента
	if context.Callback() != nil {
		callbackData := services.GetCallbackData(context.Callback())
		// если нажата левая кнопка (не с выбором процента проскальзывания)		
		if !strings.HasPrefix(callbackData, "slippage_choice") {
			return nil
		}
		// достаём процент из данных кнопки
		stringSlippage = strings.TrimPrefix(callbackData, "slippage_choice|")
	// если процент введён текстом
	} else {
		// парсим значение процента проскальзывания из строки (для проверки, что введено корректное число)
		stringSlippage, err = services.ParseSlippageAmount(strings.TrimSpace(context.Message().Text))
		if err != nil {
			return fmt.Errorf("BuyJettonCAHandler: %w", err)
		}
	}

	// получение машины состояний текущего юзера
	userStateMachine := stateMachine.UserStateMachines.Get(services.GetUserID(context.Chat()))
	// установка нового состояния
	if err = userStateMachine.SetStatus("buy_jetton_ca"); err != nil {
		return fmt.Errorf("BuyJettonCAHandler: %w", err)
	}
	// установка значения процента проскальзывания
	if err = userStateMachine.SetSlippage(stringSlippage); err != nil {
		return fmt.Errorf("BuyJettonCAHandler: %w", err)
	}

	return context.Send(fmt.Sprintf("🙂‍↔️ Процент проскальзывания: %s%% \n\nТеперь введите адрес покупаемой монеты", stringSlippage))
}
