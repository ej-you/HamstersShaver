package services

import (
	"math"
	"math/big"
)


// перевод кол-ва токенов в нужный вид для транзакции tongo
func ConvertJettonsAmountToBigInt(decimals int, amount float64) *big.Int {
	poweredAmount := int64(amount * math.Pow10(decimals))
	return big.NewInt(poweredAmount)
}

// перевод кол-ва токенов из float64 в uint64
func ConvertJettonsAmountToUint(decimals int, amount float64) uint64 {
	poweredAmount := uint64(amount * math.Pow10(decimals))
	return poweredAmount
}
