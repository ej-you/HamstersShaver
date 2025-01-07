# HamstersShaver

## RESTful API for TON API interaction


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
REST_API_KEY="your-own-key-for-this-app"

# comma-separated allowed origins for CORS
CORS_ALLOWED_ORIGINS="*"
# comma-separated allowed methods for CORS
CORS_ALLOWED_METHODS="GET,HEAD,POST"

# received from TON Console
TON_API_TOKEN="F4WMGCSOMEV3K5APINOH34FKEY5TDDMQ8WH5"

```

<hr>

### Endpoints:

#### - account
1. `/api/acount/get-jettons` - returns list with info about each account jetton (exclude TON)
2. `/api/acount/get-jetton` - returns info about account jetton by its address
3. `/api/acount/get-ton` - returns info about account's TON
4. `/api/acount/get-seqno` - returns account seqno

#### - jetton
1. `/api/jettons/get-info` - returns info about jetton by its address

#### - transactions
1. `/api/transactions/buy/pre-request` - returns info about pre-request buy transaction
2. `/api/transactions/buy/send` - send buy transaction to blockchain
3. `/api/transactions/cell/pre-request` - returns info about pre-request cell transaction
4. `/api/transactions/cell/send` - send cell transaction to blockchain
5. `/api/transactions/info` - get transaction info by its hash
6. `/api/transactions/wait-next` - wait the end of next transaction

#### - services
1. `/api/services/beauty-balance` - returns rounded string balance converted from given int64 balance


### Swagger Docs can be found at `/api/swagger/index.html`

<hr>

### Used tools (for TON interaction):

1. `Stonfi` API ([Swagger link](https://api.ston.fi/swagger-ui/))
2. SDK `tonapi-go` ([Github link](https://github.com/tonkeeper/tonapi-go))
3. SDK `tongo` ([Github link](https://github.com/tonkeeper/tongo))
4. SDK `tonutils-go` ([Github link](https://github.com/xssnick/tonutils-go))

Also use TON API Key from [TON Console](https://tonconsole.com/tonapi/api-keys) for SSE funcs

<hr>
