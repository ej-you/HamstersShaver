package jettons

import (
	"math"
	"math/big"

	"github.com/tonkeeper/tongo"
	tongoJettons "github.com/tonkeeper/tongo/contract/jetton"

	myTonapiJettons "github.com/Danil-114195722/HamstersShaver/ton_api_tonapi/jettons"
	"github.com/Danil-114195722/HamstersShaver/settings"
)


// получение структуры жетона 
func GetJettonStruct(addr string) *tongoJettons.Jetton {
	master := tongo.MustParseAccountID(addr)
	jettonStruct := tongoJettons.New(master, settings.TongoTonAPI)

	return jettonStruct
}

// перевод кол-ва токенов в нужный вид для транзакции tongo
func ConvertJettonsAmountToBigInt(jetton myTonapiJettons.AccountJetton, amount float64) *big.Int {
	poweredAmount := int64(amount * math.Pow10(jetton.Decimals))
	return big.NewInt(poweredAmount)
}