package services

import (
	"fmt"
	"strconv"

	customErrors "github.com/ej-you/HamstersShaver/tg_bot/errors"
)


// парсинг значения процента проскальзывания, переданного юзером
func ParseSlippageAmount(stringSlippage string) (string, error) {
	intSlippage, err := strconv.Atoi(stringSlippage)
	if err != nil {
		validateErr := customErrors.ValidateError("failed to parse slippage value")
		return "", fmt.Errorf("parse slippage amount: %w", fmt.Errorf("parse int from %s: %v: %w", stringSlippage, err, validateErr))
	}
	// проверка нахождения числа в диапазоне 1..100
	if intSlippage < 1 || intSlippage > 100 {
		validateErr := customErrors.ValidateError("invalid slippage value")
		return "", fmt.Errorf("parse slippage amount: %w", fmt.Errorf("given slippage amount %d is not in range 1..100: %w", intSlippage, validateErr))
	}
	return stringSlippage, nil
}
