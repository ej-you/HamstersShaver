package transactions

import (
	"context"
	"encoding/base64"
	// "encoding/hex"
	"fmt"
	"math"
	"strconv"

	tonapi "github.com/tonkeeper/tonapi-go"

	"github.com/Danil-114195722/HamstersShaver/ton_api_tonapi/account"
	"github.com/Danil-114195722/HamstersShaver/ton_api_tonapi/jettons"
	"github.com/Danil-114195722/HamstersShaver/settings"
)

// JettonSwapAction
// JettonTransferAction
// SmartContractAction
// Send
// !!! Action
// Event

// SendBlockchainMessage
// SendBlockchainMessageReq
// Boc


const stonfi tonapi.JettonSwapActionDex = tonapi.JettonSwapActionDexStonfi
const dedust tonapi.JettonSwapActionDex = tonapi.JettonSwapActionDexDedust


// продажа монет
func CellJetton(ctx context.Context, dex tonapi.JettonSwapActionDex, jettonCA string, jetton jettons.AccountJetton, amount float64) error {
	// количество монет на продажу в минимальных единицах монеты (в виде строки)
	realAmount := strconv.Itoa(int(amount * math.Pow10(jetton.Decimals)))

	// получаем структуру AccountAddress по данным аккаунта
	accountAddrStruct, err := account.GetAccountAddressStruct(ctx)
	if err != nil {
		return err
	}
	// получаем структуру AccountAddress по CA монеты
	jettonAddrStruct, err := getJettonAddressStruct(jettonCA)
	if err != nil {
		return err
	}

	// формируем структуру для проведения транзакции
	var jettonSwap tonapi.JettonSwapAction = tonapi.JettonSwapAction{
		Dex: dex,
		AmountIn: "0",
		AmountOut: realAmount,
		TonIn: tonapi.OptInt64{Set: false},
		TonOut: tonapi.OptInt64{Set: false},
		UserWallet: accountAddrStruct,
		Router: jettonAddrStruct,
		JettonMasterIn: tonapi.OptJettonPreview{Set: false},
		JettonMasterOut: tonapi.OptJettonPreview{Set: false},
	}

	fmt.Printf("\njettonSwap: %v\n", jettonSwap)

	// переводим структуру для проведения транзакции в JSON-формат
	jsonJettonSwap, err := jettonSwap.MarshalJSON()
	if err != nil {
		settings.ErrorLog.Println("Failed to marshal JettonSwapAction struct to JSON:", err.Error())
		return err
	}

	// переменные для создания нового сообщения в блокчейн
	var boc tonapi.OptString = tonapi.OptString{
		Set: false,
		// Value: hex.EncodeToString(jsonJettonSwap),
		Value: base64.StdEncoding.EncodeToString(jsonJettonSwap),
	}
	var batch = []string{boc.Value}

	// структура с новым сообщением в блокчейн
	var blockchainMessage tonapi.SendBlockchainMessageReq = tonapi.SendBlockchainMessageReq{
		Boc: boc,
		Batch: batch,
	}

	fmt.Println("\nblockchainMessage:", blockchainMessage, "\n")

	// отправка нового сообщения в блокчейн
	err = settings.TonapiTonAPI.SendBlockchainMessage(ctx, &blockchainMessage)
	if err != nil {
		settings.ErrorLog.Println("Failed to send BlockchainMessage:", err.Error())
		return err
	}

	return nil
}


// покупка моонет
func BuyJetton() {

}


func getJettonAddressStruct(jettonCA string) (tonapi.AccountAddress, error) {
	var jettonAddr tonapi.AccountAddress

	// // Декодируем CA монеты из base64 в бинарные данные
    // binaryJettonCA, err := base64.URLEncoding.DecodeString(jettonCA)
    // if err != nil {
    //     settings.ErrorLog.Println("Failed to decode jettonCA from base64 to binary:", err.Error())
    //     return jettonAddr, err
    // }
	// // перевод CA монеты из бинарных данных в HEX-вид
	// hexJettonCA := "0:" + hex.EncodeToString(binaryJettonCA)

	// создание структуры tonapi.AccountAddress и заполнение её CA монеты
	jettonAddr = tonapi.AccountAddress{
		Address: jettonCA,
		Name: tonapi.OptString{Set: false},
		IsScam: false,
		Icon: tonapi.OptString{Set: false},
		IsWallet: false,
	}

	return jettonAddr, nil
}
