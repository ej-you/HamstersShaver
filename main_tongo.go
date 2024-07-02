package main

import (
	"fmt"
	"reflect"
	"context"
	"time"

	"github.com/Danil-114195722/HamstersShaver/ton_api_tongo/wallet"
	"github.com/Danil-114195722/HamstersShaver/ton_api_tongo/account"
)


// временно, для просмотра полей структуры
func ShowStructure(s interface{}) {
	fmt.Println()
    for x := 0; x < reflect.ValueOf(s).Elem().NumField(); x++ {
        fmt.Printf("Name field: `%s`  Type: `%s`\n", reflect.TypeOf(s).Elem().Field(x).Name,
            reflect.ValueOf(s).Elem().Field(x).Type())
    }
}

func main() {
	ctxMain, cancel := context.WithTimeout(context.Background(), 60*time.Second)
    defer cancel()

	walletInfo := wallet.GetWallet()

	balance, err := wallet.GetWalletBalance(context.Background(), walletInfo)
	if err == nil {
		fmt.Printf("balance: %v | balance + 1: %v\n", balance, balance + 1)
	}

	accountState, err := account.GetAccountState(ctxMain)
	if err == nil {
		fmt.Println("accountState:", accountState)
	}

	// ShowStructure(accountState)
}
