package jettons

import (
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
	Symbol 			string
	Decimals 		int
	MasterAddress 	string
	PriceUSD 		float64
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
		settings.ErrorLog.Println("Failed to get jetton data from Stonfi API:", err)
		return JettonParams{}, err
	}

	defer resp.Body.Close()
	// чтение байт тела ответа от API
	rawJsonBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		settings.ErrorLog.Println("Failed to read data from response body:", err)
		return JettonParams{}, err
	}

	// десериализация JSON-ответа в структуру jsonPairs
	err = json.Unmarshal(rawJsonBytes, &jettonInfoParse)
	if err != nil {
		settings.ErrorLog.Println("Jetton was not found. Failed to unmarshal JSON:", err)
		return JettonParams{}, err
	}

	// если не было найдено ни одной пары
	if (jettonInfoParse == AssetJettonParams{}) {
		pairsNofFoundError := errors.New("jetton info is null")
		settings.ErrorLog.Println("Jetton was not found:", pairsNofFoundError)
		return JettonParams{}, pairsNofFoundError
	}

	parsePriceToFloat, err := strconv.ParseFloat(jettonInfoParse.Asset.StringPriceUSD, 64)
	if err != nil {
		settings.ErrorLog.Println("Failed to parse float from StringPriceUSD:", err)
		return JettonParams{}, err
	}

	jettonInfo := JettonParams{
		Symbol: jettonInfoParse.Asset.Symbol,
		Decimals: jettonInfoParse.Asset.Decimals,
		MasterAddress: jettonInfoParse.Asset.MasterAddress,
		PriceUSD: parsePriceToFloat,
	}

	return jettonInfo, nil
}
