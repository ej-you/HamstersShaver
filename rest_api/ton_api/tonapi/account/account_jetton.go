package account

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	tonapi "github.com/tonkeeper/tonapi-go"

	myToutilsgoServices "github.com/ej-you/HamstersShaver/rest_api/ton_api/tonutils_go/services"

	coreErrors "github.com/ej-you/HamstersShaver/rest_api/core/errors"
	"github.com/ej-you/HamstersShaver/rest_api/ton_api/tonapi/services"
	"github.com/ej-you/HamstersShaver/rest_api/settings"
)


// описывает монету (кроме TON), имеющуюся у аккаунта
type AccountJetton struct {
	Symbol 			string `json:"symbol" example:"GRAM" description:"символ монеты"`
	Balance 		int64 `json:"balance" example:"326166742480" description:"баланс монеты на аккаунте"`
	Decimals 		int `json:"decimals" example:"9" description:"decimals монеты"`
	BeautyBalance 	string `json:"beautyBalance" example:"326.167" description:"округлённый баланс"`
	MasterAddress 	string `json:"masterAddress" example:"EQC47093oX5Xhb0xuk2lCr2RhS8rj-vul61u4W2UH5ORmG_O" description:"мастер-адрес монеты (jetton_master)"`
}


// получение инфы о монете аккаунта
func GetAccountJetton(ctx context.Context, tonapiClient *tonapi.Client, jettonCA string) (AccountJetton, error) {
	var accountJettonInfo AccountJetton
	var rawJetton *tonapi.JettonBalance
	var apiErr coreErrors.APIError

	// конфиг API для получения инфы о монете аккаунта
	accountJettonParams := tonapi.GetAccountJettonBalanceParams{
		AccountID: settings.GetJsonWallet().Hash,
		JettonID: jettonCA,
		Currencies: []string{},
	}

	// получение инфы о монете аккаунта
	rawJetton, err := tonapiClient.GetAccountJettonBalance(ctx, accountJettonParams)
	if err != nil {
		// если такой монеты нет у данного аккаунта
		if strings.HasPrefix(err.Error(), "decode response: error: code 404: {Error:account") {
			apiErr = coreErrors.New(
				fmt.Errorf("get account jetton using tonapi: account has not given jetton: %w", err),
				"account has not given jetton",
				"ton_api",
				404,
			)
			return accountJettonInfo, apiErr
		// если был дан неверный адрес
		} else if strings.HasPrefix(err.Error(), "decode response: error: code 4") {
			apiErr = coreErrors.New(
				fmt.Errorf("get account jetton using tonapi: invalid jetton address was given: %w", err),
				"invalid jetton address was given",
				"ton_api",
				400,
			)
			return accountJettonInfo, apiErr
		}
		// неизвестная ошибка
		apiErr = coreErrors.New(
			fmt.Errorf("get account jetton using tonapi: %w", err),
			"failed to get account jetton",
			"ton_api",
			500,
		)
		apiErr.CheckTimeout()
		return accountJettonInfo, apiErr
	}
	
	jettonDecimals := rawJetton.Jetton.Decimals
	// перевод баланса монеты из строкового целого представления в int64
	intJettonBalance, err := strconv.ParseInt(rawJetton.Balance, 10, 64)
	if err != nil {
		apiErr = coreErrors.New(
			fmt.Errorf("get account jetton using tonapi: parse int64 from string jetton balance: %w", err),
			"failed to get account jetton",
			"rest_api",
			500,
		)
		return accountJettonInfo, apiErr
	}
	// преобразование баланса в строку с точкой
	beautyJettonBalance := services.BeautyJettonAmountFromInt64(intJettonBalance, jettonDecimals)

	// конвертация адреса монеты из HEX в base64
	jettonAddrBase64, err := myToutilsgoServices.ConvertAddrToBase64(rawJetton.Jetton.Address)
	if err != nil {
		return accountJettonInfo, fmt.Errorf("failed to get account jetton: %w", err)
	}

	// создание структуры для новой монеты и добавление её в список к остальным
	accountJettonInfo = AccountJetton{
		Symbol: rawJetton.Jetton.Symbol,
		Balance: intJettonBalance,
		Decimals: jettonDecimals,
		BeautyBalance: beautyJettonBalance,
		// мастер-адрес монеты
		MasterAddress: jettonAddrBase64,
	}

	return accountJettonInfo, nil
}
