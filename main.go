package main

import (
	"fmt"
	"reflect"

	// "github.com/Danil-114195722/HamstersShaver/ton_api/wallet"
	"github.com/Danil-114195722/HamstersShaver/ton_api/account"
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
	// walletInfo := wallet.GetWallet()

	// balance, err := wallet.GetWalletBalance(walletInfo)
	// if err != nil {
	// 	fmt.Println("ERROR")
	// } else {
	// 	fmt.Println("balance:", balance)
	// }

	acc := account.GetAccount()

	// fmt.Printf("State: %s\n", acc.State)
	// fmt.Printf("Data: %s\n", acc.Data)
	// fmt.Printf("Code: %s\n", acc.Code)
	// fmt.Printf("LastTxLT: %s\n", acc.LastTxLT)
	// fmt.Printf("LastTxHash: %s\n", acc.LastTxHash)

	fmt.Printf("Balance: %s TON\n", acc.State.Balance)
	fmt.Printf("ExtraCurrencies: %s\n", acc.State.ExtraCurrencies)
	fmt.Printf("StorageInfo: %s\n", acc.State.StorageInfo)
	fmt.Printf("AccountStorage: %s\n", acc.State.AccountStorage)

	// fmt.Printf("ExtraCurrencies: %s\n", acc.State.ExtraCurrencies.keySz)
	// fmt.Printf("ExtraCurrencies: %s\n", acc.State.ExtraCurrencies.root)

	ShowStructure(acc.State.ExtraCurrencies)
	ShowStructure(acc.State)
}
