package services

import (
	"fmt"
	"strconv"
	"strings"

	apiClient "github.com/ej-you/HamstersShaver/tg_bot/api_client"
	customErrors "github.com/ej-you/HamstersShaver/tg_bot/errors"
)


// парсинг значения количества монет, переданного юзером
func ParseJettonsAmount(jettonCA, rawJettonsAmount string) (string, error) {
	// если введено количество монет
	if !strings.HasSuffix(rawJettonsAmount, "%") {
		// парсим значение количества из строки (для проверки, что введено корректное число)
		jettonsAmountFloat64, err := strconv.ParseFloat(rawJettonsAmount, 64)
		if err != nil {
			validateErr := customErrors.ValidateError("failed to parse amount value")
			return "", fmt.Errorf("parse jettons amount: %w", fmt.Errorf("parse float64 from %s: %v: %w", rawJettonsAmount, err, validateErr))
		}
		if jettonsAmountFloat64 <= 0 {
			validateErr := customErrors.ValidateError("invalid amount value")
			return "", fmt.Errorf("parse jettons amount: %w", fmt.Errorf("given jettons amount %f is non-positive: %w", jettonsAmountFloat64, validateErr))
		
		}
		return rawJettonsAmount, nil
	}

	// если введён процент от баланса монеты
	// парсим значение процента из строки
	percent, err := strconv.Atoi(strings.TrimSuffix(rawJettonsAmount, "%"))
	if err != nil {
		validateErr := customErrors.ValidateError("failed to parse percent value")
		return "", fmt.Errorf("parse jettons amount: %w", fmt.Errorf("parse int from %s: %v: %w", rawJettonsAmount, err, validateErr))
	}

	// получение количества монет по проценту
	var jettonAmountFromPercent apiClient.JettonAmountFromPercent
	getJettonAmountFromPercentParams := apiClient.QueryParams{
		"percent": percent,
		"masterAddress": jettonCA,
	}
	err = apiClient.GetRequest("/api/services/jetton-amount-from-percent", &getJettonAmountFromPercentParams, &jettonAmountFromPercent)
	if err != nil {
		return "", fmt.Errorf("parse jetton amount: get jetton amount from percent: %w", err)
	}

	return jettonAmountFromPercent.JettonAmount, nil
}


// получение кол-ва монет по проценту, переданному юзером
func GetJettonAmountFromPercentFromCallback(jettonCA string, callbackData string) (string, error) {
	// достаём процент из данных кнопки и переводим его в int
	intPercent, err := strconv.Atoi(strings.TrimPrefix(callbackData, "jettons_amount_choice|"))
	if err != nil {
		internalErr := customErrors.InternalError("failed to parse percent value")
		return "", fmt.Errorf("get jetton amount from callback: parse int from jettons_amount_choice button: %v: %w", err, internalErr)
	}

	// получение количества монет по проценту
	var jettonAmountFromPercent apiClient.JettonAmountFromPercent
	getJettonAmountFromPercentParams := apiClient.QueryParams{
		"percent": intPercent,
		"masterAddress": jettonCA,
	}
	err = apiClient.GetRequest("/api/services/jetton-amount-from-percent", &getJettonAmountFromPercentParams, &jettonAmountFromPercent)
	if err != nil {
		return "", fmt.Errorf("parse jetton amount: get jetton amount from percent: %w", err)
	}

	return jettonAmountFromPercent.JettonAmount, nil
}
