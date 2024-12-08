package transactions

import (
	"context"
	"errors"
	"time"

	tonapi "github.com/tonkeeper/tonapi-go"

	"github.com/ej-you/HamstersShaver/rest_api/ton_api_rest/tonapi/services"
	"github.com/ej-you/HamstersShaver/rest_api/settings/constants"
	"github.com/ej-you/HamstersShaver/rest_api/settings"
)


// данные о транзакции по её хэшу
type TransactionInfo struct {
	// хэш транзакции
	Hash 		string `json:"hash" example:"ed79dafdda1a766dc6d7745e8dd03410adf7ba57bb6fccdb33ee5f5d8c3640f4"`
	// оставшееся кол-во TON после проведения транзакции
	EndBalance 	string `json:"endBalance" example:"2.689"`
	// была ли отклонена транзакция (не означает успех или неудачу транзакции)
	Bounce		bool `json:"bounce" example:"true"`
	// название операции транзакции
	OpName		string `json:"opName" example:"jetton_notify"`
}

// @Desctiption Данные о транзакции по её хэшу со статусом выполнения транзакции
type TransactionInfoWithStatusOK struct {
	// хэш транзакции
	Hash 		string `json:"hash" example:"ed79dafdda1a766dc6d7745e8dd03410adf7ba57bb6fccdb33ee5f5d8c3640f4"`
	// оставшееся кол-во TON после проведения транзакции
	EndBalance 	string `json:"endBalance" example:"2.689"`
	// была ли отклонена транзакция (не означает успех или неудачу транзакции)
	Bounce		bool `json:"bounce" example:"true"`
	// название операции транзакции
	OpName		string `json:"opName" example:"jetton_notify"`
	// действие с монетами в транзакции (покупка/продажа)
	Action 		string `json:"action" example:"buy"`
	// успех или неудача выполнения транзакции	
	StatusOK 	bool `json:"statusOK" example:"true"`
}


// получение информации о транзакции по её хэшу
func GetTransactionInfoByHash(ctx context.Context, hash string, timeout time.Duration) (TransactionInfo, error) {
	var transInfo TransactionInfo

	// получение клиента для tonapi-go
	tonapiClient, err := settings.GetTonClientTonapiWithTimeout("mainnet", timeout)
	if err != nil {
		return transInfo, err
	}

	params := tonapi.GetBlockchainTransactionParams{TransactionID: hash}
	// получение всей информации о транзакции
	rawTransInfo, err := tonapiClient.GetBlockchainTransaction(ctx, params)
	if err != nil {
		return transInfo, err
	}

	// выбор нужной информации
	transInfo = TransactionInfo{
		Hash: rawTransInfo.Hash,
		EndBalance: services.JettonBalanceFormat(rawTransInfo.EndBalance, constants.TonDecimals),
		Bounce: rawTransInfo.InMsg.Value.Bounce,
		OpName: rawTransInfo.InMsg.Value.DecodedOpName.Value,
	}

	return transInfo, nil
}


// получение информации о транзакции по её хэшу со статусом выполнения транзакции
// action может быть "buy" или "cell" (покупка и продажа монет соответственно)
func GetTransactionInfoWithStatusOKByHash(ctx context.Context, hash string, action string, timeout time.Duration) (TransactionInfoWithStatusOK, error) {
	var transInfoWithStatusOK TransactionInfoWithStatusOK
	
	if action != "buy" && action != "cell" {
		return transInfoWithStatusOK, errors.New("Invalid action parameter was given. Only \"buy\" and \"cell\" are accepted")
	}

	// получение структуры TransactionInfo
	transInfo, err := GetTransactionInfoByHash(ctx, hash, timeout)
	if err != nil {
		return transInfoWithStatusOK, err
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

	transInfoWithStatusOK = TransactionInfoWithStatusOK{
		Hash: transInfo.Hash,
		EndBalance: transInfo.EndBalance,
		Bounce: transInfo.Bounce,
		OpName: transInfo.OpName,
		Action: action,
		StatusOK: isOk,
	}

	return transInfoWithStatusOK, nil
}
