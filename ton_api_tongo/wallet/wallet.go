package wallet

import (
	"context"

	ton_wallet "github.com/tonkeeper/tongo/wallet"
	"github.com/Danil-114195722/HamstersShaver/settings"
)


// получение реального кошелька по данным из JSON-конфига
func GetWallet() ton_wallet.Wallet {
	// получаем доступ к кошельку (строка слов сид-фразы, TongoTonAPI)
	realWallet, err := ton_wallet.DefaultWalletFromSeed(settings.JsonWallet.SeedPhrase, settings.TongoTonAPI)
	settings.DieIf(err)
	settings.InfoLog.Println("Got real wallet successfully")

	return realWallet
}

// получение баланса кошелька
func GetWalletBalance(ctx context.Context, wallet ton_wallet.Wallet) (uint64, error) {
	var balance uint64

	// получение текущего баланса кошелька
	balance, err := wallet.GetBalance(context.Background())
	if err != nil {
		settings.ErrorLog.Println("Failed to get balance: ", err.Error())
		return balance, err
	}
	settings.InfoLog.Println("Got wallet balance successfully")
	return balance, nil
}
