package transactions

import (
	"context"
	"fmt"

	tongoStonfi "github.com/tonkeeper/tongo/contract/stonfi"
	tongoTlb "github.com/tonkeeper/tongo/tlb"
	tongoTon "github.com/tonkeeper/tongo/ton"

	myStonfiJettons "github.com/ej-you/HamstersShaver/rest_api/ton_api_rest/stonfi/jettons"
	myTongoWallet "github.com/ej-you/HamstersShaver/rest_api/ton_api_rest/tongo/wallet"
	myTongoServices "github.com/ej-you/HamstersShaver/rest_api/ton_api_rest/tongo/services"

	myTonapiServices "github.com/ej-you/HamstersShaver/rest_api/ton_api_rest/tonapi/services"

	coreErrors "github.com/ej-you/HamstersShaver/rest_api/core/errors"
	"github.com/ej-you/HamstersShaver/rest_api/settings/constants"
	"github.com/ej-you/HamstersShaver/rest_api/settings"
)


// данные о последующей транзакции продажи монет (Jetton -> TON)
type PreRequestCellJetton struct {
	UsedJettons		string `json:"usedJettons" example:"200.0" description:"кол-во используемых монет на продажу в формате, удобном для человека"`
	JettonCA 		string `json:"jettonCA" example:"EQC47093oX5Xhb0xuk2lCr2RhS8rj-vul61u4W2UH5ORmG_O" description:"мастер-адрес продаваемой монеты (jetton_master)"`
	DEX 			string `json:"dex" example:"Stonfi" description:"название DEX биржи"`
	TONsOut 		string `json:"tonsOut" example:"0.114" description:"примерное кол-во монет TON, которые будут получены после транзакции"`
	MinOut	 		string `json:"minOut" example:"0.091" description:"минимальное кол-во получаемых монет TON (с учётом процента проскальзывания)"`
	JettonSymbol 	string `json:"jettonSymbol" example:"GRAM" description:"символ продаваемой монеты"`
}


// получение данных на подтверждение последующей транзакции продажи монет (Jetton -> TON)
func GetPreRequestCellJetton(jettonCA string, jettonAmount float64, slippage int) (PreRequestCellJetton, error) {
	var preRequestInfo PreRequestCellJetton

	// получение данных о продаваемой монете
	jettonInfo, err := myStonfiJettons.GetJettonInfoByAddressWithTimeout(jettonCA, constants.GetJettonInfoByAddressTimeout)
	if err != nil {
		return preRequestInfo, fmt.Errorf("get cell pre request: %w", err)
	}
	// получение данных о TON
	tonInfo, err := myStonfiJettons.GetJettonInfoByAddressWithTimeout(constants.TonInfoAddr, constants.GetJettonInfoByAddressTimeout)
	if err != nil {
		return preRequestInfo, fmt.Errorf("get cell pre request: %w", err)
	}
	// цена монеты в TON
	jettonPriceInTON := jettonInfo.PriceUSD / tonInfo.PriceUSD

	// предположительное кол-во TON на выходе без учёта изменения цены
	predictedTonAmount := jettonAmount * jettonPriceInTON
	// перевод процента проскальзывания в часть от кол-ва монет в виде float64
	slippageAmount := predictedTonAmount * (1.0 - float64(slippage) / 100)

	preRequestInfo = PreRequestCellJetton{
		UsedJettons: myTonapiServices.BeautyJettonAmountFromFloat64(jettonAmount, jettonInfo.Decimals),
		JettonCA: jettonInfo.MasterAddress,
		DEX: "Stonfi",
		TONsOut: myTonapiServices.BeautyJettonAmountFromFloat64(predictedTonAmount, tonInfo.Decimals),
		MinOut: myTonapiServices.BeautyJettonAmountFromFloat64(slippageAmount, tonInfo.Decimals),
		JettonSymbol: jettonInfo.Symbol,
	}
	return preRequestInfo, nil
}


// продажа монет (Jetton -> TON)
func CellJetton(ctx context.Context, jettonCA string, amount float64, slippage int) error {
	// создание API клиента TON для tongo
	tongoClient, err := settings.GetTonClientTongoWithTimeout("mainnet", constants.TongoClientTimeout)
	if err != nil {
		return fmt.Errorf("send cell transaction: %w", err)
	}

	// получение данных о кошельке через tongo
	realWallet, err := myTongoWallet.GetWallet(tongoClient)
	if err != nil {
		return fmt.Errorf("send cell transaction: %w", err)
	}
	// получение данных о продаваемой монете
	jettonInfo, err := myStonfiJettons.GetJettonInfoByAddressWithTimeout(jettonCA, constants.GetJettonInfoByAddressTimeout)
	if err != nil {
		return fmt.Errorf("send cell transaction: %w", err)
	}
	// получение данных о TON
	tonInfo, err := myStonfiJettons.GetJettonInfoByAddressWithTimeout(constants.TonInfoAddr, constants.GetJettonInfoByAddressTimeout)
	if err != nil {
		return fmt.Errorf("send cell transaction: %w", err)
	}
	// цена монеты в TON
	jettonPriceInTON := jettonInfo.PriceUSD / tonInfo.PriceUSD

	
	// адрес получателя (StonfiRouter)
	jettonRouter := tongoTon.MustParseAccountID(constants.StonfiRouterAddr)
	// адрес монеты (откуда) jettonCA
	jettonMaster0 := tongoTon.MustParseAccountID(jettonCA)
	// адрес монеты (куда) TON
	jettonMaster1 := tongoTon.MustParseAccountID(constants.ProxyTonMasterAddr)
	
	// структура с информацией для Swap транзакции на DEX Stonfi
	stonfiStruct, err := tongoStonfi.NewStonfi(ctx, tongoClient, jettonRouter, jettonMaster0, jettonMaster1)
	if err != nil {
		apiErr := coreErrors.New(
			fmt.Errorf("send cell transaction: create new stonfiStruct: %w", err),
			"failed to prepare message",
			"ton_api",
			500,
		)
		apiErr.CheckTimeout()
		return apiErr
	}

	// TON для газовой комиссии (0.3 TON)
	gasToncoins := tongoTlb.Grams(300_000_000)
	// TON для передачи в следующее сообщение цепочки транзакций (0.2 TON)
	forwardToncoins := tongoTlb.Grams(200_000_000)
	// кол-во монет в виде *big.Int
	bigIntAmount := myTongoServices.ConvertJettonsAmountToBigInt(jettonInfo.Decimals, amount)
	// адрес отправителя (кошелёк юзера)
	senderAddrID := tongoTon.MustParseAccountID(settings.JsonWallet.Hash)

	// предположительное кол-во TON на выходе без учёта изменения цены и газовой комиссии
	predictedTonAmount := amount * jettonPriceInTON
	// перевод процента проскальзывания в часть от кол-ва TON в виде float64
	slippageAmount := predictedTonAmount * (1.0 - float64(slippage) / 100)
	// процент проскальзывания (часть от кол-ва TON) в виде *big.Int
	minOut := myTongoServices.ConvertJettonsAmountToBigInt(constants.TonDecimals, slippageAmount)

	// структура для совершения Swap транзакции
	jettonTransfer, err := stonfiStruct.MakeSwapMessage(gasToncoins, forwardToncoins, *bigIntAmount, *minOut, senderAddrID)
	if err != nil {
		return coreErrors.New(
			fmt.Errorf("send cell transaction: make swap message: %w", err),
			"failed to make swap message",
			"ton_api",
			500,
		)
	}

	// отправка сообщения в блокчейн
	err = realWallet.Send(ctx, jettonTransfer)
	if err != nil {
		apiErr := coreErrors.New(
			fmt.Errorf("send cell transaction: send transfer message: %w", err),
			"failed to send transfer message",
			"ton_api",
			500,
		)
		apiErr.CheckTimeout()
		return apiErr
	}
	return nil
}
