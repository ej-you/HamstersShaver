package services

import (
	"math"
)


// перевод int64 баланса во float64
func ConvertBalanceToFloat64(intBalance int64, decimals int) float64 {
	return float64(intBalance) / math.Pow10(decimals)
}
