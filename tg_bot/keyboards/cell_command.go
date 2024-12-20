package keyboards

import (
	"encoding/json"
	"fmt"
	"net/http"

	telebot "gopkg.in/telebot.v3"

	"github.com/ej-you/HamstersShaver/tg_bot/settings"
)


// получение рядов кнопок для клавиатуры с монетами и их количеством на кошельке аккаунта
func SetWalletJettonsBtnRows(replyMarkup *telebot.ReplyMarkup) error {
	var inlineRows []telebot.Row

	// получение нужной информации о всех монетах на кошельке аккаунта
	allJettons, err := getJettonsInfo()
	if err != nil {
		return err
	}

	// создание рядов кнопок
	var button telebot.Btn
	for _, jetton := range allJettons {
		// создание ряда с кнопкой и добавление его в срез рядов
		button = replyMarkup.Data(fmt.Sprintf("%s — %s", jetton.Symbol, jetton.BeautyBalance), jetton.MasterAddress)
		inlineRows = append(inlineRows, replyMarkup.Row(button))
	}

	// добавление рядов кнопок в клавиатуру
	replyMarkup.Inline(inlineRows...)

	return nil
}


type jettonInfo struct {
	Symbol 			string `json:"symbol"`
	BeautyBalance 	string `json:"beautyBalance"`
	MasterAddress 	string `json:"masterAddress"`
}


// получение монет аккаунта и их кол-во
func getJettonsInfo() ([]jettonInfo, error) {
	var allJettonsInfo []jettonInfo
	client := &http.Client{}

	// обращение к API для получения баланса TON
	req, err := http.NewRequest("GET", settings.RestApiHost+"/api/account/get-jettons", nil)
	if err != nil {
		return allJettonsInfo, fmt.Errorf("create get account jettons request error: %v", err)
	}
	// добавление query-параметров
	queryParams := req.URL.Query()
	queryParams.Add("api-key", settings.RestApiKey)
	req.URL.RawQuery = queryParams.Encode()
	// отправка запроса
	resp, err := client.Do(req)
	if err != nil {
		return allJettonsInfo, fmt.Errorf("failed to get account jettons: %v", err)
	}
	defer resp.Body.Close()
	// декодирование ответа в структуру
	if err := json.NewDecoder(resp.Body).Decode(&allJettonsInfo); err != nil {
		return allJettonsInfo, fmt.Errorf("failed to decode answer from get account jettons response: %v", err)
	}

	return allJettonsInfo, nil
}
