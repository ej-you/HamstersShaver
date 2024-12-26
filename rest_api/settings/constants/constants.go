package constants

import "time"


// константы для TON API
const (
	StonfiRouterAddr = "EQB3ncyBUTjZUA5EnFKR5_EnOMI9V1tTEAAPaiU71gc4TiUt"

	// для проведения транзакции
	ProxyTonMasterAddr = "EQCM3B12QK1e4yZSf8GtBRT0aLMNyEsBc_DhVfRRtOEffLez"
	// для получения инфы о TON от Stonfi API
	TonInfoAddr = "EQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAM9c"

	TonDecimals = 9
)


// таймауты для разных запросов
const (
	// клиенты
	TonapiClientTimeout = 1*time.Second
	TongoClientTimeout = 2*time.Second

	// Stonfi API
	GetJettonInfoByAddressTimeout = 4*time.Second

	// tonapi
	GetAccountJettonContextTimeout = 1*time.Second
	GetBalanceJettonsContextTimeout = 1*time.Second
	GetBalanceTONContextTimeout = 1*time.Second
	GetTransInfoContextTimeout = 1*time.Second

	// tongo
	GetAccountSeqnoContextTimeout = 2*time.Second
	SendBuyJettonContextTimeout = 2*time.Second
	SendCellJettonContextTimeout = 2*time.Second
)
