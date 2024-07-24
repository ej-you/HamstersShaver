package transactions

import (
	"context"

	// "github.com/Danil-114195722/HamstersShaver/ton_api_tonapi/account"
	"github.com/Danil-114195722/HamstersShaver/ton_api_tonapi/jettons"
	// "github.com/Danil-114195722/HamstersShaver/settings"

)


type TransferMessage struct {
	// адрес монеты
	Jetton              *Jetton
	// адрес отправителя
	Sender              ton.AccountID
	// количество Jetton токенов, которые вы хотите обменять.
	JettonAmount        *big.Int
	// адрес получателя (может совпадать с адресом отправителя)
	Destination         ton.AccountID
	// необязательное поле, представляющее адрес, куда будет отправлен ответ после выполнения транзакции.
	ResponseDestination *ton.AccountID
	// Количество TON, которое вы прикрепляете к транзакции.
	// Эти средства могут использоваться для покрытия газовых сборов транзакции.
	// Это значение должно быть достаточно большим, чтобы покрыть все необходимые сборы, связанные с транзакцией.
	AttachedTon         tlb.Grams
	// количество TON, которое будет отправлено дальше к другому контракту или аккаунту, как часть сложной транзакции.
	ForwardTonAmount    tlb.Grams
	// может отсутствовать
	ForwardPayload      *boc.Cell
	// может отсутствовать
	CustomPayload       *boc.Cell
}


// продажа монет
func CellJetton(ctx context.Context, jettonCA string, jetton jettons.AccountJetton, amount float64) error {
	// количество монет на продажу в минимальных единицах монеты (в виде строки)
	// realAmount := strconv.Itoa(int(amount * math.Pow10(jetton.Decimals)))

	return nil
}


// func getJettonAddressStruct(jettonCA string) (tonapi.AccountAddress, error) {
// 	var jettonAddr tonapi.AccountAddress

// 	// // Декодируем CA монеты из base64 в бинарные данные
//     // binaryJettonCA, err := base64.URLEncoding.DecodeString(jettonCA)
//     // if err != nil {
//     //     settings.ErrorLog.Println("Failed to decode jettonCA from base64 to binary:", err.Error())
//     //     return jettonAddr, err
//     // }
// 	// // перевод CA монеты из бинарных данных в HEX-вид
// 	// hexJettonCA := "0:" + hex.EncodeToString(binaryJettonCA)

// 	// создание структуры tonapi.AccountAddress и заполнение её CA монеты
// 	jettonAddr = tonapi.AccountAddress{
// 		Address: jettonCA,
// 		Name: tonapi.OptString{Set: false},
// 		IsScam: false,
// 		Icon: tonapi.OptString{Set: false},
// 		IsWallet: false,
// 	}

// 	return jettonAddr, nil
// }

