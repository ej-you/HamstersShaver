package main

import (
	"fmt"
	"context"

	myTonapiAccount "github.com/Danil-114195722/HamstersShaver/ton_api_tonapi/account"
	myTonapiJettons "github.com/Danil-114195722/HamstersShaver/ton_api_tonapi/jettons"
	myTongoTransactions "github.com/Danil-114195722/HamstersShaver/ton_api_tongo/transactions"
)


const ProxyTonAddress = "EQARULUYsmJq1RiZ-YiH-IJLcAZUVkVff-KBPwEmmaQGH6aC"
const MemeAddress = "EQBDsNSrXfMwjUmiyf39gIh8flwVjkNLEmILoG-st6v6uTUi"
const NotAddress = "EQAN5ylxMPuAzArYh4iSM-Z_dvakIrtOctgjdzHD_0-YkVZs"


func main() {
	tonBalance, _ := myTonapiAccount.GetBalanceTON(context.Background())
	fmt.Printf("tonBalance: %v TON | Balance: %d\n", tonBalance.BeautyBalance, tonBalance.Balance)

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

	// процент проскальзывания (10%)
	slippage := 30

	jettonCA := "EQAvlWFDxGF2lXm67y4yzC17wYKD9A0guwPkMs1gOsM__NOT"
	err := myTongoTransactions.CellJetton(context.Background(), jettonCA, cellNot, 5.0, slippage)
	if err == nil {
		fmt.Println("GREAT!!!")
	}
}
