package transactions

// import (
// 	"fmt"
// 	"context"
// 	"reflect"

// 	"github.com/tonkeeper/tongo"
// 	tongoJettons "github.com/tonkeeper/tongo/contract/jetton"
// 	tongoTlb "github.com/tonkeeper/tongo/tlb"

// 	myTonapiJettons "github.com/Danil-114195722/HamstersShaver/ton_api_tonapi/jettons"
	
// 	myTongoWallet "github.com/Danil-114195722/HamstersShaver/ton_api_tongo/wallet"
// 	myTongoJettons "github.com/Danil-114195722/HamstersShaver/ton_api_tongo/jettons"
	
// 	"github.com/Danil-114195722/HamstersShaver/settings"
// )


// // type TransferMessage struct {
// // 	// адрес монеты и ещё что-то с блокчейном
// // 	Jetton              *Jetton
// // 	// адрес отправителя
// // 	Sender              ton.AccountID
// // 	// количество Jetton токенов, которые вы хотите обменять.
// // 	JettonAmount        *big.Int
// // 	// адрес получателя (может совпадать с адресом отправителя)
// // 	Destination         ton.AccountID
// // 	// необязательное поле, представляющее адрес, куда будет отправлен ответ после выполнения транзакции.
// // 	ResponseDestination *ton.AccountID
// // 	// Количество TON, которое вы прикрепляете к транзакции.
// // 	// Эти средства могут использоваться для покрытия газовых сборов транзакции.
// // 	// Это значение должно быть достаточно большим, чтобы покрыть все необходимые сборы, связанные с транзакцией.
// // 	AttachedTon         tlb.Grams
// // 	// количество TON, которое будет отправлено дальше к другому контракту или аккаунту, как часть сложной транзакции.
// // 	ForwardTonAmount    tlb.Grams
// // 	// может отсутствовать
// // 	ForwardPayload      *boc.Cell
// // 	// может отсутствовать
// // 	CustomPayload       *boc.Cell
// // }


// // продажа монет
// // TODO: 10 попыток до успеха (ошибка "error code: 651 message: cannot load block")
// func CellJetton(ctx context.Context, jettonCA string, jetton myTonapiJettons.AccountJetton, amount float64) error {
// 	// получение данных о кошельке через tongo
// 	realWallet, err := myTongoWallet.GetWallet()
// 	if err != nil {
// 		return err
// 	}
// 	// fmt.Println("realWallet.blockchain", realWallet.blockchain)
// 	// fmt.Println("\n")
// 	// val1 := reflect.ValueOf(realWallet)
// 	// if val1.Kind() == reflect.Ptr {
// 	// 	val1 = val1.Elem()
// 	// }
// 	// for i := 0; i < val1.NumField(); i++ {
// 	// 	fmt.Printf("realWallet.%v: %v\n", val1.Type().Field(i).Name, val1.Field(i).Interface())
// 	// }

// 	jettonStruct := myTongoJettons.GetJettonStruct(jettonCA)
// 	bigIntAmount := myTongoJettons.ConvertJettonsAmountToBigInt(jetton, amount)
// 	// адрес получателя (тот же, что и отправителя)
// 	recipientAddr := tongo.MustParseAddress(settings.JsonWallet.Hash)

// 	jettonTransfer := tongoJettons.TransferMessage{
// 		Jetton: jettonStruct,
// 		JettonAmount: bigIntAmount,
// 		Destination: recipientAddr.ID,
// 		AttachedTon: tongoTlb.Grams(300000000),  // 0.3 TON
// 		ForwardTonAmount: 0,
// 		// addition fields
// 		Sender: recipientAddr.ID,
// 		// Wallet address used to return remained TON-coins with excesses message.
// 		ResponseDestination: &recipientAddr.ID,
// 	}

// 	// Вывод значений структуры
// 	fmt.Println("\n")
// 	val := reflect.ValueOf(jettonTransfer)
// 	if val.Kind() == reflect.Ptr {
// 		val = val.Elem()
// 	}
// 	for i := 0; i < val.NumField(); i++ {
// 		fmt.Printf("jettonTransfer.%v: %v\n", val.Type().Field(i).Name, val.Field(i).Interface())
// 	}

// 	// отправка сообщения в блокчейн
// 	err = realWallet.Send(ctx, jettonTransfer)
// 	if err != nil {
// 		settings.ErrorLog.Printf("Failed to send transfer message: %v", err)
// 	}

// 	return err
// }


// // func getJettonAddressStruct(jettonCA string) (tonapi.AccountAddress, error) {
// // 	var jettonAddr tonapi.AccountAddress

// // 	// // Декодируем CA монеты из base64 в бинарные данные
// //     // binaryJettonCA, err := base64.URLEncoding.DecodeString(jettonCA)
// //     // if err != nil {
// //     //     settings.ErrorLog.Println("Failed to decode jettonCA from base64 to binary:", err.Error())
// //     //     return jettonAddr, err
// //     // }
// // 	// // перевод CA монеты из бинарных данных в HEX-вид
// // 	// hexJettonCA := "0:" + hex.EncodeToString(binaryJettonCA)

// // 	// создание структуры tonapi.AccountAddress и заполнение её CA монеты
// // 	jettonAddr = tonapi.AccountAddress{
// // 		Address: jettonCA,
// // 		Name: tonapi.OptString{Set: false},
// // 		IsScam: false,
// // 		Icon: tonapi.OptString{Set: false},
// // 		IsWallet: false,
// // 	}

// // 	return jettonAddr, nil
// // }

