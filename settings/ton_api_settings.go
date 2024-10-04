package settings

import (
	"os"
	"errors"
	"encoding/json"

	tongo "github.com/tonkeeper/tongo/liteapi"
	tonapi "github.com/tonkeeper/tonapi-go"
)


type JsonWalletData struct {
	Hash		string	`json:"hash"`
	SeedPhrase	string	`json:"seed_phrase"`
}

// путь до JSON-файла с данными кошелька
var pathToWalletData string = "./settings/config/wallet.json"

// получение данных о кошельке из JSON-файла
func getJsonWallet() JsonWalletData {
	var jsonWallet JsonWalletData

	// открытие файла
	fileData, err := os.ReadFile(pathToWalletData)
	DieIf(err)

	// перевод данных из JSON в структуру JsonWalletData
	err = json.Unmarshal(fileData, &jsonWallet)
	DieIf(err)

	return jsonWallet
}

// создание клиента TON для tonapi-go
func GetTonClientTonapi(conType string) (*tonapi.Client, error) {
	var client *tonapi.Client

	// тестовый конфиг
	if conType == "testnet" {
		client, err := tonapi.NewClient(tonapi.TestnetTonApiURL)
		if err != nil {
			return client, err
		}
		// InfoLog.Println("(tonapi) Connected to testnet TON node")
		return client, nil
	// основной конфиг
	} else if conType == "mainnet" {
		client, err := tonapi.New()
		if err != nil {
			return client, err
		}
		// InfoLog.Println("(tonapi) Connected to mainnet TON node")
		return client, nil
	// неправильный параметр конфига
	} else {
		conTypeError := errors.New("(tonapi) Invalid conType parameter was given")
		return client, conTypeError
	}
}

// создание клиента TON для tongo
func GetTonClientTongo(conType string) (*tongo.Client, error) {
	var client *tongo.Client

	// тестовый конфиг
	if conType == "testnet" {
		client, err := tongo.NewClientWithDefaultTestnet()
		if err != nil {
			return client, err
		}
		// InfoLog.Println("(tongo) Connected to testnet TON node")
		return client, nil
	// основной конфиг
	} else if conType == "mainnet" {
		client, err := tongo.NewClientWithDefaultMainnet()
		if err != nil {
			return client, err
		}
		// InfoLog.Println("(tongo) Connected to mainnet TON node")
		return client, nil
	// неправильный параметр конфига
	} else {
		conTypeError := errors.New("(tongo) Invalid conType parameter was given")
		return client, conTypeError
	}
}

// данные кошелька из JSON-конфига
var JsonWallet JsonWalletData = getJsonWallet()
