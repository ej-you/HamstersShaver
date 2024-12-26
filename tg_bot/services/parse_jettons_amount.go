package services

import (
	"fmt"
	"strconv"
	"strings"
	"time"

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

	// если введён процент от баланса
	var jettonsAmount string

	// парсим значение процента из строки
	percent, err := strconv.Atoi(strings.TrimSuffix(rawJettonsAmount, "%"))
	if err != nil {
		validateErr := customErrors.ValidateError("failed to parse percent value")
		return "", fmt.Errorf("parse jettons amount: %w", fmt.Errorf("parse int from %s: %v: %w", rawJettonsAmount, err, validateErr))
	}
	// проверка значения процента в пределах 1..100
	if percent <= 0 || percent > 100 {
		validateErr := customErrors.ValidateError("invalid percent value")
		return "", fmt.Errorf("parse jettons amount: %w", validateErr)
	}
	// получение количества монет по проценту
	jettonsAmount, err = getAmountFromPercent(jettonCA, percent)
	if err != nil {
		return "", fmt.Errorf("parse jettons amount: %w", err)
	}
	return jettonsAmount, nil
}


// получение кол-ва монет по проценту, переданному юзером
func getAmountFromPercent(jettonCA string, percent int) (string, error) {
	var err error
	var jettonsAmount int
	var jettonsDecimals int

	// если нужен баланс TON
	if jettonCA == apiClient.TONMasterAddress {
		// получение баланса TON у аккаунта
		var accountTonInfo apiClient.TONInfo
		err = apiClient.GetRequest("/api/account/get-ton", nil, &accountTonInfo, 5*time.Second)
		if err != nil {
			return "", fmt.Errorf("get tons amount from percent: %w", err)
		}
		// рассчёт процента от общего баланса TON
		jettonsAmount = accountTonInfo.Balance / 100 * percent
		jettonsDecimals = accountTonInfo.Decimals

	// если нужен баланс монеты
	} else {
		// получение баланса монеты у аккаунта
		var accountJettonInfo apiClient.AccountJetton
		getAccountJettonInfoParams := apiClient.QueryParams{Params: map[string]interface{}{"MasterAddress": jettonCA}}
		
		err = apiClient.GetRequest("/api/account/get-jetton", &getAccountJettonInfoParams, &accountJettonInfo, 5*time.Second)
		if err != nil {
			return "", fmt.Errorf("get jettons amount from percent: %w", err)
		}
		// рассчёт процента от общего баланса монеты
		jettonsAmount = accountJettonInfo.Balance / 100 * percent
		jettonsDecimals = accountJettonInfo.Decimals
	}

	// запрос на перевод кол-ва монет в человекочитаемый вид
	var beautyBalance apiClient.BeautyBalance
	getBeautyBalanceParams := apiClient.QueryParams{Params: map[string]interface{}{
		"RawBalance": jettonsAmount,
		"Decimals": jettonsDecimals,
	}}
	err = apiClient.GetRequest("/api/services/beauty-balance", &getBeautyBalanceParams, &beautyBalance, 5*time.Second)
	if err != nil {
		return "", fmt.Errorf("get jettons amount from percent: %w", err)
	}

	return beautyBalance.BeautyBalance, nil
}
