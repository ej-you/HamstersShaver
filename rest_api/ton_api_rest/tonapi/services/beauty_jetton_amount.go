package services

import (
	"math"
	"regexp"
	"strconv"
)


// кол-во цифр после запятой, которые нужно оставить при "округлении"
const decimalsForBeautyJettonAmount = 2


// чистка нулей на конце после точки, кроме первого
// (543.000 - удалятся только 2 последних нуля)
// (8673.567000 - удалятся все нули)
// (8673.560006 - ничего не удалится)
func clearZerosEnd(number string) string {
	re := regexp.MustCompile(`(\.\d+?)(0+)$`)
	clearZeroNumber := re.ReplaceAllString(number, "${1}")

	return clearZeroNumber
}


// перевод int64 значения кол-ва монет в string с точностью до decimalsForBeautyJettonAmount знаков (остальное отбрасывается)
func BeautyJettonAmountFromInt64(int64Amount int64, decimals int) string {
	// перевод кол-ва монет из int64 во float64
	float64Amount := float64(int64Amount) / math.Pow10(decimals)
	// "округлённый" результат
	return BeautyJettonAmountFromFloat64(float64Amount, decimals)
}


// перевод float64 значения кол-ва монет в string с точностью до decimalsForBeautyJettonAmount знаков (остальное отбрасывается)
func BeautyJettonAmountFromFloat64(float64Amount float64, decimals int) string {
	// отбрасываем всё после 2 знаков
	accuracy := math.Pow10(decimalsForBeautyJettonAmount)
	flooredFloat64Amount := math.Floor(float64Amount * accuracy) / accuracy

	// если целая часть числа меньше 10 и после "округления" после точки остались только нули то не используем это "округление"
	if int(float64Amount) < 10 && flooredFloat64Amount == float64(int(flooredFloat64Amount)) {
		// "неокруглённый" результат (берём decimals цифр после запятой и чистим лишние нули на конце)
		return clearZerosEnd(strconv.FormatFloat(float64Amount, 'f', decimals, 64))
	}
	// "округлённый" результат
	return strconv.FormatFloat(flooredFloat64Amount, 'f', -1, 64)
}
