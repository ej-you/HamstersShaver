package account

import (
	"errors"
	"context"

	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/tlb"

	"github.com/Danil-114195722/HamstersShaver/settings"
)


// получение аккаунта по данным из JSON-конфига
func GetAccount(ctx context.Context) *tlb.Account {
	// получение главного блока
	block, err := settings.TonutilsTonAPI.CurrentMasterchainInfo(ctx)
	settings.DieIf(err)

	// парсим адрес
	addr := address.MustParseAddr(settings.JsonWallet.Hash)

	// получение аккаунта
	acc, err := settings.TonutilsTonAPI.WaitForBlock(block.SeqNo).GetAccount(ctx, block, addr)
	settings.DieIf(err)

	// проверка того, что аккаунт активен
	if !acc.IsActive {
		accountIsNotActiveError := errors.New("account is not active")
		settings.DieIf(accountIsNotActiveError)
	}

	return acc
}
