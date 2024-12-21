package account

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	tonapi "github.com/tonkeeper/tonapi-go"

	myToutilsgoServices "github.com/ej-you/HamstersShaver/rest_api/ton_api_rest/tonutils_go/services"

	"github.com/ej-you/HamstersShaver/rest_api/ton_api_rest/tonapi/services"
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

	// конвертация адреса монеты из HEX в base64
	jettonAddrBase64, err := myToutilsgoServices.ConvertAddrToBase64(rawJetton.Jetton.Address)
	if err != nil {
		return accountJettonInfo, err
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
