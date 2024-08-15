package main

import (
	"fmt"
	"context"

	myTonapiAccount "github.com/Danil-114195722/HamstersShaver/ton_api/tonapi/account"
	myTonapiJettons "github.com/Danil-114195722/HamstersShaver/ton_api/tonapi/jettons"
	// myTongoTransactions "github.com/Danil-114195722/HamstersShaver/ton_api/tongo/transactions"
)


func main() {
	tonBalance, _ := myTonapiAccount.GetBalanceTON(context.Background())
	fmt.Printf("tonBalance: %v TON | Balance: %d | Decimals: %d\n\n", tonBalance.BeautyBalance, tonBalance.Balance, tonBalance.Decimals)

	// var cellJetton myTonapiJettons.AccountJetton

	accountJettons, _ := myTonapiJettons.GetBalanceJettons(context.Background())
	for _, accJetton := range accountJettons {
		fmt.Printf("Symbol: %s | BeautyBalance: %s", accJetton.Symbol, accJetton.BeautyBalance)
		fmt.Printf(" | Balance: %d | Decimals: %d\n", accJetton.Balance, accJetton.Decimals)
		fmt.Printf("Master: %s\n\n", accJetton.MasterAddress)

		// if accJetton.Symbol == "USD₮" {
		// 	cellJetton = accJetton
		// }
	}

	// процент проскальзывания (20%)
	// slippage := 20

	// jettonCA := "EQAvlWFDxGF2lXm67y4yzC17wYKD9A0guwPkMs1gOsM__NOT"
	// jettonCA := "EQCxE6mUtQJKFnGfaROTKOt1lZbDiiX1kCixRv7Nw2Id_sDs"
	// err := myTongoTransactions.CellJetton(context.Background(), jettonCA, cellJetton, 0.1, slippage)
	// if err == nil {
	// 	fmt.Println("GREAT!!!")
	// }

	// GRAM
	jettonInfo, err := myTonapiJettons.GetJettonInfoByAddres(context.Background(), "EQC47093oX5Xhb0xuk2lCr2RhS8rj-vul61u4W2UH5ORmG_O")
	if err == nil {
		fmt.Println("jettonInfo:", jettonInfo)
	}
}
