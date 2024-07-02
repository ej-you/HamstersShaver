package account

import (
	"context"

	"github.com/tonkeeper/tongo/tlb"
	"github.com/tonkeeper/tongo/ton"

	"github.com/Danil-114195722/HamstersShaver/settings"
)


// получение аккаунта по данным из JSON-конфига
func GetAccountState(ctx context.Context) (tlb.ShardAccount, error) {
	var accountState tlb.ShardAccount

	// парсим адрес аккаунта
	accountID := ton.MustParseAccountID(settings.JsonWallet.Hash)
	settings.InfoLog.Println("Account ID was parsed successfully")

	// получение аккаунта
	accountState, err := settings.TongoTonAPI.GetAccountState(ctx, accountID)
	if err != nil {
		settings.ErrorLog.Println("Failed to get account state:", err.Error())
		return accountState, err
	}

	// // проверка того, что аккаунт активен
	// if !acc.IsActive {
	// 	accountIsNotActiveError := errors.New("account is not active")
	// 	settings.DieIf(accountIsNotActiveError)
	// }

	settings.InfoLog.Println("Got account state successfully")
	return accountState, nil
}
