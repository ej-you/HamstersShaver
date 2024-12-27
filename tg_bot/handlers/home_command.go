package handlers

import (
	"fmt"

	telebot "gopkg.in/telebot.v3"

	apiClient "github.com/ej-you/HamstersShaver/tg_bot/api_client"
	stateMachine "github.com/ej-you/HamstersShaver/tg_bot/state_machine"
	"github.com/ej-you/HamstersShaver/tg_bot/keyboards"
	"github.com/ej-you/HamstersShaver/tg_bot/services"
)


// команда: /home
// кнопки: to_home
func HomeHandler(context telebot.Context) error {
	var err error

	// сообщение о начале подгрузки данных
	loadMsg, err := context.Bot().Send(context.Recipient(), "Загрузка...")
	if err != nil {
		return err
	}

	// получение баланса TON у аккаунта
	var TONAccountInfo apiClient.TONInfo
	err = apiClient.GetRequest("/api/account/get-ton", nil, &TONAccountInfo)
	if err != nil {
		return fmt.Errorf("HomeHandler: %w", err)
	}

	// получение актуального курса TON в долларах
	var TONJettonInfo apiClient.JettonInfo
	getTONJettonInfoParams := apiClient.QueryParams{Params: map[string]interface{}{
		"MasterAddress": apiClient.TONMasterAddress,
	}}
	err = apiClient.GetRequest("/api/jettons/get-info", &getTONJettonInfoParams, &TONJettonInfo)
	if err != nil {
		return fmt.Errorf("HomeHandler: %w", err)
	}

	// получение машины состояний текущего юзера
	userStateMachine := stateMachine.UserStateMachines.Get(services.GetUserID(context.Chat()))
	// установка нового состояния
	if err = userStateMachine.SetStatus("home"); err != nil {
		return fmt.Errorf("HomeHandler: %w", err)
	}
	// очистка кэша с информацией для новой транзакции
	if err = userStateMachine.ClearNewTransactionPreparation(); err != nil {
		return fmt.Errorf("HomeHandler: %w", err)
	}

	msgText := fmt.Sprintf(`Главное меню

💰 Текущий баланс TON: %s 
💵 Актуальный курс TON: %.2f$

❗️Для справки используйте /help`,
	TONAccountInfo.BeautyBalance, TONJettonInfo.PriceUSD)

	// редактирование сообщения о загрузке - вывод данных
	_, err = context.Bot().Edit(loadMsg, msgText, keyboards.InlineKeyboardMainMenu)
	return err
}
