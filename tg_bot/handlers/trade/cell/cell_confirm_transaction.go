package cell
// номер в диалоге: 4

import (
	"fmt"
	"strings"
	"time"

	telebot "gopkg.in/telebot.v3"

	apiClient "github.com/ej-you/HamstersShaver/tg_bot/api_client"
	stateMachine "github.com/ej-you/HamstersShaver/tg_bot/state_machine"
	"github.com/ej-you/HamstersShaver/tg_bot/keyboards"
	"github.com/ej-you/HamstersShaver/tg_bot/services"
)


func CellConfirmTransactionHandler(context telebot.Context) error {
	var err error
	userId := services.GetUserID(context.Chat())

	// парсим значение процента проскальзывания из строки (для проверки, что введено корректное число)
	stringSlippage, err := services.ParseSlippageAmount(strings.TrimSpace(context.Message().Text))
	if err != nil {
		return fmt.Errorf("ConfirmTransactionHandler for user %s: %w", userId, err)
	}

	// получение машины состояний текущего юзера
	userStateMachine := stateMachine.UserStateMachines.Get(userId)
	// установка нового состояния
	if err = userStateMachine.SetStatus("cell_confirm_transaction"); err != nil {
		return fmt.Errorf("ConfirmTransactionHandler for user %s: %w", userId, err)
	}
	// установка значения процента проскальзывания
	if err = userStateMachine.SetSlippage(stringSlippage); err != nil {
		return fmt.Errorf("ConfirmTransactionHandler for user %s: %w", userId, err)
	}

	msgText := fmt.Sprintf("🏁 Процент проскальзывания: %s%% \n\nСбор данных для новой транзакции...", stringSlippage)
	context.Send(msgText)

	// вызов функции для подтверждения транзакции
	return confirmNewTransaction(context, userStateMachine, userId)
}


// подтверждение транзакции
func confirmNewTransaction(context telebot.Context, userStateMachine stateMachine.UserStateMachine, userId string) error {
	newTransInfo, err := userStateMachine.GetNewTransactionPreparation()
	if err != nil {
		return fmt.Errorf("ConfirmTransactionHandler for user %s: %w", userId, err)
	}

	// запрос на получение информации о последующей транзакции продажи монет по собранным данным
	var cellPreRequestInfo apiClient.PreRequestCellJetton
	getCellPreRequestInfoParams := apiClient.QueryParams{Params: map[string]interface{}{
		"JettonCA": newTransInfo.JettonCA,
		"Amount": newTransInfo.Amount,
		"Slippage": newTransInfo.Slippage,
	}}
	err = apiClient.GetRequest("/api/transactions/cell/pre-request", &getCellPreRequestInfoParams, &cellPreRequestInfo, 5*time.Second)
	if err != nil {
		return fmt.Errorf("ConfirmTransactionHandler for user %s: %w", userId, err)
	}

	msgText := fmt.Sprintf(`🔁 Подтверждение транзакции продажи монет:

Монета: %s
Адрес монеты: %s
DEX-биржа: %s

Монет на продажу: %s
Проскальзывание: %s%%
Примерное количество TON, которые будут получены после транзакции: %s

Подтвердите проведение данной транзакции 👇
`,
		cellPreRequestInfo.JettonSymbol,
		cellPreRequestInfo.JettonCA,
		cellPreRequestInfo.DEX,
		cellPreRequestInfo.UsedJettons,
		newTransInfo.Slippage,
		cellPreRequestInfo.TONsOut,
	)

	return context.Send(msgText, keyboards.InlineKeyboardConfirmNewTransaction)
}
