package main

import (
	"fmt"
	"context"

	myTonapiAccount "github.com/Danil-114195722/HamstersShaver/ton_api/tonapi/account"
	myTonapiJettons "github.com/Danil-114195722/HamstersShaver/ton_api/tonapi/jettons"
	myTongoTransactions "github.com/Danil-114195722/HamstersShaver/ton_api/tongo/transactions"
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

		// if accJetton.Symbol == "GRAM" {
		// 	cellJetton = accJetton
		// }
	}

	// процент проскальзывания (20%)
	slippage := 20

	// продажа GRAM
	// jettonCA := "EQC47093oX5Xhb0xuk2lCr2RhS8rj-vul61u4W2UH5ORmG_O"
	// err := myTongoTransactions.CellJetton(context.Background(), jettonCA, cellJetton, 100, slippage)
	// if err == nil {
	// 	fmt.Println("GREAT!!!")
	// }

	// покупка DOGS
	// LP GRAM-pTON:  EQASBZLwa2vfdsgoDF2w96pdccBJJRxDNXXPUL7NMm0WdnMx
	// jettonCA := "EQC47093oX5Xhb0xuk2lCr2RhS8rj-vul61u4W2UH5ORmG_O"
	jettonCA := "EQCvxJy4eG8hyHBFsZ7eePxrRsUQSFE_jpptRAYBmcG_DOGS"
	// err := myTongoTransactions.BuyJetton(context.Background(), jettonCA, 130, slippage)
	err := myTongoTransactions.BuyJetton(context.Background(), jettonCA, 0.1, slippage)
	if err == nil {
		fmt.Println("GREAT!!!")
	}
}
