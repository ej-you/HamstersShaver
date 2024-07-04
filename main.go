package main

import (
	"fmt"
	"context"

	"github.com/Danil-114195722/HamstersShaver/ton_api_tonapi/account"
	"github.com/Danil-114195722/HamstersShaver/ton_api_tonapi/jettons"
)


func main() {
	tonBalance, err := account.GetBalanceTON(context.Background())
	if err == nil {
		fmt.Printf("tonBalance: %v TON | Balance: %d\n", tonBalance.BeautyBalance, tonBalance.Balance)
	}

	accountJettons, err := jettons.GetBalanceJettons(context.Background())
	if err == nil {
		for _, accJetton := range accountJettons {
			fmt.Printf("Symbol: %s | BeautyBalance: %s", accJetton.Symbol, accJetton.BeautyBalance)
			fmt.Printf(" | Balance: %d | Decimals: %d\n", accJetton.Balance, accJetton.Decimals)
		}
	}
}
