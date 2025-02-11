package jettons

import (
	"fmt"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	coreErrors "github.com/ej-you/HamstersShaver/rest_api/core/errors"
)


const apiUrl = "https://api.ston.fi/v1/assets/"


// описывает любую монету (функция получения информации о монете по её адресу)
type JettonParams struct {
	Symbol 			string `json:"symbol" example:"GRAM" description:"символ монеты"`
	Decimals 		int `json:"decimals" example:"9" description:"decimals монеты"`
	MasterAddress 	string `json:"masterAddress" example:"EQC47093oX5Xhb0xuk2lCr2RhS8rj-vul61u4W2UH5ORmG_O" description:"мастер-адрес монеты (jetton_master)"`
	PriceUSD 		float64 `json:"priceUsd" example:"0.002695039717551585" description:"цена монеты в долларах (USD)"`
}

// для десериализации json'а с инфой о монете в структуру
type AssetJettonParams struct {
	Asset struct{
		Symbol string `json:"symbol"`
		Decimals int `json:"decimals"`
		// мастер-адрес монеты (jetton_master)
		MasterAddress string `json:"contract_address"`
		// цена в долларах
		StringPriceUSD string `json:"dex_price_usd"`
	}
}


// получение инфы о монете по её адресу с таймаутом
func GetJettonInfoByAddressWithTimeout(jettonAddr string, timeout time.Duration) (JettonParams, error) {
	var jettonInfo JettonParams
	var jettonInfoParse AssetJettonParams
	var apiErr error

	// запрос к API
	client := http.Client{Timeout: timeout}
	resp, err := client.Get(apiUrl + jettonAddr)
	if err != nil {
		urlErr := err.(*url.Error)
		// если ошибка таймаута
		if urlErr.Timeout() {
			apiErr = fmt.Errorf("get jetton data from Stonfi API: %w", coreErrors.TimeoutError)
		// неизвестная ошибка
		} else {
			apiErr = fmt.Errorf("get jetton data from Stonfi API: failed to send request: %v: %w", err, coreErrors.RestApiError)
		}
		return jettonInfo, apiErr
	}

	defer resp.Body.Close()
	// чтение байт тела ответа от API
	rawJsonBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		apiErr = fmt.Errorf("get jetton data from Stonfi API: failed to read data from response body: %w", coreErrors.RestApiError)
		return jettonInfo, apiErr
	}

	// десериализация JSON-ответа в структуру jsonPairs
	err = json.Unmarshal(rawJsonBytes, &jettonInfoParse)
	if err != nil {
		apiErr = fmt.Errorf("get jetton data from Stonfi API: unmarshal json: %v: %w", err, coreErrors.JettonNotFoundError)
		return jettonInfo, apiErr
	}

	// если не было найдено информации о жетоне
	if (jettonInfoParse == AssetJettonParams{}) {
		apiErr = fmt.Errorf("get jetton data from Stonfi API: %w", coreErrors.JettonNotFoundError)
		return jettonInfo, apiErr
	}

	// перевод цены монеты из строки во float64
	parsePriceToFloat, err := strconv.ParseFloat(jettonInfoParse.Asset.StringPriceUSD, 64)
	if err != nil {
		apiErr = fmt.Errorf("get jetton data from Stonfi API: failed to parse float from StringPriceUSD: %w", coreErrors.RestApiError)
		return jettonInfo, apiErr
	}

	// формирование выходной структуры с данными
	jettonInfo = JettonParams{
		Symbol: jettonInfoParse.Asset.Symbol,
		Decimals: jettonInfoParse.Asset.Decimals,
		MasterAddress: jettonInfoParse.Asset.MasterAddress,
		PriceUSD: parsePriceToFloat,
	}

	return jettonInfo, nil
}
