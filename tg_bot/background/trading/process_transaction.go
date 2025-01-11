package trading

import (
	"fmt"
	"strconv"
	"time"

	telebot "gopkg.in/telebot.v3"

	apiClient "github.com/ej-you/HamstersShaver/tg_bot/api_client"
	customErrors "github.com/ej-you/HamstersShaver/tg_bot/errors"
	stateMachine "github.com/ej-you/HamstersShaver/tg_bot/state_machine"
	"github.com/ej-you/HamstersShaver/tg_bot/keyboards"
	"github.com/ej-you/HamstersShaver/tg_bot/settings"
)


const waitSeqnoIncrementTimes = 6


// изменение сообщения, отправленного после подтверждения транзакции, на ошибку
func editSentMessageToError(context *telebot.Context, sentTransMsg *telebot.Message) {
	(*context).Bot().Edit(sentTransMsg, "🤷‍♂️ Упс... Произошла ошибка 👆", keyboards.InlineKeyboardToHome)
}

// ожидание инкрементации seqno в течение ~30 секунд
func waitSeqnoIncrement(seqnoBeforeTrans apiClient.AccountSeqno) error {
	var seqnoAfterTrans apiClient.AccountSeqno
	var seqnoErr error

	for i := 0; i < waitSeqnoIncrementTimes; i++ {
		// получение seqno аккаунта после отправки транзакции
		seqnoErr = apiClient.GetRequest("/api/account/get-seqno", nil, &seqnoAfterTrans)
		if seqnoErr == nil && seqnoAfterTrans.Seqno > seqnoBeforeTrans.Seqno { // NOT err
			break
		}
		time.Sleep(5*time.Second)
	}
	// если все попытки были неуспешными
	if seqnoErr != nil {
		return fmt.Errorf("wait seqno increment: %w", seqnoErr)
	}
	// если seqno так и не увеличился
	if seqnoAfterTrans.Seqno == seqnoBeforeTrans.Seqno {
		return fmt.Errorf("wait seqno increment: %w", customErrors.InternalError("wait process transaction in mempool: timeout"))
	}
	return nil
}


// вся обработка транзакции в фоне
func ProcessTransaction(context *telebot.Context, sentTransMsg *telebot.Message, transInfo stateMachine.NewTransactionPreparation, transactionUUID string) {
	// получение seqno аккаунта до проведения транзакции
	var seqnoBeforeTrans apiClient.AccountSeqno
	err := apiClient.GetRequest("/api/account/get-seqno", nil, &seqnoBeforeTrans)
	if err != nil {
		editSentMessageToError(context, sentTransMsg)
		go customErrors.BackgroundErrorHandler("transaction", transactionUUID, fmt.Errorf("processTransaction: %w", err), context)
		return
	}

	// перевод кол-ва монет во float64
	amountFloat64, _ := strconv.ParseFloat(transInfo.Amount, 64)
	// перевод процента проскальзывания в число
	slippageInt, _ := strconv.Atoi(transInfo.Slippage)

	// POST-запрос для отправки транзакции в блокчейн
	postSendTransData := apiClient.JsonBody{
		"jettonCA": transInfo.JettonCA,
		"amount": amountFloat64,
		"slippage": slippageInt,
	}
	err = apiClient.PostRequest(fmt.Sprintf("/api/transactions/%s/send", transInfo.Action), &postSendTransData, nil)
	if err != nil {
		editSentMessageToError(context, sentTransMsg)
		go customErrors.BackgroundErrorHandler("transaction", transactionUUID, fmt.Errorf("processTransaction: %w", err), context)
		return
	}

	// изменение сообщения на "транзакция в mempool"
	settings.InfoLog.Printf("Transaction %q: was sent to mempool", transactionUUID)
	(*context).Bot().Edit(sentTransMsg, "⏸️ Транзакция отправлена в mempool 👆", keyboards.InlineKeyboardToHome)

	// ожидание инкрементации seqno в течение ~30 секунд
	if err = waitSeqnoIncrement(seqnoBeforeTrans); err != nil {
		editSentMessageToError(context, sentTransMsg)
		go customErrors.BackgroundErrorHandler("transaction", transactionUUID, fmt.Errorf("processTransaction: %w", err), context)
		return
	}

	// изменение сообщения на "ожидание окончания транзакции"
	settings.InfoLog.Printf("Transaction %q: seqno was incremented", transactionUUID)
	(*context).Bot().Edit(sentTransMsg, "🔄 Ожидание окончания транзакции... 👆", keyboards.InlineKeyboardToHome)

	// ожидание окончания следующей транзакции
	var waitedTransHash apiClient.WaitTransactionHash
	err = apiClient.SseRequest("/api/transactions/wait-next", &waitedTransHash)
	if err != nil {
		editSentMessageToError(context, sentTransMsg)
		go customErrors.BackgroundErrorHandler("transaction", transactionUUID, fmt.Errorf("processTransaction: %w", err), context)
		return
	}

	// небольшая пауза, потому что без неё не успевает обработаться информация о новой транзакции и
	// функция получения информации по хэшу транзакции возвращает ошибку
	time.Sleep(2*time.Second)
	// изменение сообщения на "транзакция завершена"
	settings.InfoLog.Printf("Transaction %q: was finished", transactionUUID)
	(*context).Bot().Edit(sentTransMsg, "✅ Транзакция завершена! 👆", keyboards.InlineKeyboardToHome)

	// получение информации по хэшу отловленной транзакции
	var endTransInfo apiClient.TransactionInfo
	getEndTransInfoParams := apiClient.QueryParams{
		"TransactionHash": waitedTransHash.Hash,
		"Action": transInfo.Action,
	}
	err = apiClient.GetRequest("/api/transactions/info", &getEndTransInfoParams, &endTransInfo)
	if err != nil {
		go customErrors.BackgroundErrorHandler("transaction", transactionUUID, fmt.Errorf("processTransaction: %w", err), context)
		return
	}

	// разные способы получения информации о монете в зависимости от успеха/неудачи проведённой транзакции
	getJettonInfoParams := apiClient.QueryParams{"MasterAddress": transInfo.JettonCA}
	// используем структуру AccountJetton и для случая с JettonInfo, потому что они имеют общие используемые поля Symbol и MasterAddress
	var jettonInfo apiClient.AccountJetton
	var beautyTransResult string
	var newJettonBalance string
	if endTransInfo.StatusOK == true {
		// получение информации о монете аккаунта по её адресу
		err = apiClient.GetRequest("/api/account/get-jetton", &getJettonInfoParams, &jettonInfo)
		if err != nil {
			go customErrors.BackgroundErrorHandler("transaction", transactionUUID, fmt.Errorf("processTransaction: %w", err), context)
			return
		}
		beautyTransResult = "успешно ✅"
		newJettonBalance = fmt.Sprintf("Новый баланс монеты: %s", jettonInfo.BeautyBalance)
	} else {
		// получение информации о монете по её адресу
		err = apiClient.GetRequest("/api/jettons/get-info", &getJettonInfoParams, &jettonInfo)
		if err != nil {
			go customErrors.BackgroundErrorHandler("transaction", transactionUUID, fmt.Errorf("processTransaction: %w", err), context)
			return
		}
		beautyTransResult = "неудачно ❌"
	}

	// данные для сообщения в красивом виде
	beautyAction := "покупка монет"
	beautyWhatUsed := "TON для покупки"
	if endTransInfo.Action == "cell" {
		beautyAction = "продажа монет"
		beautyWhatUsed = "Монет на продажу"
	}

	// составление текста сообщения
	msgText := fmt.Sprintf(`💸 Транзакция завершена!

Действие: %s
DEX-биржа: %s
Результат: %s

Монета: %s
Адрес монеты: %s
%s: %s

Новый баланс TON: %s
%s
`,
		beautyAction,
		transInfo.DEX,
		beautyTransResult,

		jettonInfo.Symbol,
		jettonInfo.MasterAddress,
		beautyWhatUsed, transInfo.Amount,

		endTransInfo.EndBalance,
		newJettonBalance,
	)

	// отправка нового сообщения с данными о закончившейся транзакции
	keyboards.SetTonviewerTransLink(endTransInfo.Hash)
	(*context).Send(msgText, keyboards.InlineKeyboardTonviewerTransLink)
}
