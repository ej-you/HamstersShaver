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
func getTonClientTonapi(conType string) *tonapi.Client {
	var client *tonapi.Client

	// тестовый конфиг
	if conType == "testnet" {
		client, err := tonapi.NewClient(tonapi.TestnetTonApiURL)
		DieIf(err)
		InfoLog.Println("(tonapi) Connected to testnet TON node")
		return client
	// основной конфиг
	} else if conType == "mainnet" {
		client, err := tonapi.New()
		DieIf(err)
		InfoLog.Println("(tonapi) Connected to mainnet TON node")
		return client
	// неправильный параметр конфига
	} else {
		conTypeError := errors.New("(tonapi) Invalid conType parameter was given")
		DieIf(conTypeError)
		return client
	}
}

// создание клиента TON для tongo
func getTonClientTongo(conType string) *tongo.Client {
	var client *tongo.Client

	// тестовый конфиг
	if conType == "testnet" {
		client, err := tongo.NewClientWithDefaultTestnet()
		DieIf(err)
		InfoLog.Println("(tongo) Connected to testnet TON node")
		return client
	// основной конфиг
	} else if conType == "mainnet" {
		client, err := tongo.NewClientWithDefaultMainnet()
		DieIf(err)
		InfoLog.Println("(tongo) Connected to mainnet TON node")
		return client
	// неправильный параметр конфига
	} else {
		conTypeError := errors.New("(tongo) Invalid conType parameter was given")
		DieIf(conTypeError)
		return client
	}
}

// данные кошелька из JSON-конфига
var JsonWallet JsonWalletData = getJsonWallet()

// // создание API клиента TON для tonapi-go
// var TonapiTonAPI *tonapi.Client = getTonClientTonapi("mainnet")

// // создание API клиента TON для tongo
// var TongoTonAPI *tongo.Client = getTonClientTongo("mainnet")
