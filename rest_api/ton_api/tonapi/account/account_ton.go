package account

import (
	"fmt"
	"context"

	tonapi "github.com/tonkeeper/tonapi-go"

	coreErrors "github.com/ej-you/HamstersShaver/rest_api/core/errors"
	"github.com/ej-you/HamstersShaver/rest_api/ton_api/tonapi/services"
	"github.com/ej-you/HamstersShaver/rest_api/settings"
)


// описывает TON монету, имеющуюся у аккаунта
type TonJetton struct {
	Balance 		int64 `json:"balance" example:"1955532900" description:"баланс монеты на аккаунте"`
	Decimals 		int `json:"decimals" example:"9" description:"decimals монеты"`
	BeautyBalance 	string `json:"beautyBalance" example:"1.956" description:"округлённый баланс"`
}


// получение аккаунта по данным из JSON-конфига
func GetAccount(ctx context.Context, tonapiClient *tonapi.Client) (*tonapi.Account, error) {
	var account *tonapi.Account

	// конфиг API для получения аккаунта
	accountParams := tonapi.GetAccountParams{AccountID: settings.GetJsonWallet().Hash}

	// получение аккаунта по его адресу
	account, err := tonapiClient.GetAccount(ctx, accountParams)
	if err != nil {
		// ошибка таймаута
		if coreErrors.IsTimeout(err) {
			return account, fmt.Errorf("get account using tonapi: %w", coreErrors.TimeoutError)
		}
		// неизвестная ошибка
		return account, fmt.Errorf("get account using tonapi: %v: %w", err, coreErrors.TonApiError)
	}

	// проверка того, что аккаунт активен
	if account.Status != "active" {
		return account, fmt.Errorf("get account using tonapi: account is not active: %w", coreErrors.TonApiError)
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
		return tonJetton, fmt.Errorf("get account ton balance: %w", err)
	}

	// преобразование баланса в строку с точкой
	tonBalance = services.BeautyJettonAmountFromInt64(account.Balance, 9)

	// создание экзземпляра структуры TonJetton
	tonJetton = TonJetton{Balance: account.Balance, Decimals: 9, BeautyBalance: tonBalance}

	return tonJetton, nil
}
