package main

import (
	"fmt"
	"context"

	"github.com/Danil-114195722/HamstersShaver/ton_api_tonapi/account"
	"github.com/Danil-114195722/HamstersShaver/ton_api_tonapi/jettons"
)


func main() {
	acc, err := account.GetAccount(context.Background())
	if err == nil {
		// fmt.Println("account:", acc)
	
		// fmt.Println("Address:", acc.Address)
		// fmt.Println("Balance:", acc.Balance)
		// fmt.Println("Name:", acc.Name)
		// fmt.Println("GetMethods:", acc.GetMethods)
		// fmt.Println("IsWallet:", acc.IsWallet)

		fmt.Printf("Status: %v | type: %T\n", acc.Status, acc.Status)
	}

	tonBalance, err := account.GetBalanceTON(context.Background())
	if err == nil {
		fmt.Printf("tonBalance: %v TON\n", tonBalance)
	}

	accountJettons, err := jettons.GetBalanceJettons(context.Background())
	if err == nil {
		for _, accJetton := range accountJettons {
			fmt.Printf("Symbol: %s | Balance: %f\n", accJetton.Symbol, accJetton.Balance)
		}
	}
}
