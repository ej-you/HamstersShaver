package wallet

import (
	"fmt"
	"context"
	"strings"
	"strconv"

	"github.com/xssnick/tonutils-go/tlb"
	ton_wallet "github.com/xssnick/tonutils-go/ton/wallet"

	"github.com/Danil-114195722/HamstersShaver/settings"
)


// получение реального кошелька по данным из JSON-конфига
func GetWallet() *ton_wallet.Wallet {
	// делим сид-фразу на список из слов
	seedList := strings.Split(settings.JsonWallet.SeedPhrase, " ")

	// получаем доступ к кошельку (TonAPI, список слов сид-фразы, версия конфига)
	realWallet, err := ton_wallet.FromSeed(settings.TonAPI, seedList, ton_wallet.V4R2)
	settings.DieIf(err)

	return realWallet
}

// получение баланса кошелька
func GetWalletBalance(ctx context.Context, wallet *ton_wallet.Wallet) (tlb.Coins, error) {
	// получение главного блока
	block, err := settings.TonAPI.CurrentMasterchainInfo(ctx)
	if err != nil {
		settings.ErrorLog.Println("Failed to get masterchain info: ", err.Error())
		return tlb.Coins{}, err
	}

	// получение текущего баланса кошелька
	balance, err := wallet.GetBalance(ctx, block)
	if err != nil {
		settings.ErrorLog.Println("Failed to get balance: ", err.Error())
		return tlb.Coins{}, err
	}
	return balance, nil
}

// получение баланса кошелька в виде числа float64
func GetWalletBalanceFloat64(ctx context.Context, wallet *ton_wallet.Wallet) (float64, error) {
	var balanceFloat64 float64

	// получение текущего баланса кошелька в виде структуры tlb.Coins
	balanceInCoins, err := GetWalletBalance(ctx, wallet)
	if err != nil {
		return balanceFloat64, err
	}

	fmt.Println("balance in Coins:", balanceInCoins.String())

	// НЕ РАБОТАЕТ
	balanceFloat64, err = strconv.ParseFloat(fmt.Sprintf("%f", balanceInCoins.String()), 64)
	if err != nil {
		settings.ErrorLog.Println("Failed to convert balance in Coins to float64: ", err.Error())
		return balanceFloat64, err
	}

	return balanceFloat64, nil
}
