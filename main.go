package main

import (
	"fmt"
	"reflect"
	// "time"
	"context"

	"github.com/Danil-114195722/HamstersShaver/ton_api/wallet"
	// "github.com/Danil-114195722/HamstersShaver/ton_api/account"
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
	// ctxMain, ctxCancel := context.WithTimeout(context.Background(), 30*time.Second)
    // defer ctxCancel()

	walletInfo := wallet.GetWallet()

	// ПОКА ЭТА ФУНКЦИЯ НЕ РАБОТАЕТ
	balance, err := wallet.GetWalletBalanceFloat64(context.Background(), walletInfo)
	if err == nil {
		fmt.Println("balance: %f | balance + 1: %f", balance, balance + 1)
	}

	// acc := account.GetAccount(ctxMain)

	// fmt.Printf("State: %s\n", acc.State)
	// fmt.Printf("Data: %s\n", acc.Data)
	// fmt.Printf("Code: %s\n", acc.Code)
	// fmt.Printf("LastTxLT: %s\n", acc.LastTxLT)
	// fmt.Printf("LastTxHash: %s\n", acc.LastTxHash)

	// fmt.Printf("Balance: %s TON\n", acc.State.Balance)
	// fmt.Printf("ExtraCurrencies: %s\n", acc.State.ExtraCurrencies)
	// fmt.Printf("StorageInfo: %s\n", acc.State.StorageInfo)
	// fmt.Printf("AccountStorage: %s\n", acc.State.AccountStorage)

	// fmt.Printf("ExtraCurrencies: %s\n", acc.State.ExtraCurrencies.keySz)
	// fmt.Printf("ExtraCurrencies: %s\n", acc.State.ExtraCurrencies.root)

	// ShowStructure(acc.State.ExtraCurrencies)
	// ShowStructure(acc.State)
	// ShowStructure(acc.State.Balance)
	// ShowStructure(balance)
}
