package main

import (
	"fmt"
	"context"

	tonapi "github.com/tonkeeper/tonapi-go"

	"github.com/Danil-114195722/HamstersShaver/ton_api_tonapi/account"
	"github.com/Danil-114195722/HamstersShaver/ton_api_tonapi/jettons"
	"github.com/Danil-114195722/HamstersShaver/ton_api_tonapi/transactions"
)


func main() {
	tonBalance, _ := account.GetBalanceTON(context.Background())
	fmt.Printf("tonBalance: %v TON | Balance: %d\n", tonBalance.BeautyBalance, tonBalance.Balance)

	var cellNot jettons.AccountJetton

	accountJettons, _ := jettons.GetBalanceJettons(context.Background())
	for _, accJetton := range accountJettons {
		fmt.Printf("Symbol: %s | BeautyBalance: %s", accJetton.Symbol, accJetton.BeautyBalance)
		fmt.Printf(" | Balance: %d | Decimals: %d\n", accJetton.Balance, accJetton.Decimals)

		if accJetton.Symbol == "NOT" {
			cellNot = accJetton
		}
	}

	stonfi := tonapi.JettonSwapActionDexStonfi
	jettonCA := "EQAvlWFDxGF2lXm67y4yzC17wYKD9A0guwPkMs1gOsM__NOT"
	err := transactions.CellJetton(context.Background(), stonfi, jettonCA, cellNot, 10.0)
	if err == nil {
		fmt.Println("GREAT!!")
	}
}
