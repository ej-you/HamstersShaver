package account

import (
	"context"
	"errors"

	tonapi "github.com/tonkeeper/tonapi-go"

	"github.com/Danil-114195722/HamstersShaver/ton_api/tonapi/services"
	"github.com/Danil-114195722/HamstersShaver/settings"
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
		settings.ErrorLog.Println("Failed to get account:", err.Error())
		return account, err
	}

	// проверка того, что аккаунт активен
	if account.Status != "active" {
		accountIsNotActiveError := errors.New("account is not active")
		settings.ErrorLog.Println("Failed to interact with account:", accountIsNotActiveError.Error())
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
