package services

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)


// чистка нулей после точки, кроме первого
// (543.000 - удалятся только 2 последних нуля)
// (8673.567000 - удалятся все нули)
// (8673.560006 - ничего не удалится)
func clearZerosEnd(number string) string {
	re := regexp.MustCompile(`(\.\d+?)(0+)$`)
	clearZeroNumber := re.ReplaceAllString(number, "${1}")

	return clearZeroNumber
}

// форматирует целое число баланса монеты в округлённый float в виде строки
func JettonBalanceFormat(balance int64, decimals int) string {
	// перевод баланса из int64 во float64
	floatBalance := float64(balance) / math.Pow10(decimals)
	// преобразование в строку
	stringBalance := fmt.Sprintf("%f", floatBalance)

	// округляем до 3 знаков
	roundedBalance := fmt.Sprintf("%.3f", floatBalance)
	roundedFloatBalance, _ := strconv.ParseFloat(roundedBalance, 64)

	// если целая часть числа меньше 10 и после округления после точки остались только нули
	// то не используем это округление в return
	if int(floatBalance) < 10 && roundedFloatBalance == float64(int(roundedFloatBalance)) {
		// чистим от лишних нулей неокруглённый результат
		return clearZerosEnd(stringBalance)
	}

	// чистим от лишних нулей округлённый результат
	return clearZerosEnd(roundedBalance)
}
