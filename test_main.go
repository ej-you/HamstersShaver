package main

import (
	"fmt"
	"context"
	"time"

	myTonapiAccount "github.com/Danil-114195722/HamstersShaver/ton_api/tonapi/account"
	
	myTongoWallet "github.com/Danil-114195722/HamstersShaver/ton_api/tongo/wallet"
	myTongoTransactions "github.com/Danil-114195722/HamstersShaver/ton_api/tongo/transactions"
	
	"github.com/Danil-114195722/HamstersShaver/settings"
)


func main() {
	// создание API клиента TON для tonapi-go с таймаутом в 3 секунд
	tonapiClient, err := settings.GetTonClientTonapiWithTimeout("mainnet", 3*time.Second)
	if err != nil {
		panic(err)
	}
	// создание API клиента TON для tongo с таймаутом в 3 секунд
	tongoClient, err := settings.GetTonClientTongoWithTimeout("mainnet", 3*time.Second)
	if err != nil {
		panic(err)
	}

	// получение баланса TON на аккаунте
	tonBalance, _ := myTonapiAccount.GetBalanceTON(context.Background(), tonapiClient)
	fmt.Printf("tonBalance: %v TON | Balance: %d | Decimals: %d\n\n", tonBalance.BeautyBalance, tonBalance.Balance, tonBalance.Decimals)

	// получение балансов монет на аккаунте
	accountJettons, _ := myTonapiAccount.GetBalanceJettons(context.Background(), tonapiClient)
	for _, accJetton := range accountJettons {
		fmt.Printf("Symbol: %s | BeautyBalance: %s", accJetton.Symbol, accJetton.BeautyBalance)
		fmt.Printf(" | Balance: %d | Decimals: %d\n", accJetton.Balance, accJetton.Decimals)
		fmt.Printf("MasterAddress: %s\n", accJetton.MasterAddress)
	}

	// продажа GRAM
	jettonCA := "EQC47093oX5Xhb0xuk2lCr2RhS8rj-vul61u4W2UH5ORmG_O"

	preRequestCellInfo, err := myTongoTransactions.GetPreRequestCellJetton(jettonCA, 100, 20, 3*time.Second)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("\npreRequestCellInfo:", preRequestCellInfo)

	err = myTongoTransactions.CellJetton(context.Background(), 10*time.Second, jettonCA, 100, 20)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("\nCell transaction was sent successfully!!!")


	// покупка DOGS
	jettonCA = "EQCvxJy4eG8hyHBFsZ7eePxrRsUQSFE_jpptRAYBmcG_DOGS"

	preRequestBuyInfo, err := myTongoTransactions.GetPreRequestBuyJetton(jettonCA, 0.1, 20, 3*time.Second)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("\npreRequestBuyInfo:", preRequestBuyInfo)

	// err = myTongoTransactions.BuyJetton(context.Background(), 10*time.Second, jettonCA, 0.1, 20)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println("\Buy transaction was sent successfully!!!")


	// получение данных о кошельке через tongo
	realWallet, err := myTongoWallet.GetWallet(tongoClient)
	if err != nil {
		fmt.Println("Error (while getting wallet):", err)
		return
	}
	// получение Seqno
	seqno, err := myTonapiAccount.GetAccountSeqno(context.Background(), tonapiClient, realWallet)
	if err != nil {
		fmt.Println("Error (while getting seqno):", err)
		return
	}
	fmt.Println("\nSeqno: ", seqno)


	// //  получение баланса монеты на аккаунте
	// accountJettonInfo, err := myTonapiAccount.GetAccountJetton(context.Background(), tonapiClient, "EQC47093oX5Xhb0xuk2lCr2RhS8rj-vul61u4W2UH5ORmG_O")
	// if err != nil {
	// 	fmt.Printf("\n\nERROR! %v\n", err)
	// 	return
	// }
	// fmt.Printf("\n\nSymbol: %s | BeautyBalance: %s", accountJettonInfo.Symbol, accountJettonInfo.BeautyBalance)
	// fmt.Printf(" | Balance: %d | Decimals: %d\n", accountJettonInfo.Balance, accountJettonInfo.Decimals)
	// fmt.Printf("MasterAddress: %s\n", accountJettonInfo.MasterAddress)
}
