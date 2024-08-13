package main

import (
	"fmt"
	"context"

	tongoTon "github.com/tonkeeper/tongo/ton"

	myTonapiAccount "github.com/Danil-114195722/HamstersShaver/ton_api_tonapi/account"
	myTonapiJettons "github.com/Danil-114195722/HamstersShaver/ton_api_tonapi/jettons"
	// myTongoTransactions "github.com/Danil-114195722/HamstersShaver/ton_api_tongo/transactions"
	myTongoPool "github.com/Danil-114195722/HamstersShaver/ton_api_tongo/pool"
)


func main() {
	tonBalance, _ := myTonapiAccount.GetBalanceTON(context.Background())
	fmt.Printf("tonBalance: %v TON | Balance: %d\n", tonBalance.BeautyBalance, tonBalance.Balance)

	// var cellNot myTonapiJettons.AccountJetton

	accountJettons, _ := myTonapiJettons.GetBalanceJettons(context.Background())
	for _, accJetton := range accountJettons {
		fmt.Printf("Symbol: %s | BeautyBalance: %s", accJetton.Symbol, accJetton.BeautyBalance)
		fmt.Printf(" | Balance: %d | Decimals: %d\n", accJetton.Balance, accJetton.Decimals)
		fmt.Printf("Master: %s\n\n", accJetton.MasterAddress)

		// if accJetton.Symbol == "NOT" {
		// 	cellNot = accJetton
		// }
	}

	// // процент проскальзывания (30%)
	// slippage := 30

	// jettonCA := "EQAvlWFDxGF2lXm67y4yzC17wYKD9A0guwPkMs1gOsM__NOT"
	// err := myTongoTransactions.CellJetton(context.Background(), jettonCA, cellNot, 5.0, slippage)
	// if err == nil {
	// 	fmt.Println("GREAT!!!")
	// }

	// NOT
	jettonMaster0 := tongoTon.MustParseAccountID("EQAvlWFDxGF2lXm67y4yzC17wYKD9A0guwPkMs1gOsM__NOT")
	// TON
	// jettonMaster1 := tongoTon.MustParseAccountID("EQCM3B12QK1e4yZSf8GtBRT0aLMNyEsBc_DhVfRRtOEffLez")
	jettonMaster1 := tongoTon.MustParseAccountID("EQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAM9c")

	poolAddr, err := myTongoPool.GetJettonsPoolAddress(context.Background(), jettonMaster0, jettonMaster1)
	if err == nil {
		fmt.Println("poolAddr:", poolAddr)
	}
}
