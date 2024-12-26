package keyboards

import (
	"fmt"
	"time"

	telebot "gopkg.in/telebot.v3"

	apiClient "github.com/ej-you/HamstersShaver/tg_bot/api_client"
)


var InlineKeyboardWalletJettons = &telebot.ReplyMarkup{}

// получение рядов кнопок для клавиатуры с монетами и их количеством на кошельке аккаунта
func SetWalletJettonsButtons() error {
	var inlineRows []telebot.Row

	// получение нужной информации о всех монетах на кошельке аккаунта
	var allJettons []apiClient.AccountJetton
	err := apiClient.GetRequest("/api/account/get-jettons", nil, &allJettons, 5*time.Second)
	if err != nil {
		return fmt.Errorf("set wallet jettons buttons: %w", err)
	}

	// создание рядов кнопок
	var button telebot.Btn
	for _, jetton := range allJettons {
		// создание ряда с кнопкой и добавление его в срез рядов
		button = InlineKeyboardWalletJettons.Data(fmt.Sprintf("%s — %s", jetton.Symbol, jetton.BeautyBalance), "", jetton.MasterAddress)
		inlineRows = append(inlineRows, InlineKeyboardWalletJettons.Row(button))
	}

	// добавление рядов кнопок в клавиатуру
	InlineKeyboardWalletJettons.Inline(inlineRows...)
	return nil
}
