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
// 	jettonRouter := tongoTon.MustParseAccountID(constants.StonfiRouterAddr)
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

// 	// TON для газовой комиссии (0.6 TON)
// 	gasToncoins := tongoTlb.Grams(600_000_000)
// 	// минимальное кол-во TON для возврата от газовой комиссии (0.2 TON)
// 	forwardToncoins := tongoTlb.Grams(200_000_000)
// 	// кол-во TON в виде *big.Int
// 	bigIntAmount := myTongoJettons.ConvertJettonsAmountToBigInt(9, amount)
// 	// адрес отправителя (кошелёк юзера)
// 	senderAddrID := tongoTon.MustParseAccountID(settings.JsonWallet.Hash)

// 	// информация о пуле продаваемой монеты и TON
// 	poolInfo, err := myDexscreenerJettons.GetJettonsPoolInfo(constants.ProxyTonMasterAddr, jettonCA)
// 	if err != nil {
// 		return err
// 	}

// 	// предположительное кол-во монет на выходе без учёта изменения цены
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

	// информация о пуле покупаемой монеты и TON
	poolInfo, err := myDexscreenerJettons.GetJettonsPoolInfo(constants.ProxyTonMasterAddr, jettonCA)
	if err != nil {
		return err
	}
	fmt.Println("\npoolInfo: ", poolInfo)


	// адрес получателя (StonfiRouter)
	jettonRouter := tongoTon.MustParseAccountID(constants.StonfiRouterAddr)
	// адрес монеты (откуда) TON
	jettonMaster0 := tongoTon.MustParseAccountID(constants.ProxyTonMasterAddr)
	// адрес монеты (куда) jettonCA
	jettonMaster1 := tongoTon.MustParseAccountID(jettonCA)
	
	// структура с информацией для Swap транзакции на DEX Stonfi
	// stonfiStruct, err := tongoStonfi.NewStonfi(ctx, settings.TongoTonAPI, jettonRouter, jettonMaster0, jettonMaster1)
	// if err != nil {
	// 	settings.ErrorLog.Printf("Failed to create new stonfiStruct: %v", err)
	// 	return err
	// }

	jet1 := tongoJetton.New(jettonMaster1, settings.TongoTonAPI)
	token1, _ := jet1.GetJettonWallet(ctx, jettonRouter)  // router's

	fmt.Println("\ntoken1", token1)


	// TON для газовой комиссии (0.4 TON)
	gasToncoins := tongoTlb.Grams(400_000_000)
	// минимальное кол-во TON для возврата от газовой комиссии (0.2 TON)
	forwardToncoins := tongoTlb.Grams(200_000_000)
	// кол-во TON в виде *big.Int
	bigIntAmount := myTongoJettons.ConvertJettonsAmountToBigInt(9, amount)
	// адрес отправителя (кошелёк юзера)
	senderAddrID := tongoTon.MustParseAccountID(settings.JsonWallet.Hash)


	// предположительное кол-во монет на выходе без учёта изменения цены
	predictedJettonsAmount := amount / poolInfo.PriceNative
	// перевод процента проскальзывания в часть от кол-ва монет в виде float64
	slippageAmount := predictedJettonsAmount * (1.0 - float64(slippage) / 100)
	// процент проскальзывания (часть от кол-ва монет) в виде *big.Int
	minOut := myTongoJettons.ConvertJettonsAmountToBigInt(jetton.Decimals, slippageAmount)

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
		TokenWallet: token1.ToMsgAddress(),
		MinOut:      tongoTlb.VarUInteger16(*minOut),
		ToAddress:   senderAddrID.ToMsgAddress(),
	}

	// преобразование ForwardPayload в cell представление
	cell := tongoBoc.NewCell()
	if err := cell.WriteUint(0x25938561, 32); err != nil {
		return err
	}
	if err := tongoTlb.Marshal(cell, payload); err != nil {
		return err
	}

	jettonTransfer := tongoJetton.TransferMessage{
		Sender:           	 jettonRouter,  // это для задания поля dest (в значение jetton_wallet stonfi_router'а)
		Jetton:           	 tongoJetton.New(jettonMaster0, settings.TongoTonAPI),
		JettonAmount:     	 bigIntAmount,
		Destination:      	 jettonRouter,
		AttachedTon:      	 gasToncoins,
		ForwardTonAmount: 	 forwardToncoins,
		ForwardPayload:   	 cell,
	}


	fmt.Println("\njettonTransfer", jettonTransfer)

	// отправка сообщения в блокчейн
	fmt.Println("\nrealWallet:", realWallet)
	// err = realWallet.Send(ctx, jettonTransfer)
	// if err != nil {
	// 	settings.ErrorLog.Printf("Failed to send transfer message: %v", err)
	// 	return err
	// }

	// msgHash, err := realWallet.SendV2(ctx, 0, jettonTransfer)
	// fmt.Printf("\nmsgHash: %v\nhex msgHash: %s\n", msgHash, msgHash.Hex())


	// ПОПРОБОВАТЬ ПРОВЕРЯТЬ SEQNO КОШЕЛЬКА ДЛЯ ОТСЛЕЖИВАНИЯ, ПРОШЛА ЛИ ТРАНЗАКЦИЯ И БЫЛО ЛИ УСПЕШНО
	return nil
}
