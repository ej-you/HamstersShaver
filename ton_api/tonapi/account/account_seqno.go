package account

import (
	"context"
	"errors"
	"fmt"

	tonapi "github.com/tonkeeper/tonapi-go"
	tongoWallet "github.com/tonkeeper/tongo/wallet"
)


func GetAccountSeqno(ctx context.Context, tonapiClient *tonapi.Client, realWallet tongoWallet.Wallet) (uint32, error) {
	var seqno uint32

	// получение seqno
	seqno, err := tonapiClient.GetSeqno(ctx, realWallet.GetAddress())
	if err != nil {
		getSeqnoError := errors.New(fmt.Sprintf("Failed to get seqno: %s", err.Error()))
		return seqno, getSeqnoError
	}
	
	return seqno, nil
}
