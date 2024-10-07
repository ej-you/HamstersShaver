package transactions

// import (
// 	"fmt"
// 	"context"
// 	// "time"

// 	tongoStonfi "github.com/tonkeeper/tongo/contract/stonfi"
// 	tongoTlb "github.com/tonkeeper/tongo/tlb"
// 	tongoTon "github.com/tonkeeper/tongo/ton"

// 	myStonfiJettons "github.com/Danil-114195722/HamstersShaver/ton_api/stonfi/jettons"
// 	myTongoWallet "github.com/Danil-114195722/HamstersShaver/ton_api/tongo/wallet"
// 	myTongoServices "github.com/Danil-114195722/HamstersShaver/ton_api/tongo/services"

// 	"github.com/Danil-114195722/HamstersShaver/settings/constants"
// 	"github.com/Danil-114195722/HamstersShaver/settings"
// )



// // TODO: 10 попыток до успеха (ошибка "error code: 651 message: cannot load block")
// // продажа монет (Jetton -> TON)
// func CellJetton(ctx context.Context, jettonCA string, amount float64, slippage int) error {
// 	// получение данных о кошельке через tongo
// 	realWallet, err := myTongoWallet.GetWallet()
// 	if err != nil {
// 		return err
// 	}
	
// 	// адрес получателя (StonfiRouter)
// 	jettonRouter := tongoTon.MustParseAccountID(constants.StonfiRouterAddr)
// 	// адрес монеты (откуда) jettonCA
// 	jettonMaster0 := tongoTon.MustParseAccountID(jettonCA)
// 	// адрес монеты (куда) TON
// 	jettonMaster1 := tongoTon.MustParseAccountID(constants.ProxyTonMasterAddr)
	
// 	// структура с информацией для Swap транзакции на DEX Stonfi
// 	stonfiStruct, err := tongoStonfi.NewStonfi(ctx, settings.TongoTonAPI, jettonRouter, jettonMaster0, jettonMaster1)
// 	if err != nil {
// 		settings.ErrorLog.Printf("Failed to create new stonfiStruct: %v", err)
// 		return err
// 	}

// 	// получение данных о покупаемой монете
// 	jettonInfo, err := myStonfiJettons.GetJettonInfoByAddres(jettonCA)
// 	if err != nil {
// 		return err
// 	}
// 	// получение данных о TON
// 	tonInfo, err := myStonfiJettons.GetJettonInfoByAddres(constants.TonInfoAddr)
// 	if err != nil {
// 		return err
// 	}
// 	// цена монеты в TON
// 	jettonPriceInTON := jettonInfo.PriceUSD / tonInfo.PriceUSD


// 	// TON для газовой комиссии (0.3 TON)
// 	gasToncoins := tongoTlb.Grams(300_000_000)
// 	// TON для передачи в следующее сообщение цепочки транзакций (0.2 TON)
// 	forwardToncoins := tongoTlb.Grams(200_000_000)
// 	// кол-во монет в виде *big.Int
// 	bigIntAmount := myTongoServices.ConvertJettonsAmountToBigInt(jettonInfo.Decimals, amount)
// 	// адрес отправителя (кошелёк юзера)
// 	senderAddrID := tongoTon.MustParseAccountID(settings.JsonWallet.Hash)

// 	// предположительное кол-во TON на выходе без учёта изменения цены и газовой комиссии
// 	predictedTonAmount := amount * jettonPriceInTON
// 	// перевод процента проскальзывания в часть от кол-ва TON в виде float64
// 	slippageAmount := predictedTonAmount * (1.0 - float64(slippage) / 100)
// 	// процент проскальзывания (часть от кол-ва TON) в виде *big.Int
// 	minOut := myTongoServices.ConvertJettonsAmountToBigInt(constants.TonDecimals, slippageAmount)

// 	fmt.Printf("\nbigIntAmount: %v | predictedTonAmount: %v | minOut: %v\n", bigIntAmount, predictedTonAmount, minOut)


// 	// структура для совершения Swap транзакции
// 	jettonTransfer, err := stonfiStruct.MakeSwapMessage(gasToncoins, forwardToncoins, *bigIntAmount, *minOut, senderAddrID)
// 	if err != nil {
// 		settings.ErrorLog.Printf("Failed to make swap message: %v", err)
// 		return err
// 	}

// 	fmt.Println("\njettonTransfer", jettonTransfer)

// 	// отправка сообщения в блокчейн
// 	fmt.Println("\nrealWallet:", realWallet)
// 	// err = realWallet.Send(ctx, jettonTransfer)
// 	// if err != nil {
// 	// 	settings.ErrorLog.Printf("Failed to send transfer message: %v", err)
// 	// 	return err
// 	// }

// 	return nil
// }
