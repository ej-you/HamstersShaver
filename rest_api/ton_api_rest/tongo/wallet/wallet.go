package wallet

import (
	"fmt"

	tongo "github.com/tonkeeper/tongo/liteapi"
	tongoWallet "github.com/tonkeeper/tongo/wallet"

	coreErrors "github.com/ej-you/HamstersShaver/rest_api/core/errors"
	"github.com/ej-you/HamstersShaver/rest_api/settings"
)


// получение реального кошелька по данным из JSON-конфига
func GetWallet(tongoClient *tongo.Client) (tongoWallet.Wallet, error) {
	var realWallet tongoWallet.Wallet
	var err error

	// получаем доступ к кошельку (строка слов сид-фразы, tongoClient)
	realWallet, err = tongoWallet.DefaultWalletFromSeed(settings.JsonWallet.SeedPhrase, tongoClient)
	
	if err != nil {
		apiErr := coreErrors.New(
			fmt.Errorf("get wallet using tongo: %w", err),
			"get wallet",
			"ton_api",
			500,
		)
		return realWallet, apiErr
	}

	return realWallet, nil
}
