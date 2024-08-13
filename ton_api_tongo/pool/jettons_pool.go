package pool

import (
	"fmt"
	"errors"
	"context"

	tongoTon "github.com/tonkeeper/tongo/ton"
	tongoAbi "github.com/tonkeeper/tongo/abi"

	"github.com/Danil-114195722/HamstersShaver/settings"
	"github.com/Danil-114195722/HamstersShaver/settings/constants"
)


func GetJettonsPoolAddress(ctx context.Context, jettonMaster0, jettonMaster1 tongoTon.AccountID) (string, error) {
	var poolAddr string
	var err error

	// перевод адреса StonfiRouter в структуру tongoTon.AccountID
	StonfiRouterAccountID := tongoTon.MustParseAccountID(constants.StonfiRouterAddr)

	// получение пула для данной пары монет на Stonfi
	_, value, err := tongoAbi.GetPoolAddress(ctx, settings.TongoTonAPI, StonfiRouterAccountID, jettonMaster0.ToMsgAddress(), jettonMaster1.ToMsgAddress())
	if err != nil {
		settings.ErrorLog.Println("Failed to get pool address:", err)
		return poolAddr, err
	}

	// проверка принадлежности значения value к типу tongoAbi.GetPoolAddress_StonfiResult
	result, ok := value.(tongoAbi.GetPoolAddress_StonfiResult)
	if !ok {
		// ошибка утверждения
		assertError := errors.New("Assertion error")
		settings.ErrorLog.Println("Failed to assert abi.GetPoolAddress value:", assertError)
		return poolAddr, assertError
	}

	// перевод результата в структуру tongoTon.AccountID
	poolAccountID, err := tongoTon.AccountIDFromTlb(result.PoolAddress)
	if err != nil {
		settings.ErrorLog.Println("Failed to get Pool's AccountID from Tlb:", err)
		return poolAddr, err
	}

	fmt.Printf("poolAccountID: %T | %v\n", *poolAccountID, *poolAccountID)
	return "GREAT", nil
}

// 0:1c754f54c40fb81d1e5081ff7917976d050e58caa3ce7c55c986ad6cfbb6a2c9
// EQAcdU9UxA-4HR5Qgf95F5dtBQ5YyqPOfFXJhq1s-7aiyWWX


// Dexscreener
// pair: EQCaY8Ifl2S6lRBMBJeY35LIuMXPc8JfItWG4tl7lBGrSoR2
// TON:  EQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAM9c