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
GO_PORT=8000
```


### Used tools (for TON interaction):

1. `Stonfi` API ([Swagger link](https://api.ston.fi/swagger-ui/))
2. SDK `tonapi-go` ([Github link](https://github.com/tonkeeper/tonapi-go))
3. SDK `tongo` ([Github link](https://github.com/tonkeeper/tongo))
