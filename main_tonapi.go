package main

import (
	"fmt"
	"context"
	// "time"

	"github.com/Danil-114195722/HamstersShaver/ton_api_tonapi/account"
)


func main() {
	// ctxMain, cancel := context.WithTimeout(context.Background(), 60*time.Second)
    // defer cancel()

	// walletInfo := wallet.GetWallet()

	// balance, err := wallet.GetWalletBalance(context.Background(), walletInfo)
	// if err == nil {
	// 	fmt.Printf("balance: %v | balance + 1: %v\n", balance, balance + 1)
	// }

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
}
