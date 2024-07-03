package jettons

import (
	"context"
	"errors"
	"math"
	"strconv"

	tonapi "github.com/tonkeeper/tonapi-go"

	"github.com/Danil-114195722/HamstersShaver/settings"
)


type AccountJetton struct {
	Symbol string
	Balance float64
}


// получение монет и их количества у аккаунта по данным из JSON-конфига
func GetBalanceJettons(ctx context.Context) ([]AccountJetton, error) {
	var rawJettons *tonapi.JettonsBalances

	// переменные для перебора монет в цикле
	var loopAccountJetton AccountJetton
	var loopJettonSymbol string
	var intLoopJettonBalance int
	var floatDevider int
	var floatLoopJettonBalance float64
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

	for _, rawJetton := range rawJettons.Balances {
		// краткое название монеты (полное название - rawJetton.Jetton.Name)
		loopJettonSymbol = rawJetton.Jetton.Symbol

		// перевод баланса монеты из строкового целого представления в int
		intLoopJettonBalance, loopErr = strconv.Atoi(rawJetton.Balance)
		if loopErr != nil {
			settings.ErrorLog.Printf("Failed to parse float64 from string jetton %q balance: %v", loopJettonSymbol, loopErr.Error())
			continue
		}
		// на это нужно делить, чтобы получить число с точкой
		floatDevider = rawJetton.Jetton.Decimals
		// преобразование баланса во float64
		floatLoopJettonBalance = float64(intLoopJettonBalance) / math.Pow10(floatDevider)

		// создание структуры для новой монеты и добавление её в список к остальным
		loopAccountJetton = AccountJetton{Symbol: loopJettonSymbol, Balance: floatLoopJettonBalance}
		accountJettonsList = append(accountJettonsList, loopAccountJetton)
	}

	if len(accountJettonsList) == 0 {
		emptyJetonsListError := errors.New("no one account jetton was gotten")
		settings.ErrorLog.Println("Empty account jettons list:", emptyJetonsListError.Error())
		return accountJettonsList, emptyJetonsListError
	}

	return accountJettonsList, nil
}
