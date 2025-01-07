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
)


const waitSeqnoIncrementTimes = 6


// вся обработка транзакции в фоне
func ProcessTransaction(context telebot.Context, sentTransMsg *telebot.Message, transInfo stateMachine.NewTransactionPreparation) {
	// получение seqno аккаунта до проведения транзакции
	var seqnoBeforeTrans, seqnoAfterTrans apiClient.AccountSeqno
	err := apiClient.GetRequest("/api/account/get-seqno", nil, &seqnoBeforeTrans)
	if err != nil {
		go customErrors.MainErrorHandler(fmt.Errorf("processTransaction: %w", err), context)
		return
	}

	// перевод кол-ва монет во float64
	amountFloat64, err := strconv.ParseFloat(transInfo.Amount, 64)
	if err != nil {
		internalErr := customErrors.InternalError("failed to parse amount to float value")
		go customErrors.MainErrorHandler(fmt.Errorf("processTransaction: %v: %w", err, internalErr), context)
		return
	}
	// перевод процента проскальзывания в число
	slippageInt, err := strconv.Atoi(transInfo.Slippage)
	if err != nil {
		internalErr := customErrors.InternalError("failed to parse slippage to int value")
		go customErrors.MainErrorHandler(fmt.Errorf("processTransaction: %v: %w", err, internalErr), context)
		return
	}

	// POST-запрос для отправки транзакции в блокчейн
	postSendTransData := apiClient.JsonBody{Data: map[string]interface{}{
		"jettonCA": transInfo.JettonCA,
		"amount": amountFloat64,
		"slippage": slippageInt,
	}}
	err = apiClient.PostRequest(fmt.Sprintf("/api/transactions/%s/send", transInfo.Action), &postSendTransData, nil)
	if err != nil {
		go customErrors.MainErrorHandler(fmt.Errorf("processTransaction: %w", err), context)
		return
	}

	// изменение сообщения на "транзакция в mempool"
	context.Bot().Edit(sentTransMsg, "⏸️ Транзакция отправлена в mempool 👆", keyboards.InlineKeyboardToHome)

	// ожидание инкрементации seqno в течение ~30 секунд
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
		go customErrors.MainErrorHandler(fmt.Errorf("processTransaction: %w", err), context)
		return
	}
	// если seqno так и не увеличился
	if seqnoAfterTrans.Seqno == seqnoBeforeTrans.Seqno {
		internalErr := customErrors.InternalError("wait process transaction in mempool: timeout")
		go customErrors.MainErrorHandler(fmt.Errorf("processTransaction: %w", internalErr), context)
		return
	}

	// изменение сообщения на "ожидание окончания транзакции"
	context.Bot().Edit(sentTransMsg, "🔄 Ожидание окончания транзакции... 👆", keyboards.InlineKeyboardToHome)



	// запуск функции ожидания следующей транзакции
	// transHash, err := tonApiSSE.SubscribeToNextTransaction(5*time.Minute)
	// if err != nil {
	// 	go customErrors.MainErrorHandler(fmt.Errorf("processTransaction: %w", err), context)
	// 	return
	// }


	// получение информации по хэшу отловленной транзакции



	// изменение сообщения на "транзакция завершена"
	context.Bot().Edit(sentTransMsg, "✅ Транзакция завершена! 👆", keyboards.InlineKeyboardToHome)



	// отправка нового сообщения с данными о закончившейся транзакции

}
