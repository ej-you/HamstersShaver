package handlers

import (
	"fmt"
	"time"

	telebot "gopkg.in/telebot.v3"

	apiClient "github.com/ej-you/HamstersShaver/tg_bot/api_client"
	stateMachine "github.com/ej-you/HamstersShaver/tg_bot/state_machine"
	"github.com/ej-you/HamstersShaver/tg_bot/keyboards"
	"github.com/ej-you/HamstersShaver/tg_bot/services"
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
		return fmt.Errorf("HomeHandler for user %s: %w", userId, err)
	}
	// очистка кэша с информацией для новой транзакции
	if err = userStateMachine.ClearNewTransactionPreparation(); err != nil {
		return fmt.Errorf("HomeHandler for user %s: %w", userId, err)
	}

	// получение баланса TON у аккаунта
	var TONAccountInfo apiClient.TONInfo
	err = apiClient.GetRequest("/api/account/get-ton", nil, &TONAccountInfo, 5*time.Second)
	if err != nil {
		return fmt.Errorf("HomeHandler for user %s: %w", userId, err)
	}

	// получение актуального курса TON в долларах
	var TONJettonInfo apiClient.JettonInfo
	getTONJettonInfoParams := apiClient.QueryParams{Params: map[string]string{"MasterAddress": "EQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAM9c"}}
	err = apiClient.GetRequest("/api/jettons/get-info", &getTONJettonInfoParams, &TONJettonInfo, 5*time.Second)
	if err != nil {
		return fmt.Errorf("HomeHandler for user %s: %w", userId, err)
	}

	msgText := fmt.Sprintf(`Главное меню

💰 Текущий баланс TON: %s 
💵 Актуальный курс TON: %.2f$

❗️Для справки используйте /help`,
	TONAccountInfo.BeautyBalance, TONJettonInfo.PriceUSD)

	return context.Send(msgText, keyboards.InlineKeyboardMainMenu)
}
