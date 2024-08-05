package jettons

import (
	"context"
	"errors"
	"strconv"

	tonapi "github.com/tonkeeper/tonapi-go"

	"github.com/Danil-114195722/HamstersShaver/ton_api_tonapi/services"
	"github.com/Danil-114195722/HamstersShaver/settings"
)

// описывает монету (кроме TON), имеющуюся у аккаунта
type AccountJetton struct {
	Symbol string
	Balance int64
	Decimals int
	BeautyBalance string
	// мастер-адрес монеты
	MasterAddress string
}


// получение монет и их количества у аккаунта по данным из JSON-конфига
func GetBalanceJettons(ctx context.Context) ([]AccountJetton, error) {
	var rawJettons *tonapi.JettonsBalances

	// переменные для перебора монет в цикле
	var loopAccountJetton AccountJetton
	var loopJettonSymbol string
	var intJettonBalance int64
	var jettonDecimals int
	var beautyLoopJettonBalance string
	var loopErr error
	// переменная для сохранения информации о монетах в виде списка структур AccountJetton
	accountJettonsList := []AccountJetton{}

	// конфиг API для получения монет аккаунта
	accountJettonsParams := tonapi.GetAccountJettonsBalancesParams{AccountID: settings.JsonWallet.Hash}

	// получение всех монет аккаунта
	rawJettons, err := settings.TonapiTonAPI.GetAccountJettonsBalances(ctx, accountJettonsParams)
	if err != nil {
		settings.ErrorLog.Println("Failed to get account jettons:", err.Error())
		return accountJettonsList, err
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
			settings.ErrorLog.Printf("Failed to parse int64 from string jetton %q balance: %v", loopJettonSymbol, loopErr.Error())
			continue
		}
		// на это нужно делить, чтобы получить число с точкой
		jettonDecimals = rawJetton.Jetton.Decimals
		// преобразование баланса в строку с точкой
		beautyLoopJettonBalance = services.JettonBalanceFormat(intJettonBalance, jettonDecimals)

		// создание структуры для новой монеты и добавление её в список к остальным
		loopAccountJetton = AccountJetton{
			Symbol: loopJettonSymbol,
			Balance: intJettonBalance,
			Decimals: jettonDecimals,
			BeautyBalance: beautyLoopJettonBalance,
			// мастер-адрес монеты
			MasterAddress: rawJetton.Jetton.Address,
		}
		accountJettonsList = append(accountJettonsList, loopAccountJetton)
	}

	// если ни одна монета не была найдена на счету аккаунта
	if len(accountJettonsList) == 0 {
		emptyJetonsListError := errors.New("no one account jetton was gotten")
		settings.ErrorLog.Println("Empty account jettons list:", emptyJetonsListError.Error())
		return accountJettonsList, emptyJetonsListError
	}

	return accountJettonsList, nil
}
