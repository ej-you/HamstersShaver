# HamstersShaver

## SSE API for TON API interaction


### Needed `.env` variables:

```dotenv
# wallet address
TON_API_WALLET_HASH="sample4ch9wko3g3rkjowfw3lpgfkejg5h49eomi45g"

# TON API key for SSE requests (received from TON Console)
SSE_API_TON_API_TOKEN="F4WMGCSOMEV3K5APINOH34FKEY5TDDMQ8WH5"

```

<hr>

### Authorization use header "Authorization" and must be like"
```
Authorization: apiKey your-own-key-for-this-app
```

<hr>

### Used tools (for TON interaction):

1. SDK `tonapi-go` ([Github link](https://github.com/tonkeeper/tonapi-go))

Also use TON API Key from [TON Console](https://tonconsole.com/tonapi/api-keys) for SSE funcs

<hr>