package api_client


// запрос на /api/account/get-jetton
// запрос на /api/account/get-jettons
type AccountJetton struct {
	Symbol 			string `json:"symbol"`
	BeautyBalance 	string `json:"beautyBalance"`
	MasterAddress 	string `json:"masterAddress"`
}

// запрос на /api/jettons/get-info
type JettonInfo struct {
	Symbol 			string `json:"symbol"`
	MasterAddress 	string `json:"masterAddress"`
	PriceUSD 		float64 `json:"priceUsd"`
}

// запрос на /api/account/get-ton
type TONInfo struct {
	BeautyBalance 	string `json:"beautyBalance"`
}
