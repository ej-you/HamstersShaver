package main

import (
	"fmt"
	"context"

	// myCryptocompareJettons "github.com/Danil-114195722/HamstersShaver/ton_api/cryptocompare/jettons"
	// myDexscreenerJettons "github.com/Danil-114195722/HamstersShaver/ton_api/dexscreener/jettons"
	myTonapiAccount "github.com/Danil-114195722/HamstersShaver/ton_api/tonapi/account"
	myTonapiJettons "github.com/Danil-114195722/HamstersShaver/ton_api/tonapi/jettons"
	myTongoTransactions "github.com/Danil-114195722/HamstersShaver/ton_api/tongo/transactions"
)


func main() {
	tonBalance, _ := myTonapiAccount.GetBalanceTON(context.Background())
	fmt.Printf("tonBalance: %v TON | Balance: %d | Decimals: %d\n\n", tonBalance.BeautyBalance, tonBalance.Balance, tonBalance.Decimals)

	var cellNot myTonapiJettons.AccountJetton

	accountJettons, _ := myTonapiJettons.GetBalanceJettons(context.Background())
	for _, accJetton := range accountJettons {
		fmt.Printf("Symbol: %s | BeautyBalance: %s", accJetton.Symbol, accJetton.BeautyBalance)
		fmt.Printf(" | Balance: %d | Decimals: %d\n", accJetton.Balance, accJetton.Decimals)
		fmt.Printf("Master: %s\n\n", accJetton.MasterAddress)

		if accJetton.Symbol == "NOT" {
			cellNot = accJetton
		}
	}

	// процент проскальзывания (30%)
	slippage := 30

	jettonCA := "EQAvlWFDxGF2lXm67y4yzC17wYKD9A0guwPkMs1gOsM__NOT"
	err := myTongoTransactions.CellJetton(context.Background(), jettonCA, cellNot, 20.0, slippage)
	if err == nil {
		fmt.Println("GREAT!!!")
	}

	// pTON := "EQCM3B12QK1e4yZSf8GtBRT0aLMNyEsBc_DhVfRRtOEffLez"
	// NOT := "EQAvlWFDxGF2lXm67y4yzC17wYKD9A0guwPkMs1gOsM__NOT"
	// MEM := "EQAWpz2_G0NKxlG2VvgFbgZGPt8Y1qe0cGj-4Yw5BfmYR5iF"

	// notPool, err := myDexscreenerJettons.GetJettonsPoolInfo(pTON, NOT)
	// if err == nil {
	// 	fmt.Println("notPool:", notPool)
	// }

	// memPrice, err := myDexscreenerJettons.GetJettonPriceUSD(pTON, MEM)
	// if err == nil {
	// 	fmt.Println("MEM in USD:", memPrice)
	// }

	// tonPrice, err := myCryptocompareJettons.GetTonPriseUSD()
	// if err == nil {
	// 	fmt.Println("TON in USD:", tonPrice)
	// }
}
