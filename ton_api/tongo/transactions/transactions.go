package transactions

import (
	"fmt"
	"context"

	tongoStonfi "github.com/tonkeeper/tongo/contract/stonfi"
	tongoAbi "github.com/tonkeeper/tongo/abi"
	tongoBoc "github.com/tonkeeper/tongo/boc"
	tongoTlb "github.com/tonkeeper/tongo/tlb"
	tongoTon "github.com/tonkeeper/tongo/ton"
	tongoJetton "github.com/tonkeeper/tongo/contract/jetton"

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
	bigIntAmount := myTongoJettons.ConvertJettonsAmountToBigInt(jetton.Decimals, amount)
	// адрес отправителя (кошелёк юзера)
	senderAddrID := tongoTon.MustParseAccountID(settings.JsonWallet.Hash)

	// информация о пуле продаваемой монеты и TON
	poolInfo, err := myDexscreenerJettons.GetJettonsPoolInfo(constants.ProxyTonMasterAddr, jettonCA)
	if err != nil {
		return err
	}

	// предположительное кол-во TON на выходе без учёта изменения цены и газовой комиссии
	predictedTonAmount := amount * poolInfo.PriceNative
	// перевод процента проскальзывания в часть от кол-ва TON в виде float64
	slippageAmount := predictedTonAmount * (1.0 - float64(slippage) / 100)
	// процент проскальзывания (часть от кол-ва TON) в виде *big.Int
	minOut := myTongoJettons.ConvertJettonsAmountToBigInt(9, slippageAmount)

	fmt.Printf("\nbigIntAmount: %v | predictedTonAmount: %v | minOut: %v\n", bigIntAmount, predictedTonAmount, minOut)


	// структура для совершения Swap транзакции
	jettonTransfer, err := stonfiStruct.MakeSwapMessage(gasToncoins, forwardToncoins, *bigIntAmount, *minOut, senderAddrID)
	if err != nil {
		settings.ErrorLog.Printf("Failed to make swap message: %v", err)
		return err
	}

	fmt.Println("\njettonTransfer", jettonTransfer)

	// отправка сообщения в блокчейн
	fmt.Println("\nrealWallet:", realWallet)
	// err = realWallet.Send(ctx, jettonTransfer)
	// if err != nil {
	// 	settings.ErrorLog.Printf("Failed to send transfer message: %v", err)
	// 	return err
	// }

	return nil
}


// покупка монет (TON -> Jetton)
// func BuyJetton(ctx context.Context, jettonCA string, amount float64, slippage int) error {
// 	// получение данных о кошельке через tongo
// 	realWallet, err := myTongoWallet.GetWallet()
// 	if err != nil {
// 		return err
// 	}
// 	// получение данных о покупаемой монете
// 	jetton, err := myTonapiJettons.GetJettonInfoByAddres(ctx, jettonCA)
// 	if err != nil {
// 		return err
// 	}
	
// 	// адрес получателя (StonfiRouter)
// 	jettonRouter := tongoTon.MustParseAccountID(settings.JsonWallet.Hash)
// 	// адрес монеты (откуда)
// 	jettonMaster0 := tongoTon.MustParseAccountID(constants.ProxyTonMasterAddr)
// 	// адрес монеты (куда)
// 	jettonMaster1 := tongoTon.MustParseAccountID(jettonCA)
	
// 	// структура с информацией для Swap транзакции на DEX Stonfi
// 	stonfiStruct, err := tongoStonfi.NewStonfi(ctx, settings.TongoTonAPI, jettonRouter, jettonMaster0, jettonMaster1)
// 	if err != nil {
// 		settings.ErrorLog.Printf("Failed to create new stonfiStruct: %v", err)
// 		return err
// 	}

// 	// TON для газовой комиссии (0.3 TON)
// 	gasToncoins := tongoTlb.Grams(300_000_000)
// 	// минимальное кол-во TON для возврата от газовой комиссии (0.01 TON)
// 	forwardToncoins := tongoTlb.Grams(10_000_000)
// 	// кол-во монет в виде *big.Int
// 	bigIntAmount := myTongoJettons.ConvertJettonsAmountToBigInt(9, amount)
// 	// адрес отправителя (кошелёк юзера)
// 	senderAddrID := tongoTon.MustParseAccountID(constants.StonfiRouterAddr)

// 	// информация о пуле продаваемой монеты и TON
// 	poolInfo, err := myDexscreenerJettons.GetJettonsPoolInfo(constants.ProxyTonMasterAddr, jettonCA)
// 	if err != nil {
// 		return err
// 	}

// 	// предположительное кол-во монет на выходе без учёта изменения цены и газовой комиссии
// 	predictedJettonsAmount := amount / poolInfo.PriceNative
// 	// перевод процента проскальзывания в часть от кол-ва монет в виде float64
// 	slippageAmount := predictedJettonsAmount * (1.0 - float64(slippage) / 100)
// 	// процент проскальзывания (часть от кол-ва монет) в виде *big.Int
// 	minOut := myTongoJettons.ConvertJettonsAmountToBigInt(jetton.Decimals, slippageAmount)

// 	fmt.Printf("\nbigIntAmount: %v | predictedJettonsAmount: %v | minOut: %v\n", bigIntAmount, predictedJettonsAmount, minOut)


// 	// структура для совершения Swap транзакции
// 	jettonTransfer, err := stonfiStruct.MakeSwapMessage(gasToncoins, forwardToncoins, *bigIntAmount, *minOut, senderAddrID)
// 	if err != nil {
// 		settings.ErrorLog.Printf("Failed to make swap message: %v", err)
// 		return err
// 	}
// 	jettonTransfer.ResponseDestination = &jettonRouter

// 	// jettonTransfer := tongoJetton.TransferMessage{
// 	// 	Jetton:           tongoJetton.New(jettonMaster1, settings.TongoTonAPI),
// 	// 	JettonAmount:     bigIntAmount,
// 	// 	Destination:      senderAddrID,
// 	// 	AttachedTon:      gasToncoins,
// 	// 	ForwardTonAmount: forwardToncoins,
// 	// }

// 	fmt.Println("\njettonTransfer", jettonTransfer)

// 	// отправка сообщения в блокчейн
// 	// fmt.Println("\nrealWallet:", realWallet)
// 	err = realWallet.Send(ctx, jettonTransfer)
// 	if err != nil {
// 		settings.ErrorLog.Printf("Failed to send transfer message: %v", err)
// 		return err
// 	}

// 	return nil
// }


// покупка монет (TON -> Jetton)
func BuyJetton(ctx context.Context, jettonCA string, amount float64, slippage int) error {
	// получение данных о кошельке через tongo
	realWallet, err := myTongoWallet.GetWallet()
	if err != nil {
		return err
	}
	// получение данных о покупаемой монете
	jetton, err := myTonapiJettons.GetJettonInfoByAddres(ctx, jettonCA)
	if err != nil {
		return err
	}
	
	// адрес получателя (StonfiRouter)
	jettonRouter := tongoTon.MustParseAccountID(constants.StonfiRouterAddr)
	// адрес монеты (куда)
	jettonMaster0 := tongoTon.MustParseAccountID(jettonCA)
	// адрес монеты (откуда)
	jettonMaster1 := tongoTon.MustParseAccountID(constants.ProxyTonMasterAddr)
	
	// структура с информацией для Swap транзакции на DEX Stonfi
	// stonfiStruct, err := tongoStonfi.NewStonfi(ctx, settings.TongoTonAPI, jettonRouter, jettonMaster0, jettonMaster1)
	// if err != nil {
	// 	settings.ErrorLog.Printf("Failed to create new stonfiStruct: %v", err)
	// 	return err
	// }


	jet0 := tongoJetton.New(jettonMaster0, settings.TongoTonAPI)
	token0, _ := jet0.GetJettonWallet(ctx, jettonRouter)

	jet1 := tongoJetton.New(jettonMaster1, settings.TongoTonAPI)
	// token1, _ := jet1.GetJettonWallet(ctx, jettonRouter)
	token1_my, _ := jet1.GetJettonWallet(ctx, tongoTon.MustParseAccountID(settings.JsonWallet.Hash))

	// fmt.Println("\ntoken1", token1)
	fmt.Println("token1_my", token1_my)


	// stonfiStruct := tongoStonfi.Stonfi{
	// 	cli: settings.TongoTonAPI,
	// 	router: jettonRouter,
	// 	master0: jettonMaster0,
	// 	token0: token0,
	// 	master1: jettonMaster1,
	// 	token1: token1,
	// }

	// TON для газовой комиссии (0.3 TON)
	gasToncoins := tongoTlb.Grams(300_000_000)
	// минимальное кол-во TON для возврата от газовой комиссии (0.1 TON)
	forwardToncoins := tongoTlb.Grams(100_000_000)
	// кол-во монет в виде *big.Int
	bigIntAmount := myTongoJettons.ConvertJettonsAmountToBigInt(jetton.Decimals, amount)
	// адрес отправителя (кошелёк юзера)
	senderAddrID := tongoTon.MustParseAccountID(settings.JsonWallet.Hash)

	// информация о пуле продаваемой монеты и TON
	poolInfo, err := myDexscreenerJettons.GetJettonsPoolInfo(constants.ProxyTonMasterAddr, jettonCA)
	if err != nil {
		return err
	}

	// предположительное кол-во монет на выходе без учёта изменения цены и газовой комиссии
	predictedJettonsAmount := amount * poolInfo.PriceNative
	// перевод процента проскальзывания в часть от кол-ва монет в виде float64
	slippageAmount := predictedJettonsAmount * (1.0 - float64(slippage) / 100)
	// процент проскальзывания (часть от кол-ва монет) в виде *big.Int
	minOut := myTongoJettons.ConvertJettonsAmountToBigInt(9, slippageAmount)

	fmt.Printf("\nbigIntAmount: %v | predictedJettonsAmount: %v | minOut: %v\n", bigIntAmount, predictedJettonsAmount, minOut)


	// структура для совершения Swap транзакции
	// jettonTransfer, err := stonfiStruct.MakeSwapMessage(gasToncoins, forwardToncoins, *bigIntAmount, *minOut, senderAddrID)
	// if err != nil {
	// 	settings.ErrorLog.Printf("Failed to make swap message: %v", err)
	// 	return err
	// }
	// jettonTransfer.ResponseDestination = &senderAddrID

	// jettonTransfer := tongoJetton.TransferMessage{
	// 	Jetton:           tongoJetton.New(jettonMaster1, settings.TongoTonAPI),
	// 	JettonAmount:     bigIntAmount,
	// 	Destination:      senderAddrID,
	// 	AttachedTon:      gasToncoins,
	// 	ForwardTonAmount: forwardToncoins,
	// }


	payload := tongoAbi.StonfiSwapJettonPayload{
		TokenWallet: token0.ToMsgAddress(),  // !!! check !!!
		MinOut:      tongoTlb.VarUInteger16(*minOut),
		ToAddress:   senderAddrID.ToMsgAddress(),
	}
	cell := tongoBoc.NewCell()
	if err := cell.WriteUint(0x25938561, 32); err != nil {
		return err
	}
	if err := tongoTlb.Marshal(cell, payload); err != nil {
		return err
	}
	jettonTransfer := tongoJetton.TransferMessage{
		Sender:           	 token1_my,// senderAddrID,  // !!! check !!!
		Jetton:           	 tongoJetton.New(jettonMaster1, settings.TongoTonAPI),  // !!! check !!!
		JettonAmount:     	 bigIntAmount,
		Destination:      	 jettonRouter,  // !!! check !!!
		AttachedTon:      	 gasToncoins,
		ForwardTonAmount: 	 forwardToncoins,
		ForwardPayload:   	 cell,
		ResponseDestination: &senderAddrID,
	}



	fmt.Println("\njettonTransfer", jettonTransfer)

	// отправка сообщения в блокчейн
	fmt.Println("\nrealWallet:", realWallet)
	// err = realWallet.Send(ctx, jettonTransfer)
	// if err != nil {
	// 	settings.ErrorLog.Printf("Failed to send transfer message: %v", err)
	// 	return err
	// }

	return nil
}
