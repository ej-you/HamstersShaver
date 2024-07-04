package account

import (
	"fmt"
	"context"
	"errors"

	tonapi "github.com/tonkeeper/tonapi-go"

	"github.com/Danil-114195722/HamstersShaver/settings"
)


type TonJetton struct {
	Balance int64
	Decimals int
	BeautyBalance string
}

// получение аккаунта по данным из JSON-конфига
func GetAccount(ctx context.Context) (*tonapi.Account, error) {
	var account *tonapi.Account

	// конфиг API для получения аккаунта
	accountParams := tonapi.GetAccountParams{AccountID: settings.JsonWallet.Hash}

	// получение аккаунта по его адресу
	account, err := settings.TonapiTonAPI.GetAccount(context.Background(), accountParams)
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
func GetBalanceTON(ctx context.Context) (TonJetton, error) {
	var tonBalance string
	var tonJetton TonJetton

	// получение аккаунта
	account, err := GetAccount(ctx)
	if err != nil {
		return tonJetton, err
	}
	
	// преобразование баланса из нано-числа в число с точкой с округлением до 2 знаков (в виде строки)
	tonBalance = fmt.Sprintf("%.2f", float64(account.Balance) / 1e9)

	// создание экзземпляра структуры TonJetton
	tonJetton = TonJetton{Balance: account.Balance, Decimals: 9, BeautyBalance: tonBalance}

	return tonJetton, nil
}
