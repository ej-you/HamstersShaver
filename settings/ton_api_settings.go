package settings

import (
	"os"
	"encoding/json"
	"context"

	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/ton"
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


// тестовый конфиг
// var tonConfigUrl string = "https://ton-blockchain.github.io/testnet-global.config.json"
// реальный конфиг
var tonConfigUrl string = "https://ton.org/global.config.json"

// создание клиента TON
func getTonClient() ton.APIClientWrapped {
	client := liteclient.NewConnectionPool()

	// подключение с URL-конфигом
	err := client.AddConnectionsFromConfigUrl(context.Background(), tonConfigUrl)
	DieIf(err)

	// API-клиент с полной проверкой
	api := ton.NewAPIClient(client, ton.ProofCheckPolicyFast).WithRetry()

	return api
}


// создание API клиента TON
var TonAPI ton.APIClientWrapped = getTonClient()
// данные кошелька из JSON-конфига
var JsonWallet JsonWalletData = getJsonWallet()
