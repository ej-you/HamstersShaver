package transactions

import (
	"context"
	"fmt"

	tongoAbi "github.com/tonkeeper/tongo/abi"
	tongoBoc "github.com/tonkeeper/tongo/boc"
	tongoTlb "github.com/tonkeeper/tongo/tlb"
	tongoTon "github.com/tonkeeper/tongo/ton"
	tongoJetton "github.com/tonkeeper/tongo/contract/jetton"

	myStonfiJettons "github.com/ej-you/HamstersShaver/rest_api/ton_api_rest/stonfi/jettons"
	
	myTongoWallet "github.com/ej-you/HamstersShaver/rest_api/ton_api_rest/tongo/wallet"
	myTongoServices "github.com/ej-you/HamstersShaver/rest_api/ton_api_rest/tongo/services"

	myTonapiServices "github.com/ej-you/HamstersShaver/rest_api/ton_api_rest/tonapi/services"
	
	coreErrors "github.com/ej-you/HamstersShaver/rest_api/core/errors"
	"github.com/ej-you/HamstersShaver/rest_api/settings/constants"
	"github.com/ej-you/HamstersShaver/rest_api/settings"
)


// данные о последующей транзакции покупки монет (TON -> Jetton)
type PreRequestBuyJetton struct {
	UsedTON 		string `json:"usedTon" example:"0.1" description:"кол-во используемых TON для покупки в формате, удобном для человека"`
	JettonCA 		string `json:"jettonCA" example:"EQC47093oX5Xhb0xuk2lCr2RhS8rj-vul61u4W2UH5ORmG_O" description:"мастер-адрес покупаемой монеты (jetton_master)"`
	DEX 			string `json:"dex" example:"Stonfi" description:"название DEX биржи"`
	JettonsOut 		string `json:"jettonsOut" example:"175.137" description:"примерное кол-во монет, которые будут получены после транзакции"`
	MinOut	 		string `json:"minOut" example:"140.11" description:"минимальное кол-во покупаемых монет (с учётом процента проскальзывания) "`
	JettonSymbol 	string `json:"jettonSymbol" example:"GRAM" description:"символ покупаемой монеты"`
}


// получение данных на подтверждение последующей транзакции покупки монет (TON -> Jetton)
func GetPreRequestBuyJetton(jettonCA string, tonAmount float64, slippage int) (PreRequestBuyJetton, error) {
	var preRequestInfo PreRequestBuyJetton

	// получение данных о покупаемой монете
	jettonInfo, err := myStonfiJettons.GetJettonInfoByAddressWithTimeout(jettonCA, constants.GetJettonInfoByAddressTimeout)
	if err != nil {
		return preRequestInfo, fmt.Errorf("get buy pre request: %w", err)
	}
	// получение данных о TON
	tonInfo, err := myStonfiJettons.GetJettonInfoByAddressWithTimeout(constants.TonInfoAddr, constants.GetJettonInfoByAddressTimeout)
	if err != nil {
		return preRequestInfo, fmt.Errorf("get buy pre request: %w", err)
	}
	// цена монеты в TON
	jettonPriceInTON := jettonInfo.PriceUSD / tonInfo.PriceUSD

	// предположительное кол-во монет на выходе без учёта изменения цены
	predictedJettonsAmount := tonAmount / jettonPriceInTON
	// перевод процента проскальзывания в часть от кол-ва монет в виде float64
	slippageAmount := predictedJettonsAmount * (1.0 - float64(slippage) / 100)

	preRequestInfo = PreRequestBuyJetton{
		UsedTON: myTonapiServices.JettonFloatAmountFormat(tonAmount, tonInfo.Decimals),
		JettonCA: jettonInfo.MasterAddress,
		DEX: "Stonfi",
		JettonsOut: myTonapiServices.JettonFloatAmountFormat(predictedJettonsAmount, jettonInfo.Decimals),
		MinOut: myTonapiServices.JettonFloatAmountFormat(slippageAmount, jettonInfo.Decimals),
		JettonSymbol: jettonInfo.Symbol,
	}
	return preRequestInfo, nil
}


// покупка монет (TON -> Jetton)
func BuyJetton(ctx context.Context, jettonCA string, amount float64, slippage int) error {
	// создание API клиента TON для tongo
	tongoClient, err := settings.GetTonClientTongoWithTimeout("mainnet", constants.TongoClientTimeout)
	if err != nil {
		return fmt.Errorf("send buy transaction: %w", err)
	}

	// получение данных о кошельке через tongo
	realWallet, err := myTongoWallet.GetWallet(tongoClient)
	if err != nil {
		return fmt.Errorf("send buy transaction: %w", err)
	}
	// получение данных о покупаемой монете
	jettonInfo, err := myStonfiJettons.GetJettonInfoByAddressWithTimeout(jettonCA, constants.GetJettonInfoByAddressTimeout)
	if err != nil {
		return fmt.Errorf("send buy transaction: %w", err)
	}
	// получение данных о TON
	tonInfo, err := myStonfiJettons.GetJettonInfoByAddressWithTimeout(constants.TonInfoAddr, constants.GetJettonInfoByAddressTimeout)
	if err != nil {
		return fmt.Errorf("send buy transaction: %w", err)
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
	jettonToBuy := tongoJetton.New(jettonMaster1, tongoClient)
	routersJettonWallet, err := jettonToBuy.GetJettonWallet(ctx, jettonRouter)
	if err != nil {
		apiErr := coreErrors.New(
			fmt.Errorf("send buy transaction: get jetton wallet using jetton master: %w", err),
			"failed to get jetton wallet",
			"ton_api",
			500,
		)
		apiErr.CheckTimeout()
		return apiErr
	}

	// кол-во TON для покупки монет (в *big.Int)
	bigIntAmount := myTongoServices.ConvertJettonsAmountToBigInt(constants.TonDecimals, amount)
	// кол-во TON для покупки монет (в uint64)
	tonAmount := myTongoServices.ConvertJettonsAmountToUint(constants.TonDecimals, amount)

	// TON для газовой комиссии (0.3 TON)
	gasToncoins := tongoTlb.Grams(300_000_000)
	// прикреплённые TON для газа в сумме с TON для покупки монет
	attachedToncoins := gasToncoins + tongoTlb.Grams(tonAmount)


	// предположительное кол-во монет на выходе без учёта изменения цены
	predictedJettonsAmount := amount / jettonPriceInTON
	// перевод процента проскальзывания в часть от кол-ва монет в виде float64
	slippageAmount := predictedJettonsAmount * (1.0 - float64(slippage) / 100)
	// процент проскальзывания (часть от кол-ва монет) в виде *big.Int
	minOut := myTongoServices.ConvertJettonsAmountToBigInt(jettonInfo.Decimals, slippageAmount)

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
		Jetton:           	 tongoJetton.New(jettonMaster0, tongoClient),
		JettonAmount:     	 bigIntAmount,
		Destination:      	 jettonRouter,
		AttachedTon:      	 attachedToncoins,
		ForwardTonAmount: 	 gasToncoins,
		ForwardPayload:   	 cell,
	}

	// отправка сообщения в блокчейн
	if err := realWallet.Send(ctx, jettonTransfer); err != nil {
		apiErr := coreErrors.New(
			fmt.Errorf("send buy transaction: send transfer message: %w", err),
			"failed to send transfer message",
			"ton_api",
			500,
		)
		apiErr.CheckTimeout()
		return apiErr
	}
	return nil
}
