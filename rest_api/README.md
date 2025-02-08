# HamstersShaver

## RESTful API for TON API interaction


### Needed `.env` variables:

```dotenv
# wallet address
TON_API_WALLET_HASH="sample4ch9wko3g3rkjowfw3lpgfkejg5h49eomi45g"
# wallet mnemonics
TON_API_WALLET_SEED_PHRASE="your long seed phrase containing twenty four words"

REST_API_TON_API_PORT=8000
MY_APIS_KEY="your-own-key-for-this-app"

# comma-separated allowed origins for CORS
REST_API_TON_API_CORS_ALLOWED_ORIGINS="*"
# comma-separated allowed methods for CORS
REST_API_TON_API_CORS_ALLOWED_METHODS="GET,HEAD,POST"

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

#### - services
1. `/api/services/jetton-amount-from-percent` - returns not rounded jettons amount from percent of its balance
2. `/api/services/ton-amount-from-percent` - returns not rounded TON amount from percent of its balance


### Swagger Docs can be found at `/api/swagger`

### Authorization use header "Authorization" and must be like"
```
Authorization: apiKey your-own-key-for-this-app
```


<hr>

### Used tools (for TON interaction):

1. `Stonfi` API ([Swagger link](https://api.ston.fi/swagger-ui/))
2. SDK `tonapi-go` ([Github link](https://github.com/tonkeeper/tonapi-go))
3. SDK `tongo` ([Github link](https://github.com/tonkeeper/tongo))
4. SDK `tonutils-go` ([Github link](https://github.com/xssnick/tonutils-go))

Also use TON API Key from [TON Console](https://tonconsole.com/tonapi/api-keys) for SSE funcs

<hr>
