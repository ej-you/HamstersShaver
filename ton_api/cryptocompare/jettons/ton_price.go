package jettons

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/Danil-114195722/HamstersShaver/settings"
)


type jsonTonPrice struct {
	USD float64 `json:"USD"`
}


// получение стоимости TONcoin в долларах
func GetTonPriseUSD() (float64, error) {
	var jsonResp jsonTonPrice
	var tonPrice float64

	// создание URL для запроса к API
	url := "https://min-api.cryptocompare.com/data/price?fsym=TONcoin&tsyms=USD"
	urlWithApiKey := fmt.Sprintf("%s&api_key=%s", url, settings.CryptocompareApiKey)

	// запрос к API
	resp, err := http.Get(urlWithApiKey)
	if err != nil {
		settings.ErrorLog.Println("Failed to get TON price from cryptocompare API:", err)
		return tonPrice, err
	}

	defer resp.Body.Close()
	// чтение байт тела ответа от API
	rawJsonBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		settings.ErrorLog.Println("Failed to read data from response body:", err)
		return tonPrice, err
	}

	// десериализация JSON-ответа в структуру jsonTonPrice
	err = json.Unmarshal(rawJsonBytes, &jsonResp)
	if err != nil {
		settings.ErrorLog.Println("Failed to unmarshal JSON-data from response body:", err)
		return tonPrice, err
	}

	// если при парсинге не получилось вытащить значение "USD"
	if tonPrice == jsonResp.USD {
		valueNofFoundError := errors.New("USD value no found")
		settings.ErrorLog.Println("Failed to parse USD value from JSON-response:", valueNofFoundError)
		return tonPrice, valueNofFoundError
	}

	// возвращаем стоимость TONcoin в USD
	tonPrice = jsonResp.USD
	return tonPrice, nil
}
