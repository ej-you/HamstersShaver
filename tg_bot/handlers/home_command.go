package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	telebot "gopkg.in/telebot.v3"

	"github.com/ej-you/HamstersShaver/tg_bot/keyboards"
	"github.com/ej-you/HamstersShaver/tg_bot/services"
	stateMachine "github.com/ej-you/HamstersShaver/tg_bot/state_machine"
	"github.com/ej-you/HamstersShaver/tg_bot/settings"
)


// команды: /home || /cancel
// кнопки: to_home
func HomeHandler(context telebot.Context) error {
	var err error
	userId := services.GetUserID(context.Chat())

	// получение машины состоянию текущего юзера
	userStateMachine := stateMachine.UserStateMachines.Get(userId)
	// установка нового состояния
	if err = userStateMachine.SetStatus("home"); err != nil {
		return err
	}
	// очистка кэша с информацией для новой транзакции
	if err = userStateMachine.ClearNewTransactionPreparation(); err != nil {
		return err
	}

	fullTonInfo, err := getTONInfo()
	if err != nil {
		settings.ErrorLog.Printf("ERROR FROM API: %v", err)
		msgText := fmt.Sprintf("😬 Ошибка в получении информации о TON:\n\t%s", err.Error())
		return context.Send(msgText, keyboards.InlineKeyboardToHome)
	}

	return context.Send(homeGetMessageText(fullTonInfo), keyboards.InlineKeyboardMainMenu)
}


type tonInfo struct {
	BeautyBalance 	string `json:"beautyBalance"`
	PriceUSD 		float64 `json:"priceUsd"`
}


// получение баланса TON у аккаунта и актуальный курс TON в долларах
func getTONInfo() (tonInfo, error) {
	var fullTonInfo tonInfo
	client := &http.Client{Timeout: 5*time.Second}

	// обращение к API для получения баланса TON
	req, err := http.NewRequest("GET", settings.RestApiHost+"/api/account/get-ton", nil)
	if err != nil {
		return fullTonInfo, fmt.Errorf("create get TON balance request error: %v", err)
	}
	// добавление query-параметров
	queryParams := req.URL.Query()
	queryParams.Add("api-key", settings.RestApiKey)
	req.URL.RawQuery = queryParams.Encode()
	// отправка запроса
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return fullTonInfo, fmt.Errorf("failed to get TON balance: %v", err)
	}
	// декодирование ответа в структуру (берём только округлённый баланс в виде строки)
	if err := json.NewDecoder(resp.Body).Decode(&fullTonInfo); err != nil {
		return fullTonInfo, fmt.Errorf("failed to decode answer from get TON balance response: %v", err)
	}

	// обращение к API для получения информации о монете TON
	req, err = http.NewRequest("GET", settings.RestApiHost+"/api/jettons/get-info", nil)
	if err != nil {
		return fullTonInfo, fmt.Errorf("create get TON info request error: %v", err)
	}
	// добавление query-параметров
	queryParams = req.URL.Query()
	queryParams.Add("api-key", settings.RestApiKey)
	queryParams.Add("MasterAddress", "EQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAM9c")
	req.URL.RawQuery = queryParams.Encode()
	// отправка запроса
	resp, err = client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return fullTonInfo, fmt.Errorf("failed to get TON info: %v", err)
	}
	// декодирование ответа в структуру (берём только цену TON в долларах)
	if err := json.NewDecoder(resp.Body).Decode(&fullTonInfo); err != nil {
		return fullTonInfo, fmt.Errorf("failed to decode answer from get TON info response: %v", err)
	}

	return fullTonInfo, nil
}

// формирование текста ответа
func homeGetMessageText(fullTonInfo tonInfo) string {
	return fmt.Sprintf(`Главное меню

💰 Текущий баланс TON: %s 
💵 Актуальный курс TON: %.2f$

❗️Для справки используйте /help`, fullTonInfo.BeautyBalance, fullTonInfo.PriceUSD)

}
