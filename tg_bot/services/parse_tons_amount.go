package services

import (
	"fmt"
	"strconv"
	"strings"

	apiClient "github.com/ej-you/HamstersShaver/tg_bot/api_client"
	customErrors "github.com/ej-you/HamstersShaver/tg_bot/errors"
)


// парсинг значения количества TON, переданного юзером
func ParseTonAmount(rawTonAmount string) (string, error) {
	// если введено количество TON
	if !strings.HasSuffix(rawTonAmount, "%") {
		// парсим значение количества TON из строки (для проверки, что введено корректное число)
		tonsAmountFloat64, err := strconv.ParseFloat(rawTonAmount, 64)
		if err != nil {
			validateErr := customErrors.ValidateError("failed to parse amount value")
			return "", fmt.Errorf("parse jettons amount: %w", fmt.Errorf("parse float64 from %s: %v: %w", rawTonAmount, err, validateErr))
		}
		if tonsAmountFloat64 <= 0 {
			validateErr := customErrors.ValidateError("invalid amount value")
			return "", fmt.Errorf("parse jettons amount: %w", fmt.Errorf("given jettons amount %f is non-positive: %w", tonsAmountFloat64, validateErr))
		}
		return rawTonAmount, nil
	}

	// если введён процент от баланса TON
	// парсим значение процента из строки
	percent, err := strconv.Atoi(strings.TrimSuffix(rawTonAmount, "%"))
	if err != nil {
		validateErr := customErrors.ValidateError("failed to parse percent value")
		return "", fmt.Errorf("parse jettons amount: %w", fmt.Errorf("parse int from %s: %v: %w", rawTonAmount, err, validateErr))
	}

	// получение количества TON по проценту
	var tonAmountFromPercent apiClient.TonAmountFromPercent
	getTonAmountFromPercentParams := apiClient.QueryParams{Params: map[string]interface{}{
		"percent": percent,
	}}
	err = apiClient.GetRequest("/api/services/ton-amount-from-percent", &getTonAmountFromPercentParams, &tonAmountFromPercent)
	if err != nil {
		return "", fmt.Errorf("parse TON amount: get TON amount from percent: %w", err)
	}

	return tonAmountFromPercent.TonAmount, nil
}


// получение кол-ва монет по проценту, переданному юзером при нажатии на кнопку
func GetTonAmountFromPercentFromCallback(callbackData string) (string, error) {
	// достаём процент из данных кнопки и переводим его в int
	intPercent, err := strconv.Atoi(strings.TrimPrefix(callbackData, "jettons_amount_choice|"))
	if err != nil {
		internalErr := customErrors.InternalError("failed to parse percent value")
		return "", fmt.Errorf("get TON amount from callback: parse int from jettons_amount_choice button: %v: %w", err, internalErr)
	}
	
	// получение количества TON по проценту
	var tonAmountFromPercent apiClient.TonAmountFromPercent
	getTonAmountFromPercentParams := apiClient.QueryParams{Params: map[string]interface{}{
		"percent": intPercent,
	}}
	err = apiClient.GetRequest("/api/services/ton-amount-from-percent", &getTonAmountFromPercentParams, &tonAmountFromPercent)
	if err != nil {
		return "", fmt.Errorf("get TON amount from callback: get TON amount from percent: %w", err)
	}

	return tonAmountFromPercent.TonAmount, nil
}
