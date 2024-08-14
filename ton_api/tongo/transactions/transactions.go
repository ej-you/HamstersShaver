package transactions

import (
	"math/big"
	"fmt"
	"context"

	tongoStonfi "github.com/tonkeeper/tongo/contract/stonfi"
	tongoTlb "github.com/tonkeeper/tongo/tlb"
	tongoTon "github.com/tonkeeper/tongo/ton"

	myTonapiJettons "github.com/Danil-114195722/HamstersShaver/ton_api/tonapi/jettons"
	
	myTongoWallet "github.com/Danil-114195722/HamstersShaver/ton_api/tongo/wallet"
	myTongoJettons "github.com/Danil-114195722/HamstersShaver/ton_api/tongo/jettons"
	
	"github.com/Danil-114195722/HamstersShaver/settings/constants"
	"github.com/Danil-114195722/HamstersShaver/settings"
)


// продажа монет
// TODO: 10 попыток до успеха (ошибка "error code: 651 message: cannot load block")
func CellJetton(ctx context.Context, jettonCA string, jetton myTonapiJettons.AccountJetton, amount float64, slippage int) error {
	// получение данных о кошельке через tongo
	realWallet, err := myTongoWallet.GetWallet()
	if err != nil {
		return err
	}
	
	// адрес получателя (StonfiRouter)
	jettonRouter := tongoTon.MustParseAccountID(constants.StonfiRouterAddr)
	// адрес монеты (откуда) NOT
	jettonMaster0 := tongoTon.MustParseAccountID(jettonCA)
	// адрес монеты (куда) TON
	jettonMaster1 := tongoTon.MustParseAccountID(constants.ProxyTonMasterAddr)
	
	// структура с информацией для Swap транзакции на DEX Stonfi
	stonfiStruct, err := tongoStonfi.NewStonfi(ctx, settings.TongoTonAPI, jettonRouter, jettonMaster0, jettonMaster1)
	if err != nil {
		settings.ErrorLog.Printf("Failed to create new stonfiStruct: %v", err)
		return err
	}
	fmt.Println("\nstonfiStruct", stonfiStruct)

	// TON для газовой комиссии (0.3 TON)
	gasToncoins := tongoTlb.Grams(300_000_000)
	// минимальное кол-во TON для возврата от газовой комиссии (0.2 TON)
	forwardToncoins := tongoTlb.Grams(200_000_000)
	// кол-во монет в виде *big.Int
	bigIntAmount := myTongoJettons.ConvertJettonsAmountToBigInt(jetton, amount)
	// адрес отправителя (кошелёк юзера)
	senderAddrID := tongoTon.MustParseAccountID(settings.JsonWallet.Hash)






	// перевод процента проскальзывания в часть от кол-ва монет в виде float64
	slippageAmount := amount - amount * (float64(slippage) / 100)
	// процент проскальзывания (часть от кол-ва монет) в виде *big.Int
	minOut := myTongoJettons.ConvertJettonsAmountToBigInt(jetton, slippageAmount)

	fmt.Printf("bigIntAmount: %v | minOut: %v\n", bigIntAmount, minOut)

	// 0.01 TON
	minOut = big.NewInt(10_000_000)

	fmt.Printf("NEW bigIntAmount: %v | minOut: %v\n", bigIntAmount, minOut)



	// структура для совершения Swap транзакции
	jettonTransfer, err := stonfiStruct.MakeSwapMessage(gasToncoins, forwardToncoins, *bigIntAmount, *minOut, senderAddrID)
	if err != nil {
		settings.ErrorLog.Printf("Failed to make swap message: %v", err)
		return err
	}

	fmt.Println("\njettonTransfer", jettonTransfer)

	// отправка сообщения в блокчейн
	// fmt.Println("\nrealWallet:", realWallet)
	err = realWallet.Send(ctx, jettonTransfer)
	if err != nil {
		settings.ErrorLog.Printf("Failed to send transfer message: %v", err)
		return err
	}

	return nil
}
