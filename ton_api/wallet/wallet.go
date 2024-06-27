package wallet

import (
	"context"
	"strings"

	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/ton/wallet"
	
	"github.com/Danil-114195722/HamstersShaver/ton_api"
	"github.com/Danil-114195722/HamstersShaver/settings"
)


// получение данных о кошельке из JSON-конфига
var JsonWallet settings.JsonWallet = settings.GetJsonWallet()

// получение реального кошелька по данным из JSON-конфига
func GetWallet() *wallet.Wallet {
	// делим сид-фразу на список из слов
	seedList := strings.Split(JsonWallet.SeedPhrase, " ")

	// получаем доступ к кошельку (API, список слов сид-фразы, версия конфига)
	realWallet, err := wallet.FromSeed(ton_api.API, seedList, wallet.V4R2)
	settings.DieIf(err)

	return realWallet
}

// получение баланса кошелька
func GetWalletBalance(wallet *wallet.Wallet) (tlb.Coins, error) {
	// получение главного блока
	block, err := ton_api.API.CurrentMasterchainInfo(context.Background())
	if err != nil {
		settings.ErrorLog.Println("Failed to get masterchain info: ", err.Error())
		return tlb.Coins{}, err
	}

	// получение текущего баланса кошелька
	balance, err := wallet.GetBalance(context.Background(), block)
	if err != nil {
		settings.ErrorLog.Println("Failed to get balance: ", err.Error())
		return tlb.Coins{}, err
	}

	return balance, nil
}
