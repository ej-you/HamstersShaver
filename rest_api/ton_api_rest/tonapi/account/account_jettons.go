package account

import (
	"fmt"
	"context"
	"strconv"

	tonapi "github.com/tonkeeper/tonapi-go"

	myToutilsgoServices "github.com/ej-you/HamstersShaver/rest_api/ton_api_rest/tonutils_go/services"

	coreErrors "github.com/ej-you/HamstersShaver/rest_api/core/errors"
	"github.com/ej-you/HamstersShaver/rest_api/ton_api_rest/tonapi/services"
	"github.com/ej-you/HamstersShaver/rest_api/settings"
)


// получение монет и их количества у аккаунта по данным из JSON-конфига
func GetBalanceJettons(ctx context.Context, tonapiClient *tonapi.Client) ([]AccountJetton, error) {
	var rawJettons *tonapi.JettonsBalances

	// переменные для перебора монет в цикле
	var loopAccountJetton AccountJetton  // структура описана в файле account_jetton.go
	var loopJettonSymbol string
	var intJettonBalance int64
	var jettonDecimals int
	var beautyLoopJettonBalance string
	var jettonAddrBase64 string
	var loopErr error
	// переменная для сохранения информации о монетах в виде списка структур AccountJetton
	accountJettonsList := []AccountJetton{}

	// конфиг API для получения монет аккаунта
	accountJettonsParams := tonapi.GetAccountJettonsBalancesParams{AccountID: settings.JsonWallet.Hash}

	// получение всех монет аккаунта
	rawJettons, err := tonapiClient.GetAccountJettonsBalances(ctx, accountJettonsParams)
	if err != nil {
		apiErr := coreErrors.New(
			fmt.Errorf("get account jettons using tonapi: %w", err),
			"failed to get account jettons",
			"ton_api",
			500,
		)
		apiErr.CheckTimeout()
		return accountJettonsList, apiErr
	}

	// перебор всех найденных монет аккаунта (сохраняется вся история монет, которые были на кошельке)
	for _, rawJetton := range rawJettons.Balances {
		// если в данный момент баланс монеты "0"
		if rawJetton.Balance == "0" {
			continue
		}

		// краткое название монеты (полное название - rawJetton.Jetton.Name)
		loopJettonSymbol = rawJetton.Jetton.Symbol

		// перевод баланса монеты из строкового целого представления в int64
		intJettonBalance, loopErr = strconv.ParseInt(rawJetton.Balance, 10, 64)
		if loopErr != nil {
			settings.ErrorLog.Printf("get account jettons using tonapi: parse int64 from string jetton %q balance: %v", loopJettonSymbol, loopErr.Error())
			continue
		}
		// на это нужно делить, чтобы получить число с точкой
		jettonDecimals = rawJetton.Jetton.Decimals
		// преобразование баланса в строку с точкой
		beautyLoopJettonBalance = services.BeautyJettonAmountFromInt64(intJettonBalance, jettonDecimals)

		// конвертация адреса монеты из HEX в base64
		jettonAddrBase64, loopErr = myToutilsgoServices.ConvertAddrToBase64(rawJetton.Jetton.Address)
		if err != nil {
			settings.ErrorLog.Printf("get account jettons using tonapi: convert raw jetton addr %q to base64: %v", rawJetton.Jetton.Address, loopErr.Error())
			continue
		}

		// создание структуры для новой монеты и добавление её в список к остальным
		loopAccountJetton = AccountJetton{
			Symbol: loopJettonSymbol,
			Balance: intJettonBalance,
			Decimals: jettonDecimals,
			BeautyBalance: beautyLoopJettonBalance,
			// мастер-адрес монеты
			MasterAddress: jettonAddrBase64,
		}
		accountJettonsList = append(accountJettonsList, loopAccountJetton)
	}

	// если ни одна монета не была найдена на счету аккаунта
	if len(accountJettonsList) == 0 {
		apiErr := coreErrors.New(
			fmt.Errorf("get account jettons using tonapi: empty account jettons list: no one account jetton was gotten"),
			"no one account jetton was gotten",
			"rest_api",
			500,
		)
		return accountJettonsList, apiErr
	}

	return accountJettonsList, nil
}
