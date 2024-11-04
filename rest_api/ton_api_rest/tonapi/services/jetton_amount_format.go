package services

import (
	"fmt"
	"strconv"
)


// форматирует float64 число кол-ва монеты в округлённый float в виде строки
func JettonFloatAmountFormat(floatBalance float64, decimals int) string {
	// преобразование в строку
	stringBalance := fmt.Sprintf("%f", floatBalance)

	// округляем до 3 знаков
	roundedBalance := fmt.Sprintf("%.3f", floatBalance)
	roundedFloatBalance, _ := strconv.ParseFloat(roundedBalance, 64)

	// если целая часть числа меньше 10 и после округления после точки остались только нули
	// то не используем это округление в return
	if int(floatBalance) < 10 && roundedFloatBalance == float64(int(roundedFloatBalance)) {
		// чистим от лишних нулей неокруглённый результат (функция в файле ./jetton_balance_format)
		return clearZerosEnd(stringBalance)
	}

	// чистим от лишних нулей округлённый результат (функция в файле ./jetton_balance_format)
	return clearZerosEnd(roundedBalance)
}
