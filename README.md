# HamstersShaver

## TG bot for fast transactions


### Config dir (`./settings/config/`) must contain the next files:

#### 1. `wallet.json`
##### Content:

```json5
{
	// wallet address
	"hash": "sample4ch9wko3g3rkjowfw3lpgfkejg5h49eomi45g",
	// wallet mnemonics
	"seed_phrase": "your long seed phrase containing twenty four words"
}
```

#### 2. `.env`
##### Content:

```dotenv
# API key for cryptocompare.com service API
CRYPTOCOMPARE_API_KEY=sampleg5k7ywg7l5jmg8ho5wdg4ih8yho34htw45e895gpt
```


### Used tools (for TON interaction):

1. API `cryptocompare.com` ([Documentation link](https://min-api.cryptocompare.com/documentation))
2. API `dexscreener.com` ([Documentation link](https://docs.dexscreener.com/api/reference))
3. SDK `tonapi-go` ([Github link](https://github.com/tonkeeper/tonapi-go))
4. SDK `tongo` ([Github link](https://github.com/tonkeeper/tongo))
