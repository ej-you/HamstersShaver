package transactions

import (
	"fmt"
	"context"

	tongoStonfi "github.com/tonkeeper/tongo/contract/stonfi"
	tongoTlb "github.com/tonkeeper/tongo/tlb"
	tongoTon "github.com/tonkeeper/tongo/ton"

	myTonapiJettons "github.com/Danil-114195722/HamstersShaver/ton_api/tonapi/jettons"
	
	myTongoWallet "github.com/Danil-114195722/HamstersShaver/ton_api/tongo/wallet"
	myTongoJettons "github.com/Danil-114195722/HamstersShaver/ton_api/tongo/jettons"
	
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
	
	// адрес получателя
	// jettonRouter := myTongoAccount.GetAccountIDByAddress(settings.JsonWallet.Hash)
	jettonRouter := tongoTon.MustParseAccountID("EQB3ncyBUTjZUA5EnFKR5_EnOMI9V1tTEAAPaiU71gc4TiUt") // StonfiRouter
	// адрес монеты (откуда) NOT
	jettonMaster0 := tongoTon.MustParseAccountID("EQAvlWFDxGF2lXm67y4yzC17wYKD9A0guwPkMs1gOsM__NOT")
	// адрес монеты (куда) MEM
	jettonMaster1 := tongoTon.MustParseAccountID("0:16a73dbf1b434ac651b656f8056e06463edf18d6a7b47068fee18c3905f99847")
	
	// структура с информацией для Swap транзакции на DEX Stonfi
	stonfiStruct, err := tongoStonfi.NewStonfi(ctx, settings.TongoTonAPI, jettonRouter, jettonMaster0, jettonMaster1)
	if err != nil {
		settings.ErrorLog.Printf("Failed to create new stonfiStruct: %v", err)
		return err
	}
	fmt.Println("\nstonfiStruct", stonfiStruct)

	// TON для газовой комиссии (0.3 TON)
	gasToncoins := tongoTlb.Grams(300_000_000)
	// кол-во монет в виде *big.Int
	bigIntAmount := myTongoJettons.ConvertJettonsAmountToBigInt(jetton, amount)
	// адрес отправителя (тот же, что и получателя)
	senderAddrID := tongoTon.MustParseAccountID(settings.JsonWallet.Hash)

	// перевод процента проскальзывания в часть от кол-ва монет в виде float64
	slippageAmount := amount - amount * (float64(slippage) / 100)
	// процент проскальзывания (часть от кол-ва монет) в виде *big.Int
	minOut := myTongoJettons.ConvertJettonsAmountToBigInt(jetton, slippageAmount)

	fmt.Printf("bigIntAmount: %v | minOut: %v\n", bigIntAmount, minOut)

	// структура для совершения Swap транзакции
	jettonTransfer, err := stonfiStruct.MakeSwapMessage(gasToncoins, tongoTlb.Grams(200_000_000), *bigIntAmount, *minOut, senderAddrID)
	if err != nil {
		settings.ErrorLog.Printf("Failed to make swap message: %v", err)
		return err
	}

	fmt.Println("\njettonTransfer", jettonTransfer)

	// отправка сообщения в блокчейн
	// fmt.Println("realWallet:", realWallet)
	err = realWallet.Send(ctx, jettonTransfer)
	if err != nil {
		settings.ErrorLog.Printf("Failed to send transfer message: %v", err)
		return err
	}

	return nil
}
