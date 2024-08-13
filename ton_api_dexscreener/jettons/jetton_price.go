package jettons


// Stonfi API

// {
//   "pool": {
//     "address": "EQCaY8Ifl2S6lRBMBJeY35LIuMXPc8JfItWG4tl7lBGrSoR2",
//     "router_address": "EQB3ncyBUTjZUA5EnFKR5_EnOMI9V1tTEAAPaiU71gc4TiUt",
	//     "reserve0": "327026382098709038",
	//     "reserve1": "562652470746670",
//     "token0_address": "EQAvlWFDxGF2lXm67y4yzC17wYKD9A0guwPkMs1gOsM__NOT",
//     "token1_address": "EQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAM9c",
//     "lp_total_supply": "5619027430185",
//     "lp_total_supply_usd": "7374303.063769328007550320",
//     "lp_fee": "70",
//     "protocol_fee": "30",
//     "ref_fee": "30",
//     "protocol_fee_address": "EQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAM9c",
//     "collected_token0_protocol_fee": "28906878962105608",
//     "collected_token1_protocol_fee": "47919876460478",
//     "lp_price_usd": "1312.380684272",
//     "apy_1d": "0.12829214305895675",
//     "apy_7d": "0.24082791713214402",
//     "apy_30d": "0.24328527026084895",
//     "deprecated": false
//   }
// }

// Цена NOTcoin в TONcoin = reserve1 / reserve0
// Цена NOTcoin в USD = (NOTcoin в TONcoin) * (Цена TONcoin в USD)

// https://api.ston.fi/v1/pools/{pool_address}
// EQCaY8Ifl2S6lRBMBJeY35LIuMXPc8JfItWG4tl7lBGrSoR2



// tongo abi.GetPoolAddress (/abi/get_methods.go)
