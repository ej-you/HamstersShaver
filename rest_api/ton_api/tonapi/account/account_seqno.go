package account

import (
	"context"
	"fmt"

	tonapi "github.com/tonkeeper/tonapi-go"
	tongoWallet "github.com/tonkeeper/tongo/wallet"

	coreErrors "github.com/ej-you/HamstersShaver/rest_api/core/errors"
)


func GetAccountSeqno(ctx context.Context, tonapiClient *tonapi.Client, realWallet tongoWallet.Wallet) (uint32, error) {
	var seqno uint32

	// получение seqno
	seqno, err := tonapiClient.GetSeqno(ctx, realWallet.GetAddress())
	if err != nil {
		apiErr := coreErrors.New(
			fmt.Errorf("get account seqno using tonapi: %w", err),
			"failed to get account seqno",
			"tonApi",
			500,
		)
		apiErr.CheckTimeout()
		return seqno, apiErr
	}
	return seqno, nil
}
