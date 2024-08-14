package transactions

import (
	"math"
	"math/big"

	"fmt"
	"context"

	tongoStonfi "github.com/tonkeeper/tongo/contract/stonfi"
	tongoTlb "github.com/tonkeeper/tongo/tlb"
	tongoTon "github.com/tonkeeper/tongo/ton"

	myDexscreenerJettons "github.com/Danil-114195722/HamstersShaver/ton_api/dexscreener/jettons"
	myTonapiJettons "github.com/Danil-114195722/HamstersShaver/ton_api/tonapi/jettons"
	
	myTongoWallet "github.com/Danil-114195722/HamstersShaver/ton_api/tongo/wallet"
	myTongoJettons "github.com/Danil-114195722/HamstersShaver/ton_api/tongo/jettons"
	
	"github.com/Danil-114195722/HamstersShaver/settings/constants"
	"github.com/Danil-114195722/HamstersShaver/settings"
)


// TODO: 10 попыток до успеха (ошибка "error code: 651 message: cannot load block")
// продажа монет (Jetton -> TON)
func CellJetton(ctx context.Context, jettonCA string, jetton myTonapiJettons.AccountJetton, amount float64, slippage int) error {
	// получение данных о кошельке через tongo
	realWallet, err := myTongoWallet.GetWallet()
	if err != nil {
		return err
	}
	
	// адрес получателя (StonfiRouter)
	jettonRouter := tongoTon.MustParseAccountID(constants.StonfiRouterAddr)
	// адрес монеты (откуда) jettonCA
	jettonMaster0 := tongoTon.MustParseAccountID(jettonCA)
	// адрес монеты (куда) TON
	jettonMaster1 := tongoTon.MustParseAccountID(constants.ProxyTonMasterAddr)
	
	// структура с информацией для Swap транзакции на DEX Stonfi
	stonfiStruct, err := tongoStonfi.NewStonfi(ctx, settings.TongoTonAPI, jettonRouter, jettonMaster0, jettonMaster1)
	if err != nil {
		settings.ErrorLog.Printf("Failed to create new stonfiStruct: %v", err)
		return err
	}

	// TON для газовой комиссии (0.3 TON)
	gasToncoins := tongoTlb.Grams(300_000_000)
	// минимальное кол-во TON для возврата от газовой комиссии (0.2 TON)
	forwardToncoins := tongoTlb.Grams(200_000_000)
	// кол-во монет в виде *big.Int
	bigIntAmount := myTongoJettons.ConvertJettonsAmountToBigInt(jetton, amount)
	// адрес отправителя (кошелёк юзера)
	senderAddrID := tongoTon.MustParseAccountID(settings.JsonWallet.Hash)

	// информация о пуле продаваемой монеты и TON
	poolInfo, err := myDexscreenerJettons.GetJettonsPoolInfo(constants.ProxyTonMasterAddr, jettonCA)
	if err != nil {
		return err
	}

	// предположительное кол-во TON на выходе без учёта изменения цены и газовой комиссии
	predictedTonAmount := amount * poolInfo.PriceNative
	// перевод процента проскальзывания в часть от кол-ва монет в виде float64
	slippageAmount := predictedTonAmount * (1.0 - float64(slippage) / 100)
	// процент проскальзывания (часть от кол-ва монет) в виде *big.Int
	minOut := big.NewInt(int64(slippageAmount * math.Pow10(9)))

	fmt.Printf("\nbigIntAmount: %v | predictedTonAmount: %v | minOut: %v\n", bigIntAmount, predictedTonAmount, minOut)


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
