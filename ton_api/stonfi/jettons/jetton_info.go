package jettons

import (
	"fmt"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/Danil-114195722/HamstersShaver/settings"
)


const apiUrl = "https://api.ston.fi/v1/assets/"


// описывает любую монету (функция получения информации о монете по её адресу)
type JettonParams struct {
	Symbol 			string `json:"symbol"`
	Decimals 		int `json:"decimals"`
	MasterAddress 	string `json:"masterAddress"`
	PriceUSD 		float64 `json:"priseUsd"`
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
func GetJettonInfoByAddres(jettonAddr string) (JettonParams, error) {
	var jettonInfoParse AssetJettonParams

	// запрос к API
	resp, err := http.Get(apiUrl + jettonAddr)
	if err != nil {
		getJettonDataError := errors.New(fmt.Sprintf("Failed to get jetton data from Stonfi API: %s", err.Error()))
		settings.ErrorLog.Println(getJettonDataError.Error())
		return JettonParams{}, getJettonDataError
	}

	defer resp.Body.Close()
	// чтение байт тела ответа от API
	rawJsonBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		readRawDataError := errors.New(fmt.Sprintf("Failed to read data from response body: %s", err.Error()))
		settings.ErrorLog.Println(readRawDataError.Error())
		return JettonParams{}, readRawDataError
	}

	// десериализация JSON-ответа в структуру jsonPairs
	err = json.Unmarshal(rawJsonBytes, &jettonInfoParse)
	if err != nil {
		parseJsonError := errors.New("Jetton was not found")
		settings.ErrorLog.Println(parseJsonError.Error())
		return JettonParams{}, parseJsonError
	}

	// если не было найдено информации о жетоне
	if (jettonInfoParse == AssetJettonParams{}) {
		pairsNofFoundError := errors.New("Jetton was not found")
		settings.ErrorLog.Println(pairsNofFoundError.Error())
		return JettonParams{}, pairsNofFoundError
	}

	// перевод цены монеты из строки во float64
	parsePriceToFloat, err := strconv.ParseFloat(jettonInfoParse.Asset.StringPriceUSD, 64)
	if err != nil {
		parseFloatError := errors.New(fmt.Sprintf("Failed to parse float from StringPriceUSD: %s", err.Error()))
		settings.ErrorLog.Println(parseFloatError.Error())
		return JettonParams{}, parseFloatError
	}

	// формирование выходной структуры с данными
	jettonInfo := JettonParams{
		Symbol: jettonInfoParse.Asset.Symbol,
		Decimals: jettonInfoParse.Asset.Decimals,
		MasterAddress: jettonInfoParse.Asset.MasterAddress,
		PriceUSD: parsePriceToFloat,
	}

	return jettonInfo, nil
}
