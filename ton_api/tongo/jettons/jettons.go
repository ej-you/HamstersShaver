package jettons

import (
	"math"
	"math/big"

	"github.com/tonkeeper/tongo"
	tongoJettons "github.com/tonkeeper/tongo/contract/jetton"

	"github.com/Danil-114195722/HamstersShaver/settings"
)


// получение структуры жетона 
func GetJettonStruct(addr string) *tongoJettons.Jetton {
	master := tongo.MustParseAccountID(addr)
	jettonStruct := tongoJettons.New(master, settings.TongoTonAPI)

	return jettonStruct
}

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
