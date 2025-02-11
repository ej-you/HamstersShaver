package transactions

import (
	"context"
	"fmt"

	tonapi "github.com/tonkeeper/tonapi-go"

	"github.com/ej-you/HamstersShaver/rest_api/ton_api/tonapi/services"

	coreErrors "github.com/ej-you/HamstersShaver/rest_api/core/errors"
	"github.com/ej-you/HamstersShaver/rest_api/settings/constants"
	"github.com/ej-you/HamstersShaver/rest_api/settings"
)


// данные о транзакции по её хэшу со статусом выполнения транзакции
type TransactionInfo struct {
	Hash 		string `json:"hash" example:"4f8ff3378e1d4cc80488750fda3bcc6b730b71b69429d9c44a775b377bdc66a4" description:"хэш транзакции"`
	LastTxHash 	string `json:"lastTxHash" example:"a8ec992c341230a885f9adfe6598eb307660c306a5f00cf0d302c72e7d966389" description:"хэш последней операции транзакции"`
	EndTime 	int64 `json:"endTime" example:"1735413815" description:"время окончания транзакции в UNIX-формате"`
	EndBalance 	string `json:"endBalance" example:"2.689" description:"оставшееся кол-во TON после проведения транзакции"`
	Bounce		bool `json:"bounce" example:"true" description:"была ли отклонена операция (не означает успех или неудачу транзакции)"`
	OpName		string `json:"opName" example:"jetton_notify" description:"название операции транзакции"`
	Action 		string `json:"action" example:"buy" description:"действие с монетами в транзакции (покупка/продажа)"`
	StatusOK 	bool `json:"statusOK" example:"true" description:"успех или неудача выполнения транзакции"`
}


// разворачивание списка операций транзакции до последней операции
func unwrapTrace(trace *tonapi.Trace) *tonapi.Trace {
	if trace.Children == nil {
		return trace
	}
	return unwrapTrace(&trace.Children[0])
}


// получение информации о завершённой транзакции по хэшу её первой операции
func getTransactionInfo(ctx context.Context, txHash string) (TransactionInfo, error) {
	var transInfo TransactionInfo

	// получение клиента для tonapi-go
	tonapiClient, err := settings.GetTonClientTonapiWithTimeout("mainnet", constants.TonapiClientTimeout)
	if err != nil {
		return transInfo, fmt.Errorf("get transaction info using tonapi: %w", err)
	}

	// получаем список со всеми операциями транзакции
	fullTrace, err := tonapiClient.GetTrace(ctx, tonapi.GetTraceParams{TraceID: txHash})
	if err != nil {
		// ошибка таймаута
		if coreErrors.IsTimeout(err) {
			return transInfo, fmt.Errorf("get transaction info using tonapi: get transaction trace: %w", coreErrors.TimeoutError)
		}
		// неизвестная ошибка
		return transInfo, fmt.Errorf("get transaction info using tonapi: get transaction trace: %v: %w", err, coreErrors.TonApiError)
	}
	// получаем информацию о последней операции
	lastTransInfo := unwrapTrace(fullTrace).Transaction

	// выбор нужной информации
	transInfo = TransactionInfo{
		Hash: txHash,
		LastTxHash: lastTransInfo.Hash,
		EndTime: lastTransInfo.Utime,
		EndBalance: services.BeautyJettonAmountFromInt64(lastTransInfo.EndBalance, constants.TonDecimals),
		Bounce: lastTransInfo.InMsg.Value.Bounce,
		OpName: lastTransInfo.InMsg.Value.DecodedOpName.Value,
	}

	return transInfo, nil
}


// получение информации о завершённой транзакции по хэшу её первой операции со статусом выполнения транзакции
// action может быть "buy" или "cell" (покупка и продажа монет соответственно)
func GetTransactionInfoWithStatusOK(ctx context.Context, txHash, action string) (TransactionInfo, error) {
	var transInfo TransactionInfo

	if action != "buy" && action != "cell" {
		return transInfo, fmt.Errorf("get transaction info with status: invalid action parameter: %s: %w", action, coreErrors.RestApiError)
	}

	// получение структуры TransactionInfo
	transInfo, err := getTransactionInfo(ctx, txHash)
	if err != nil {
		return transInfo, fmt.Errorf("get transaction info with status: %w", err)
	}

	// была подмечена закономерность, что при успешной транзакции продажи монет её OpName == "jetton_notify" и отскок Bounce == true
	// а при успешной транзакции покупки монет её OpName == "excess" и отскок Bounce == false
	// при неуспешных же транзакциях всё в точности наоборот
	var isOk bool
	if action == "buy" && !transInfo.Bounce || action == "cell" && transInfo.Bounce {
		isOk = true
	} else if action == "buy" && transInfo.Bounce || action == "cell" && !transInfo.Bounce {
		isOk = false
	}

	transInfo.Action = action
	transInfo.StatusOK = isOk

	return transInfo, nil
}
