package jettons

// Определение структуры для токена
type token struct {
	Address string `json:"address"`
	Name    string `json:"name"`
	Symbol  string `json:"symbol"`
}

// Определение структуры для пары
type jsonPair struct {
	ChainId     string `json:"chainId"`
	DexId       string `json:"dexId"`
	Url         string `json:"url"`
	PairAddress string `json:"pairAddress"`
	BaseToken   token  `json:"baseToken"`
	QuoteToken  token  `json:"quoteToken"`
	PriceNative string `json:"priceNative"`
	PriceUsd    string `json:"priceUsd"`
}

// Структура, содержащая список пар
type jsonPairs struct {
	Pairs 	[]jsonPair `json:"pairs"`
}


// API data sample:
// {
//     "schemaVersion": "1.0.0",
//     "pairs": [
//         {
//             "chainId": "ton",
//             "dexId": "stonfi",
//             "url": "https://dexscreener.com/ton/eqcay8ifl2s6lrbmbjey35liumxpc8jfitwg4tl7lbgrsor2",
//             "pairAddress": "EQCaY8Ifl2S6lRBMBJeY35LIuMXPc8JfItWG4tl7lBGrSoR2",
//             "baseToken": {
//                 "address": "EQAvlWFDxGF2lXm67y4yzC17wYKD9A0guwPkMs1gOsM__NOT",
//                 "name": "Notcoin",
//                 "symbol": "NOT"
//             },
//             "quoteToken": {
//                 "address": "EQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAM9c",
//                 "name": "TON",
//                 "symbol": "TON"
//             },
//             "priceNative": "0.001747",
//             "priceUsd": "0.01127",
//             "txns": {
//                 "m5": {
//                     "buys": 35,
//                     "sells": 30
//                 },
//                 "h1": {
//                     "buys": 362,
//                     "sells": 391
//                 },
//                 "h6": {
//                     "buys": 2946,
//                     "sells": 2335
//                 },
//                 "h24": {
//                     "buys": 5792,
//                     "sells": 5968
//                 }
//             },
//             "volume": {
//                 "h24": 577180.74,
//                 "h6": 177374.56,
//                 "h1": 35027.98,
//                 "m5": 1849.99
//             },
//             "priceChange": {
//                 "m5": -0.07,
//                 "h1": 2.8,
//                 "h6": 3.42,
//                 "h24": -2.08
//             },
//             "liquidity": {
//                 "usd": 7262696.53,
//                 "base": 322206116,
//                 "quote": 562999
//             },
//             "fdv": 1155101000,
//             "pairCreatedAt": 1715808468000,
//             "info": {
//                 "imageUrl": "https://dd.dexscreener.com/ds-data/tokens/ton/eqavlwfdxgf2lxm67y4yzc17wykd9a0guwpkms1gosm__not.png",
//                 "websites": [
//                     {
//                         "label": "Website",
//                         "url": "https://notco.in"
//                     }
//                 ],
//                 "socials": [
//                     {
//                         "type": "twitter",
//                         "url": "https://twitter.com/thenotcoin"
//                     },
//                     {
//                         "type": "telegram",
//                         "url": "https://t.me/notcoin"
//                     }
//                 ]
//             }
//         },
//         ...
//     ]
// }
