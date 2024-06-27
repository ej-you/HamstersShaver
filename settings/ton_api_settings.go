package settings

import (
	"os"
	"encoding/json"
	"context"

	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/ton"
)


type JsonWallet struct {
	Hash		string	`json:"hash"`
	SeedPhrase	string	`json:"seed_phrase"`
}

// путь до JSON-файла с данными кошелька
var pathToWalletData string = "./settings/config/wallet.json"

// получение данных о кошельке из JSON-файла
func GetJsonWallet() JsonWallet {
	var jsonWallet JsonWallet

	// открытие файла
	fileData, err := os.ReadFile(pathToWalletData)
	DieIf(err)

	// перевод данных из JSON в структуру JsonWallet
	err = json.Unmarshal(fileData, &jsonWallet)
	DieIf(err)

	return jsonWallet
}


// тестовый конфиг
// var tonConfigUrl string = "https://ton-blockchain.github.io/testnet-global.config.json"
// реальный конфиг
var tonConfigUrl string = "https://ton.org/global.config.json"

// создание клиента TON
func GetTonClient() ton.APIClientWrapped {
	client := liteclient.NewConnectionPool()

	// получение конфига
	cfg, err := liteclient.GetConfigFromUrl(context.Background(), tonConfigUrl)
	DieIf(err)

	// подключение с полученным конфигом
	err = client.AddConnectionsFromConfig(context.Background(), cfg)
	DieIf(err)

	// API-клиент с полной проверкой
	api := ton.NewAPIClient(client, ton.ProofCheckPolicyFast).WithRetry()

	return api
}
