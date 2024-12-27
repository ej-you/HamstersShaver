package buy
// номер в диалоге: 4

import (
	"fmt"
	"strings"

	telebot "gopkg.in/telebot.v3"

	apiClient "github.com/ej-you/HamstersShaver/tg_bot/api_client"
	stateMachine "github.com/ej-you/HamstersShaver/tg_bot/state_machine"
	"github.com/ej-you/HamstersShaver/tg_bot/keyboards"
	"github.com/ej-you/HamstersShaver/tg_bot/services"
)


func BuyConfirmTransactionHandler(context telebot.Context) error {
	var err error

	// получение адреса монеты от юзера
	jettonCA := strings.TrimSpace(context.Message().Text)

	// получение машины состояний текущего юзера
	userStateMachine := stateMachine.UserStateMachines.Get(services.GetUserID(context.Chat()))
	// установка нового состояния
	if err = userStateMachine.SetStatus("buy_confirm_transaction"); err != nil {
		return fmt.Errorf("BuyConfirmTransactionHandler: %w", err)
	}
	// установка значения процента проскальзывания
	if err = userStateMachine.SetJettonCA(jettonCA); err != nil {
		return fmt.Errorf("BuyConfirmTransactionHandler: %w", err)
	}

	msgText := fmt.Sprintf("🏁 Адрес монеты: %s \n\nСбор данных для новой транзакции...", jettonCA)
	context.Send(msgText)

	// вызов функции для подтверждения транзакции
	return confirmNewTransaction(context, userStateMachine)
}


// подтверждение транзакции
func confirmNewTransaction(context telebot.Context, userStateMachine stateMachine.UserStateMachine) error {
	newTransInfo, err := userStateMachine.GetNewTransactionPreparation()
	if err != nil {
		return fmt.Errorf("BuyConfirmTransactionHandler: %w", err)
	}

	// запрос на получение информации о последующей транзакции продажи монет по собранным данным
	var buyPreRequestInfo apiClient.PreRequestBuyJetton
	getBuyPreRequestInfoParams := apiClient.QueryParams{Params: map[string]interface{}{
		"JettonCA": newTransInfo.JettonCA,
		"Amount": newTransInfo.Amount,
		"Slippage": newTransInfo.Slippage,
	}}
	err = apiClient.GetRequest("/api/transactions/buy/pre-request", &getBuyPreRequestInfoParams, &buyPreRequestInfo)
	if err != nil {
		return fmt.Errorf("BuyConfirmTransactionHandler: %w", err)
	}

	msgText := fmt.Sprintf(`🔁 Подтверждение транзакции покупки монет:

Покупаемая монета: %s
Адрес монеты: %s
DEX-биржа: %s

TON для покупки: %s
Проскальзывание: %s%%
Примерное количество монет, которые будут получены после транзакции: %s

Подтвердите проведение данной транзакции 👇
`,
		buyPreRequestInfo.JettonSymbol,
		buyPreRequestInfo.JettonCA,
		newTransInfo.DEX,
		buyPreRequestInfo.UsedTON,
		newTransInfo.Slippage,
		buyPreRequestInfo.JettonsOut,
	)

	return context.Send(msgText, keyboards.InlineKeyboardConfirmNewTransaction)
}
