package jettons

import (
	"fmt"
	"encoding/json"
	"io"
	"net/http"
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


// структура для получения результата от горутины в функции с таймаутом
type jettonInfoByAddress struct {
	JettonInfo 	JettonParams
	Error 		error
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


// получение инфы о монете по её адресу
func getJettonInfoByAddress(jettonAddr string) jettonInfoByAddress {
	var jettonInfoParse AssetJettonParams
	var apiErr coreErrors.APIError

	// запрос к API
	resp, err := http.Get(apiUrl + jettonAddr)
	if err != nil {
		apiErr = coreErrors.New(
			fmt.Errorf("get jetton data from Stonfi API: %w", err),
			"failed to get jetton data",
			"stonfi_api",
			500,
		)
		return jettonInfoByAddress{JettonInfo: JettonParams{}, Error: apiErr}
	}

	defer resp.Body.Close()
	// чтение байт тела ответа от API
	rawJsonBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		apiErr = coreErrors.New(
			fmt.Errorf("get jetton data from Stonfi API: read data from response body: %w", err),
			"failed to get jetton data",
			"stonfi_api",
			500,
		)
		return jettonInfoByAddress{JettonInfo: JettonParams{}, Error: apiErr}
	}

	// десериализация JSON-ответа в структуру jsonPairs
	err = json.Unmarshal(rawJsonBytes, &jettonInfoParse)
	if err != nil {
		apiErr = coreErrors.New(
			fmt.Errorf("get jetton data from Stonfi API: jetton was not found: unmarshal json: %w", err),
			"jetton was not found",
			"stonfi_api",
			400,
		)
		return jettonInfoByAddress{JettonInfo: JettonParams{}, Error: apiErr}
	}

	// если не было найдено информации о жетоне
	if (jettonInfoParse == AssetJettonParams{}) {
		apiErr = coreErrors.New(
			fmt.Errorf("get jetton data from Stonfi API: jetton was not found"),
			"jetton was not found",
			"stonfi_api",
			400,
		)
		return jettonInfoByAddress{JettonInfo: JettonParams{}, Error: apiErr}
	}

	// перевод цены монеты из строки во float64
	parsePriceToFloat, err := strconv.ParseFloat(jettonInfoParse.Asset.StringPriceUSD, 64)
	if err != nil {
		apiErr = coreErrors.New(
			fmt.Errorf("get jetton data from Stonfi API: parse float from StringPriceUSD: %w", err),
			"failed to get jetton data",
			"rest_api",
			500,
		)
		return jettonInfoByAddress{JettonInfo: JettonParams{}, Error: apiErr}
	}

	// формирование выходной структуры с данными
	jettonInfo := JettonParams{
		Symbol: jettonInfoParse.Asset.Symbol,
		Decimals: jettonInfoParse.Asset.Decimals,
		MasterAddress: jettonInfoParse.Asset.MasterAddress,
		PriceUSD: parsePriceToFloat,
	}

	return jettonInfoByAddress{JettonInfo: jettonInfo, Error: nil}
}

// получение инфы о монете по её адресу с таймаутом
func GetJettonInfoByAddressWithTimeout(jettonAddr string, timeout time.Duration) (JettonParams, error) {
	// если таймаут равен 0
	if timeout == 0 {
		result := getJettonInfoByAddress(jettonAddr)
		return result.JettonInfo, result.Error
	}

	// создание небуферизированного канала
	ch := make(chan jettonInfoByAddress)
	// вызов горутины
	go func() {
		ch <- getJettonInfoByAddress(jettonAddr)
	}()

	select {
		// если данные получены, то возвращаем их
		case result := <- ch:
			return result.JettonInfo, result.Error
		// если прошло время timeout, а данные не получены, то возвращаем ошибку таймаута
		case <- time.After(timeout):
			timeoutAPIError := coreErrors.NewTimeout(
				fmt.Errorf("get jetton data from Stonfi API: timeout error"),
				"get jetton data via stonfi: timeout error",
				"timeout",
				500,
			)
			return JettonParams{}, timeoutAPIError
	}
}
