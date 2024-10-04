package transactions

import (
	"fmt"
	"context"
	// "time"

	tongoStonfi "github.com/tonkeeper/tongo/contract/stonfi"
	tongoAbi "github.com/tonkeeper/tongo/abi"
	tongoBoc "github.com/tonkeeper/tongo/boc"
	tongoTlb "github.com/tonkeeper/tongo/tlb"
	tongoTon "github.com/tonkeeper/tongo/ton"
	tongoJetton "github.com/tonkeeper/tongo/contract/jetton"

	tonapi "github.com/tonkeeper/tonapi-go"

	myStonfiJettons "github.com/Danil-114195722/HamstersShaver/ton_api/stonfi/jettons"
	myTongoWallet "github.com/Danil-114195722/HamstersShaver/ton_api/tongo/wallet"
	myTongoJettons "github.com/Danil-114195722/HamstersShaver/ton_api/tongo/jettons"

	"github.com/Danil-114195722/HamstersShaver/settings/constants"
	"github.com/Danil-114195722/HamstersShaver/settings"
)



// TODO: 10 попыток до успеха (ошибка "error code: 651 message: cannot load block")
// продажа монет (Jetton -> TON)
func CellJetton(ctx context.Context, jettonCA string, amount float64, slippage int) error {
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

	// получение данных о покупаемой монете
	jettonInfo, err := myStonfiJettons.GetJettonInfoByAddres(jettonCA)
	if err != nil {
		return err
	}
	// получение данных о TON
	tonInfo, err := myStonfiJettons.GetJettonInfoByAddres(constants.TonInfoAddr)
	if err != nil {
		return err
	}
	// цена монеты в TON
	jettonPriceInTON := jettonInfo.PriceUSD / tonInfo.PriceUSD


	// TON для газовой комиссии (0.3 TON)
	gasToncoins := tongoTlb.Grams(300_000_000)
	// TON для передачи в следующее сообщение цепочки транзакций (0.2 TON)
	forwardToncoins := tongoTlb.Grams(200_000_000)
	// кол-во монет в виде *big.Int
	bigIntAmount := myTongoJettons.ConvertJettonsAmountToBigInt(jettonInfo.Decimals, amount)
	// адрес отправителя (кошелёк юзера)
	senderAddrID := tongoTon.MustParseAccountID(settings.JsonWallet.Hash)

	// предположительное кол-во TON на выходе без учёта изменения цены и газовой комиссии
	predictedTonAmount := amount * jettonPriceInTON
	// перевод процента проскальзывания в часть от кол-ва TON в виде float64
	slippageAmount := predictedTonAmount * (1.0 - float64(slippage) / 100)
	// процент проскальзывания (часть от кол-ва TON) в виде *big.Int
	minOut := myTongoJettons.ConvertJettonsAmountToBigInt(constants.TonDecimals, slippageAmount)

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
func BuyJetton(ctx context.Context, jettonCA string, amount float64, slippage int) error {
	// получение данных о кошельке через tongo
	realWallet, err := myTongoWallet.GetWallet()
	if err != nil {
		return err
	}
	// получение данных о покупаемой монете
	jettonInfo, err := myStonfiJettons.GetJettonInfoByAddres(jettonCA)
	if err != nil {
		return err
	}
	// получение данных о TON
	tonInfo, err := myStonfiJettons.GetJettonInfoByAddres(constants.TonInfoAddr)
	if err != nil {
		return err
	}
	// цена монеты в TON
	jettonPriceInTON := jettonInfo.PriceUSD / tonInfo.PriceUSD

	// адрес получателя (StonfiRouter)
	jettonRouter := tongoTon.MustParseAccountID(constants.StonfiRouterAddr)
	// адрес монеты (откуда) TON
	jettonMaster0 := tongoTon.MustParseAccountID(constants.ProxyTonMasterAddr)
	// адрес монеты (куда) jettonCA
	jettonMaster1 := tongoTon.MustParseAccountID(jettonCA)

	// получение jetton_wallet stonfi_router'а покупаемой монеты
	jettonToBuy := tongoJetton.New(jettonMaster1, settings.TongoTonAPI)
	routersJettonWallet, err := jettonToBuy.GetJettonWallet(ctx, jettonRouter)
	
	if err != nil {
		return err
	}

	// кол-во TON для покупки монет (в *big.Int)
	bigIntAmount := myTongoJettons.ConvertJettonsAmountToBigInt(constants.TonDecimals, amount)
	// кол-во TON для покупки монет (в uint64)
	tonAmount := myTongoJettons.ConvertJettonsAmountToUint(constants.TonDecimals, amount)

	// TON для газовой комиссии (0.3 TON)
	gasToncoins := tongoTlb.Grams(300_000_000)
	// прикреплённые TON для газа в сумме с TON для покупки монет
	attachedToncoins := gasToncoins + tongoTlb.Grams(tonAmount)


	// предположительное кол-во монет на выходе без учёта изменения цены
	predictedJettonsAmount := amount / jettonPriceInTON
	// перевод процента проскальзывания в часть от кол-ва монет в виде float64
	slippageAmount := predictedJettonsAmount * (1.0 - float64(slippage) / 100)
	// процент проскальзывания (часть от кол-ва монет) в виде *big.Int
	minOut := myTongoJettons.ConvertJettonsAmountToBigInt(jettonInfo.Decimals, slippageAmount)

	fmt.Printf("\nbigIntAmount: %v | predictedJettonsAmount: %v | minOut: %v\n", bigIntAmount, predictedJettonsAmount, minOut)

	// адрес получателя (кошелёк юзера)
	toAddrID := tongoTon.MustParseAccountID(settings.JsonWallet.Hash)

	// определение ForwardPayload
	payload := tongoAbi.StonfiSwapJettonPayload{
		TokenWallet: routersJettonWallet.ToMsgAddress(),
		MinOut:      tongoTlb.VarUInteger16(*minOut),
		ToAddress:   toAddrID.ToMsgAddress(),
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
		// jettonRouter - для задания поля dest (в значение jetton_wallet pTON'а stonfi_router'а)
		Sender:           	 jettonRouter,
		Jetton:           	 tongoJetton.New(jettonMaster0, settings.TongoTonAPI),
		JettonAmount:     	 bigIntAmount,
		Destination:      	 jettonRouter,
		AttachedTon:      	 attachedToncoins,
		ForwardTonAmount: 	 gasToncoins,
		ForwardPayload:   	 cell,
	}


	seqnoBefore, err := settings.TongoTonAPI.GetSeqno(ctx, realWallet.GetAddress())
	if err != nil {
		return err
	}
	fmt.Println("\nseqnoBefore: ", seqnoBefore)


	fmt.Println("\njettonTransfer", jettonTransfer)
	fmt.Println("\nrealWallet:", realWallet)
	
	// отправка сообщения в блокчейн
	// msgHash, err := realWallet.SendV2(ctx, 0, jettonTransfer)
	// if err != nil {
	// 	settings.ErrorLog.Printf("Failed to send transfer message: %v", err)
	// 	return err
	// }
	// fmt.Printf("\nmsgHash: %v\nhex msgHash: %s\n", msgHash, msgHash.Hex())

	msgHashParams := tonapi.GetBlockchainTransactionByMessageHashParams{MsgID: "4c6ba8804b6996325dc344fda7199a9bca7e8948b2cd9831e3d48a028cfd107a"}
	transactionInfo, err := settings.TonapiTonAPI.GetBlockchainTransactionByMessageHash(ctx, msgHashParams)
	if err != nil {
		return err
	}
	fmt.Println("\ntransactionInfo: ", transactionInfo)

	// GetBlockchainTransactionByMessageHash
	
	// msgHash: [76 107 168 128 75 105 150 50 93 195 68 253 167 25 154 155 202 126 137 72 178 205 152 49 227 212 138 2 140 253 16 122]
	// hex msgHash: 4c6ba8804b6996325dc344fda7199a9bca7e8948b2cd9831e3d48a028cfd107a

	// var seqnoAfter uint32
	// for i := 0; i < 25; i++ {
	// 	seqnoAfter, err = settings.TongoTonAPI.GetSeqno(ctx, realWallet.GetAddress())
	// 	if err == nil && seqnoAfter > seqnoBefore {
	// 		fmt.Printf("\nTransaction was sent successfully!")
	// 		return nil
	// 	}
	// 	time.Sleep(5 * time.Second)
	// }

	return nil
}
