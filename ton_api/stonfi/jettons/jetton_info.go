package jettons

import (
	"fmt"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/Danil-114195722/HamstersShaver/settings"
)


const apiUrl = "https://api.ston.fi/v1/assets/"


// описывает любую монету (функция получения информации о монете по её адресу)
type JettonParams struct {
	Symbol 			string `json:"symbol"`
	Decimals 		int `json:"decimals"`
	MasterAddress 	string `json:"masterAddress"`
	PriceUSD 		float64 `json:"priceUsd"`
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

	// запрос к API
	resp, err := http.Get(apiUrl + jettonAddr)
	if err != nil {
		getJettonDataError := errors.New(fmt.Sprintf("Failed to get jetton data from Stonfi API: %s", err.Error()))
		settings.ErrorLog.Println(getJettonDataError.Error())
		return jettonInfoByAddress{JettonInfo: JettonParams{}, Error: getJettonDataError}
	}

	defer resp.Body.Close()
	// чтение байт тела ответа от API
	rawJsonBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		readRawDataError := errors.New(fmt.Sprintf("Failed to read data from response body: %s", err.Error()))
		settings.ErrorLog.Println(readRawDataError.Error())
		return jettonInfoByAddress{JettonInfo: JettonParams{}, Error: readRawDataError}
	}

	// десериализация JSON-ответа в структуру jsonPairs
	err = json.Unmarshal(rawJsonBytes, &jettonInfoParse)
	if err != nil {
		parseJsonError := errors.New("Jetton was not found")
		settings.ErrorLog.Println(parseJsonError.Error())
		return jettonInfoByAddress{JettonInfo: JettonParams{}, Error: parseJsonError}
	}

	// если не было найдено информации о жетоне
	if (jettonInfoParse == AssetJettonParams{}) {
		infoNofFoundError := errors.New("Jetton was not found")
		settings.ErrorLog.Println(infoNofFoundError.Error())
		return jettonInfoByAddress{JettonInfo: JettonParams{}, Error: infoNofFoundError}
	}

	// перевод цены монеты из строки во float64
	parsePriceToFloat, err := strconv.ParseFloat(jettonInfoParse.Asset.StringPriceUSD, 64)
	if err != nil {
		parseFloatError := errors.New(fmt.Sprintf("Failed to parse float from StringPriceUSD: %s", err.Error()))
		settings.ErrorLog.Println(parseFloatError.Error())
		return jettonInfoByAddress{JettonInfo: JettonParams{}, Error: parseFloatError}
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
			timeoutError := errors.New("Failed to get jetton data from Stonfi API: timeout error")
			settings.ErrorLog.Println(timeoutError.Error())
			return JettonParams{}, timeoutError
	}
}
