package jettons

import (
	"context"
	"strconv"

	tonapi "github.com/tonkeeper/tonapi-go"

	"github.com/Danil-114195722/HamstersShaver/settings"
)


// описывает любую монету (функция получения информации о монете по её адресу)
type JettonParams struct {
	Symbol string
	Decimals int
	// мастер-адрес монеты
	MasterAddress string
}

// получение информации о монете по её адресу
func GetJettonInfoByAddres(ctx context.Context, addr string) (JettonParams, error) {
	var rawJettonInfo *tonapi.JettonInfo
	var jettonInfo JettonParams

	// параметры для получения информации о монете
	params := tonapi.GetJettonInfoParams{AccountID: addr}

	// запрос данных
	rawJettonInfo, err := settings.TonapiTonAPI.GetJettonInfo(ctx, params)
	if err != nil {
		settings.ErrorLog.Println("Failed to get jetton info:", err)
		return jettonInfo, err
	}

	// перевод значения rawJettonInfo.Metadata.Decimals из строки в число
	intDecimals, err := strconv.Atoi(rawJettonInfo.Metadata.Decimals)
	if err != nil {
		settings.ErrorLog.Println("Failed to parse int from jetton decimals string:", err)
		return jettonInfo, err
	}

	// заполнение структуры нужной информацией
	jettonInfo = JettonParams{
		Symbol: rawJettonInfo.Metadata.Symbol,
		Decimals: intDecimals,
		MasterAddress: rawJettonInfo.Metadata.Address,
	}
	return jettonInfo, nil
}
