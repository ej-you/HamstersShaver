package account

import (
	"context"
	"errors"

	tonapi "github.com/tonkeeper/tonapi-go"

	"github.com/Danil-114195722/HamstersShaver/settings"
)


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
func GetBalanceTON(ctx context.Context) (float64, error) {
	var tonBalance float64

	// получение аккаунта
	account, err := GetAccount(ctx)
	if err != nil {
		return tonBalance, err
	}
	// преобразование баланса из нано-числа в число с точкой
	tonBalance = float64(account.Balance) / 1e9

	return tonBalance, nil
}
