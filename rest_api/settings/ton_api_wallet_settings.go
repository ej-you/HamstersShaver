package settings

import "sync"


type JsonWalletData struct {
	Hash		string	`json:"hash"`
	SeedPhrase	string	`json:"seed_phrase"`
}


var once sync.Once
// данные кошелька
var jsonWallet *JsonWalletData


// получение данных о кошельке из JSON-файла
func GetJsonWallet() *JsonWalletData {
	once.Do(func() {
		// создание структуры
		jsonWallet = &JsonWalletData{
			Hash: hash,
			SeedPhrase: seedPhrase,
		}
	})
	return jsonWallet
}
