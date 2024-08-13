package wallet

import (
	ton_wallet "github.com/tonkeeper/tongo/wallet"
	"github.com/Danil-114195722/HamstersShaver/settings"
)


// получение реального кошелька по данным из JSON-конфига
func GetWallet() (ton_wallet.Wallet, error) {
	var realWallet ton_wallet.Wallet
	var err error

	// получаем доступ к кошельку (строка слов сид-фразы, TongoTonAPI)
	realWallet, err = ton_wallet.DefaultWalletFromSeed(settings.JsonWallet.SeedPhrase, settings.TongoTonAPI)
	
	if err != nil {
		settings.ErrorLog.Println("(tongo) Failed to get wallet:", err.Error())
		return realWallet, err
	}

	return realWallet, nil
}
