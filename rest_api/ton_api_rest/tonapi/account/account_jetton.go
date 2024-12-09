package account

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	tonapi "github.com/tonkeeper/tonapi-go"

	"github.com/ej-you/HamstersShaver/rest_api/ton_api_rest/tonapi/services"
	"github.com/ej-you/HamstersShaver/rest_api/settings"
)


// @Description Описывает монету (кроме TON), имеющуюся у аккаунта
type AccountJetton struct {
	// символ монеты
	Symbol 			string `json:"symbol" example:"GRAM"`
	// баланс монеты на аккаунте
	Balance 		int64 `json:"balance" example:"326166742480"`
	// decimals монеты
	Decimals 		int `json:"decimals" example:"9"`
	// округлённый баланс
	BeautyBalance 	string `json:"beautyBalance" example:"326.167"`
	// мастер-адрес монеты (jetton_master)
	MasterAddress 	string `json:"masterAddress" example:"0:b8ef4f77a17e5785bd31ba4da50abd91852f2b8febee97ad6ee16d941f939198"`
}


// получение инфы о монете аккаунта
func GetAccountJetton(ctx context.Context, tonapiClient *tonapi.Client, jettonCA string) (AccountJetton, error) {
	var accountJettonInfo AccountJetton
	var rawJetton *tonapi.JettonBalance

	// конфиг API для получения инфы о монете аккаунта
	accountJettonParams := tonapi.GetAccountJettonBalanceParams{
		AccountID: settings.JsonWallet.Hash,
		JettonID: jettonCA,
		Currencies: []string{},
	}

	// получение инфы о монете аккаунта
	rawJetton, err := tonapiClient.GetAccountJettonBalance(ctx, accountJettonParams)
	if err != nil {
		getAccountJettonError := errors.New(fmt.Sprintf("Failed to get account jetton info: %s", err.Error()))
		return accountJettonInfo, getAccountJettonError
	}
	
	jettonDecimals := rawJetton.Jetton.Decimals
	// перевод баланса монеты из строкового целого представления в int64
	intJettonBalance, err := strconv.ParseInt(rawJetton.Balance, 10, 64)
	if err != nil {
		parseIntError := errors.New(fmt.Sprintf("Failed to parse int64 from string jetton balance: %s", err.Error()))
		return accountJettonInfo, parseIntError
	}
	// преобразование баланса в строку с точкой
	beautyJettonBalance := services.JettonBalanceFormat(intJettonBalance, jettonDecimals)

	// создание структуры для новой монеты и добавление её в список к остальным
	accountJettonInfo = AccountJetton{
		Symbol: rawJetton.Jetton.Symbol,
		Balance: intJettonBalance,
		Decimals: jettonDecimals,
		BeautyBalance: beautyJettonBalance,
		// мастер-адрес монеты
		MasterAddress: rawJetton.Jetton.Address,
	}

	return accountJettonInfo, nil
}
