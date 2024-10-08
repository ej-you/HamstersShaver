package transactions

import (
	"context"
	"errors"
	"fmt"
	"time"

	tongoAbi "github.com/tonkeeper/tongo/abi"
	tongoBoc "github.com/tonkeeper/tongo/boc"
	tongoTlb "github.com/tonkeeper/tongo/tlb"
	tongoTon "github.com/tonkeeper/tongo/ton"
	tongoJetton "github.com/tonkeeper/tongo/contract/jetton"

	myStonfiJettons "github.com/Danil-114195722/HamstersShaver/ton_api/stonfi/jettons"
	
	myTongoWallet "github.com/Danil-114195722/HamstersShaver/ton_api/tongo/wallet"
	myTongoServices "github.com/Danil-114195722/HamstersShaver/ton_api/tongo/services"

	"github.com/Danil-114195722/HamstersShaver/settings/constants"
	"github.com/Danil-114195722/HamstersShaver/settings"
)


// данные о последующей транзакции покупки монет (TON -> Jetton)
type PreRequestBuyJetton struct {
	UsedTON 		float64 `json:"usedTon"`
	JettonCA 		string `json:"jettonCA"`
	DEX 			string `json:"dex"`
	JettonsOut 		float64 `json:"jettonsOut"`
	MinOut	 		float64 `json:"minOut"`
	JettonSymbol 	string `json:"jettonSymbol"`
}


// получение данных на подтверждение последующей транзакции покупки монет (TON -> Jetton)
func GetPreRequestBuyJetton(jettonCA string, tonAmount float64, slippage int, timeout time.Duration) (PreRequestBuyJetton, error) {
	var preRequestInfo PreRequestBuyJetton

	// получение данных о покупаемой монете
	jettonInfo, err := myStonfiJettons.GetJettonInfoByAddressWithTimeout(jettonCA, timeout)
	if err != nil {
		return preRequestInfo, err
	}
	// получение данных о TON
	tonInfo, err := myStonfiJettons.GetJettonInfoByAddressWithTimeout(constants.TonInfoAddr, timeout)
	if err != nil {
		return preRequestInfo, err
	}
	// цена монеты в TON
	jettonPriceInTON := jettonInfo.PriceUSD / tonInfo.PriceUSD

	// предположительное кол-во монет на выходе без учёта изменения цены
	predictedJettonsAmount := tonAmount / jettonPriceInTON
	// перевод процента проскальзывания в часть от кол-ва монет в виде float64
	slippageAmount := predictedJettonsAmount * (1.0 - float64(slippage) / 100)

	preRequestInfo = PreRequestBuyJetton{
		UsedTON: tonAmount,
		JettonCA: jettonCA,
		DEX: "Stonfi",
		JettonsOut: predictedJettonsAmount,
		MinOut: slippageAmount,
		JettonSymbol: jettonInfo.Symbol,
	}
	return preRequestInfo, nil
}


// покупка монет (TON -> Jetton)
func BuyJetton(ctx context.Context, timeout time.Duration, jettonCA string, amount float64, slippage int) error {
	// создание API клиента TON для tongo с таймаутом timeout
	tongoClient, err := settings.GetTonClientTongoWithTimeout("mainnet", timeout)
	if err != nil {
		return err
	}

	// получение данных о кошельке через tongo
	realWallet, err := myTongoWallet.GetWallet(tongoClient)
	if err != nil {
		return err
	}
	// получение данных о покупаемой монете с таймаутом timeout
	jettonInfo, err := myStonfiJettons.GetJettonInfoByAddressWithTimeout(jettonCA, timeout)
	if err != nil {
		return err
	}
	// получение данных о TON с таймаутом timeout
	tonInfo, err := myStonfiJettons.GetJettonInfoByAddressWithTimeout(constants.TonInfoAddr, timeout)
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
	jettonToBuy := tongoJetton.New(jettonMaster1, tongoClient)
	routersJettonWallet, err := jettonToBuy.GetJettonWallet(ctx, jettonRouter)
	if err != nil {
		getJettonWalletError := errors.New(fmt.Sprintf("Failed to get jetton wallet: %s", err.Error()))
		return getJettonWalletError
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
		sendMEssageError := errors.New(fmt.Sprintf("Failed to send transfer message: %s", err.Error()))
		return sendMEssageError
	}
	return nil
}

// ЗАТЕМ ВЫДАВАТЬ БАЛАНС TON НА АККАУНТЕ И ПОКУПАЕМОЙ/ПРОДАВАЕМОЙ МОНЕТЫ (КОГДА ОНИ ОБА ПОМЕНЯЮТСЯ, НО С ТАЙМАУТОМ 2min)
// (ПРОТЕСТИТЬ СИТУАЦИЮ НЕПРОХОЖДЕНИЯ ПОРОГА minOut, T.K. ТАМ НЕ ПОМЕНЯЕТСЯ БАЛАНС ПОКУПАЕМОЙ/ПРОДАВАЕМОЙ МОНЕТЫ)

// СХЕМУ (НИЖЕ) ПРОВОДИТЬ В САМОМ КЛИЕНТЕ ПОСЛЕ ОТПРАВКИ ЗАПРОСА В API НА ОТПРАВКУ ТРАНЗЫ
// Вначале попробовать сравнивать первое полученное значение баланса TON (т.е. когда нужное кол-во TON уйдёт на транзу)
// с последующими до случая с их расхождением. Затем брать баланс покупаемой/продаваемой монеты
// (вначале так же получив первое значение баланса вместе с первым значением баланса TON) и выдавать юзеру успех и
// новый баланс монеты (если он изменился) или неуспех (если он не изменился)
