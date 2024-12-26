package api_client


// мастер-адрес TON для получения информации о нём
const TONMasterAddress = "EQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAM9c"


// запрос на /api/account/get-jetton
// запрос на /api/account/get-jettons
type AccountJetton struct {
	Symbol 			string `json:"symbol"`
	Decimals 		int `json:"decimals"`
	Balance 		int `json:"balance"`
	BeautyBalance 	string `json:"beautyBalance"`
	MasterAddress 	string `json:"masterAddress"`
}

// запрос на /api/jettons/get-info
type JettonInfo struct {
	Symbol			string `json:"symbol"`
	MasterAddress	string `json:"masterAddress"`
	PriceUSD		float64 `json:"priceUsd"`
}

// запрос на /api/account/get-ton
type TONInfo struct {
	BeautyBalance 	string `json:"beautyBalance"`
	Balance 		int `json:"balance"`
	Decimals 		int `json:"decimals"`
}


// запрос на /api/services/beauty-balance
type BeautyBalance struct {
	BeautyBalance string `json:"beautyBalance"`
}

// запрос на /api/transactions/buy/pre-request
type PreRequestBuyJetton struct {
	UsedTON 		string `json:"usedTon"`
	JettonCA 		string `json:"jettonCA"`
	DEX 			string `json:"dex"`
	JettonsOut 		string `json:"jettonsOut"`
	MinOut	 		string `json:"minOut"`
	JettonSymbol 	string `json:"jettonSymbol"`
}

// запрос на /api/transactions/cell/pre-request
type PreRequestCellJetton struct {
	UsedJettons		string `json:"usedJettons"`
	JettonCA 		string `json:"jettonCA"`
	DEX 			string `json:"dex"`
	TONsOut 		string `json:"tonsOut"`
	MinOut	 		string `json:"minOut"`
	JettonSymbol 	string `json:"jettonSymbol"`
}
