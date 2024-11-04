package wallet

import (
	"fmt"
	"errors"

	tongo "github.com/tonkeeper/tongo/liteapi"
	tongoWallet "github.com/tonkeeper/tongo/wallet"

	"github.com/Danil-114195722/HamstersShaver/rest_api/settings"
)


// получение реального кошелька по данным из JSON-конфига
func GetWallet(tongoClient *tongo.Client) (tongoWallet.Wallet, error) {
	var realWallet tongoWallet.Wallet
	var err error

	// получаем доступ к кошельку (строка слов сид-фразы, tongoClient)
	realWallet, err = tongoWallet.DefaultWalletFromSeed(settings.JsonWallet.SeedPhrase, tongoClient)
	
	if err != nil {
		getWalletError := errors.New(fmt.Sprintf("(tongo) Failed to get wallet: %s", err.Error()))
		settings.ErrorLog.Println(getWalletError.Error())
		return realWallet, getWalletError
	}

	return realWallet, nil
}
