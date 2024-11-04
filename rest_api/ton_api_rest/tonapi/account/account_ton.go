package account

import (
	"fmt"
	"context"
	"errors"

	tonapi "github.com/tonkeeper/tonapi-go"

	"github.com/ej-you/HamstersShaver/rest_api/ton_api_rest/tonapi/services"
	"github.com/ej-you/HamstersShaver/rest_api/settings"
)


// описывает TON монету, имеющуюся у аккаунта
type TonJetton struct {
	Balance 		int64 `json:"balance"`
	Decimals 		int `json:"decimals"`
	BeautyBalance 	string `json:"beautyBalance"`
}


// получение аккаунта по данным из JSON-конфига
func GetAccount(ctx context.Context, tonapiClient *tonapi.Client) (*tonapi.Account, error) {
	var account *tonapi.Account

	// конфиг API для получения аккаунта
	accountParams := tonapi.GetAccountParams{AccountID: settings.JsonWallet.Hash}

	// получение аккаунта по его адресу
	account, err := tonapiClient.GetAccount(ctx, accountParams)
	if err != nil {
		getAccountError := errors.New(fmt.Sprintf("Failed to get account: %s", err.Error()))
		settings.ErrorLog.Println(getAccountError.Error())
		return account, getAccountError
	}

	// проверка того, что аккаунт активен
	if account.Status != "active" {
		accountIsNotActiveError := errors.New("Failed to interact with account: account is not active")
		settings.ErrorLog.Println(accountIsNotActiveError.Error())
		return account, accountIsNotActiveError
	}

	return account, nil
}


// получение баланса аккаунта в тонах
func GetBalanceTON(ctx context.Context, tonapiClient *tonapi.Client) (TonJetton, error) {
	var tonBalance string
	var tonJetton TonJetton

	// получение аккаунта
	account, err := GetAccount(ctx, tonapiClient)
	if err != nil {
		return tonJetton, err
	}

	// преобразование баланса в строку с точкой
	tonBalance = services.JettonBalanceFormat(account.Balance, 9)

	// создание экзземпляра структуры TonJetton
	tonJetton = TonJetton{Balance: account.Balance, Decimals: 9, BeautyBalance: tonBalance}

	return tonJetton, nil
}
