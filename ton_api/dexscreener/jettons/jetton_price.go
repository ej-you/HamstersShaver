package jettons

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/Danil-114195722/HamstersShaver/settings"
)


const apiUrl = "https://api.dexscreener.io/latest/dex/tokens/"


// обработанные данные от API
type JettonsPoolInfo struct {
	// адрес пула
	PoolAddress		string
	// котирующая монета (TONcoin)
	QuoteTokenName	string
	QuoteTokenAddr	string
	// основная монета (ALTcoin)
	BaseTokenName	string
	BaseTokenAddr	string
	// цена в единицах котирующей монеты
	PriceNative		float64
	// цена в долларах
	PriceUsd		float64
}


// получение информации о пулах по адресам двух монет от dexscreener API и десериализация JSON-ответа в структуру jsonPairs
func apiGetJettonsPairs(jettonAddr0, jettonAddr1 string) (jsonPairs, error) {
	var unmarshaledJsonPairs jsonPairs

	// URL для запроса к API
	requestUrl := apiUrl + jettonAddr0 + "," + jettonAddr1
	// запрос к API
	resp, err := http.Get(requestUrl)
	if err != nil {
		settings.ErrorLog.Println("Failed to get pool data from dexscreener API:", err)
		return unmarshaledJsonPairs, err
	}

	defer resp.Body.Close()
	// чтение байт тела ответа от API
	rawJsonBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		settings.ErrorLog.Println("Failed to read data from response body:", err)
		return unmarshaledJsonPairs, err
	}

	// десериализация JSON-ответа в структуру jsonPairs
	err = json.Unmarshal(rawJsonBytes, &unmarshaledJsonPairs)
	if err != nil {
		settings.ErrorLog.Println("Failed to unmarshal JSON-data from response body:", err)
		return unmarshaledJsonPairs, err
	}

	// если не было найдено ни одной пары
	if (unmarshaledJsonPairs.Pairs == nil) {
		pairsNofFoundError := errors.New("pairs value is null")
		settings.ErrorLog.Println("Pairs was not found:", pairsNofFoundError)
		return unmarshaledJsonPairs, pairsNofFoundError
	}

	return unmarshaledJsonPairs, nil
}

// получение информации о пуле двух монет в виде структуры JettonsPoolInfo
func GetJettonsPoolInfo(jettonAddr0, jettonAddr1 string) (JettonsPoolInfo, error) {
	var poolInfo JettonsPoolInfo

	// получение структуры-списка со всеми найденными пулами с их данными от API
	unmarshaledJsonPairs, err := apiGetJettonsPairs(jettonAddr0, jettonAddr1)
	if err != nil {
		return poolInfo, err
	}

	for _, pairInfo := range unmarshaledJsonPairs.Pairs {
		// поиск первого пула на stonfi
		if pairInfo.DexId == "stonfi" {
			// перевод строкового значения цены (в USD) во float64
			priceUSD, err := strconv.ParseFloat(pairInfo.PriceUsd, 64)
			if err != nil {
				settings.ErrorLog.Println("Failed to parse float from string PriceUSD:", err)
				return poolInfo, err
			}

			// перевод строкового значения цены (в QuoteToken) во float64
			priceNative, err := strconv.ParseFloat(pairInfo.PriceNative, 64)
			if err != nil {
				settings.ErrorLog.Println("Failed to parse float from string PriceNative:", err)
				return poolInfo, err
			}


			// заполнение структуры данными пула
			poolInfo = JettonsPoolInfo{
				PoolAddress: pairInfo.PairAddress,

				QuoteTokenName: pairInfo.QuoteToken.Symbol,
				QuoteTokenAddr: pairInfo.QuoteToken.Address,

				BaseTokenName: pairInfo.BaseToken.Symbol,
				BaseTokenAddr: pairInfo.BaseToken.Address,
				PriceNative: priceNative,
				PriceUsd: priceUSD,
			}
			break
		}
	}
	
	if (poolInfo == JettonsPoolInfo{}) {
		poolNofFoundError := errors.New("pool not found error")
		settings.ErrorLog.Println("No pools found:", poolNofFoundError)
		return poolInfo, err
	}

	return poolInfo, nil
}
