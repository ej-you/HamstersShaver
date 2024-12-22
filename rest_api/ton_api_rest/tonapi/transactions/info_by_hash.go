package transactions

import (
	"context"
	"fmt"
	"time"

	tonapi "github.com/tonkeeper/tonapi-go"

	"github.com/ej-you/HamstersShaver/rest_api/ton_api_rest/tonapi/services"

	coreErrors "github.com/ej-you/HamstersShaver/rest_api/core/errors"
	"github.com/ej-you/HamstersShaver/rest_api/settings/constants"
	"github.com/ej-you/HamstersShaver/rest_api/settings"
)


// данные о транзакции по её хэшу
type TransactionInfo struct {
	Hash 		string `json:"hash" example:"ed79dafdda1a766dc6d7745e8dd03410adf7ba57bb6fccdb33ee5f5d8c3640f4"`
	EndBalance 	string `json:"endBalance" example:"2.689"`
	Bounce		bool `json:"bounce" example:"true"`
	OpName		string `json:"opName" example:"jetton_notify"`
}

// данные о транзакции по её хэшу со статусом выполнения транзакции
type TransactionInfoWithStatusOK struct {
	Hash 		string `json:"hash" example:"ed79dafdda1a766dc6d7745e8dd03410adf7ba57bb6fccdb33ee5f5d8c3640f4" description:"хэш транзакции"`
	EndBalance 	string `json:"endBalance" example:"2.689" description:"оставшееся кол-во TON после проведения транзакции"`
	Bounce		bool `json:"bounce" example:"true" description:"была ли отклонена операция (не означает успех или неудачу транзакции)"`
	OpName		string `json:"opName" example:"jetton_notify" description:"название операции транзакции"`
	Action 		string `json:"action" example:"buy" description:"действие с монетами в транзакции (покупка/продажа)"`
	StatusOK 	bool `json:"statusOK" example:"true" description:"успех или неудача выполнения транзакции"`
}


// получение информации о транзакции по её хэшу
func GetTransactionInfoByHash(ctx context.Context, hash string, timeout time.Duration) (TransactionInfo, error) {
	var transInfo TransactionInfo

	// получение клиента для tonapi-go
	tonapiClient, err := settings.GetTonClientTonapiWithTimeout("mainnet", timeout)
	if err != nil {
		return transInfo, fmt.Errorf("get transaction info using tonapi: %w", err)
	}

	params := tonapi.GetBlockchainTransactionParams{TransactionID: hash}
	// получение всей информации о транзакции
	rawTransInfo, err := tonapiClient.GetBlockchainTransaction(ctx, params)
	if err != nil {
		apiErr := coreErrors.New(
			fmt.Errorf("get transaction info using tonapi: %w", err),
			"failed to get transaction info",
			"ton_api",
			500,
		)
		return transInfo, apiErr
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
		apiErr := coreErrors.New(
			fmt.Errorf("get transaction info using tonapi: invalid action parameter was given: %s", action),
			"invalid action parameter",
			"rest_api",
			400,
		)
		return transInfoWithStatusOK, apiErr
	}

	// получение структуры TransactionInfo
	transInfo, err := GetTransactionInfoByHash(ctx, hash, timeout)
	if err != nil {
		return transInfoWithStatusOK, fmt.Errorf("get transaction info with status using tonapi: %w", err)
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
