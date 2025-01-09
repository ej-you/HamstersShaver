package cell
// номер в диалоге: 4

import (
	"fmt"
	"strings"

	telebot "gopkg.in/telebot.v3"

	apiClient "github.com/ej-you/HamstersShaver/tg_bot/api_client"
	customErrors "github.com/ej-you/HamstersShaver/tg_bot/errors"
	stateMachine "github.com/ej-you/HamstersShaver/tg_bot/state_machine"
	"github.com/ej-you/HamstersShaver/tg_bot/keyboards"
	"github.com/ej-you/HamstersShaver/tg_bot/services"

	"github.com/ej-you/HamstersShaver/tg_bot/settings/constants"
)


func CellConfirmTransactionHandler(context telebot.Context) error {
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
			return fmt.Errorf("CellConfirmTransactionHandler: %w", err)
		}
	}

	// получение машины состояний текущего юзера
	userStateMachine := stateMachine.UserStateMachines.Get(services.GetUserID(context.Chat()))
	// установка нового состояния
	if err = userStateMachine.SetStatus("cell_confirm_transaction"); err != nil {
		return fmt.Errorf("CellConfirmTransactionHandler: %w", err)
	}
	// установка значения процента проскальзывания
	if err = userStateMachine.SetSlippage(stringSlippage); err != nil {
		return fmt.Errorf("CellConfirmTransactionHandler: %w", err)
	}

	msgText := fmt.Sprintf("🏁 Процент проскальзывания: %s%% \n\nСбор данных для новой транзакции...", stringSlippage)
	context.Send(msgText)

	// вызов функции для подтверждения транзакции
	return confirmNewTransaction(context, userStateMachine)
}


// подтверждение транзакции продажи
func confirmNewTransaction(context telebot.Context, userStateMachine stateMachine.UserStateMachine) error {
	newTransInfo, err := userStateMachine.GetNewTransactionPreparation()
	if err != nil {
		return fmt.Errorf("CellConfirmTransactionHandler: %w", err)
	}

	// получение баланса TON у аккаунта
	var TONAccountInfo apiClient.TONInfo
	err = apiClient.GetRequest("/api/account/get-ton", nil, &TONAccountInfo)
	if err != nil {
		return fmt.Errorf("CellConfirmTransactionHandler: %w", err)
	}

	// проверяем, что TON хватит для газовой комиссии
	tonBalanceFloat := services.ConvertBalanceToFloat64(TONAccountInfo.Balance, TONAccountInfo.Decimals)
	
	fmt.Println("tonBalanceFloat:", tonBalanceFloat)

	if tonBalanceFloat < constants.GasAmountFloat64  {
		internalErr := customErrors.InternalError("you must have minimun 0.3 TON for create transaction")
		return fmt.Errorf("CellConfirmTransactionHandler: not enough TONs for create transaction (balance - %s TON): %w", TONAccountInfo.BeautyBalance, internalErr)
	}

	// запрос на получение информации о последующей транзакции продажи монет по собранным данным
	var cellPreRequestInfo apiClient.PreRequestCellJetton
	getCellPreRequestInfoParams := apiClient.QueryParams{Params: map[string]interface{}{
		"JettonCA": newTransInfo.JettonCA,
		"Amount": newTransInfo.Amount,
		"Slippage": newTransInfo.Slippage,
	}}
	err = apiClient.GetRequest("/api/transactions/cell/pre-request", &getCellPreRequestInfoParams, &cellPreRequestInfo)
	if err != nil {
		return fmt.Errorf("CellConfirmTransactionHandler: %w", err)
	}

	msgText := fmt.Sprintf(`⏩ Подтверждение транзакции продажи монет:

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
		newTransInfo.DEX,
		cellPreRequestInfo.UsedJettons,
		newTransInfo.Slippage,
		cellPreRequestInfo.TONsOut,
	)

	return context.Send(msgText, keyboards.InlineKeyboardConfirmNewTransaction)
}
