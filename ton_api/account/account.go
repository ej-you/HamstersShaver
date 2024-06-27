package account

import (
	"errors"
	"context"

	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/tlb"
	
	"github.com/Danil-114195722/HamstersShaver/ton_api"
	"github.com/Danil-114195722/HamstersShaver/settings"
)


// получение данных о кошельке из JSON-конфига
var JsonWallet settings.JsonWallet = settings.GetJsonWallet()

// получение аккаунта по данным из JSON-конфига
func GetAccount() *tlb.Account {
	ctx := context.Background()

	// получение главного блока
	block, err := ton_api.API.CurrentMasterchainInfo(ctx)
	settings.DieIf(err)

	// парсим адрес
	addr := address.MustParseAddr(JsonWallet.Hash)

	// получение аккаунта
	acc, err := ton_api.API.WaitForBlock(block.SeqNo).GetAccount(ctx, block, addr)
	settings.DieIf(err)

	// проверка того, что аккаунт активен
	if !acc.IsActive {
		accountIsNotActiveError := errors.New("account is not active")
		settings.DieIf(accountIsNotActiveError)
	}

	return acc
}

