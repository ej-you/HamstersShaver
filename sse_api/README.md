# HamstersShaver

## SSE API for TON API interaction


### Needed `.env` variables:

```dotenv
# wallet address
TON_API_WALLET_HASH="sample4ch9wko3g3rkjowfw3lpgfkejg5h49eomi45g"

REST_API_TON_API_PORT=8000

# comma-separated allowed origins for CORS
SSE_API_TON_API_CORS_ALLOWED_ORIGINS="*"
# comma-separated allowed methods for CORS
SSE_API_TON_API_CORS_ALLOWED_METHODS="GET,HEAD,POST"

# TON API key for SSE requests (received from TON Console)
SSE_API_TON_API_TOKEN="F4WMGCSOMEV3K5APINOH34FKEY5TDDMQ8WH5"

```

<hr>

### Authorization: none

### Endpoints:
1. `/sse/acount-traces` - subscribe to account (with TON_API_WALLET_HASH wallet) traces

#### Docs:

##### In: nothing
##### Out (trace):

1. Transaction hash:
```
event: trace
data: c3a5bc8a6e78a711150f99d785488a896f9d471039591991518e6b99bc51f332
id: b86f5fff-b878-4989-a209-0c45e6369316
retry: 1000
```
2. Error (after this message the connection will be closed)
```
event: error
data: {
    "status": "error",
    "errors": {
        "sseError": "some error desc"
    }
}
id: b86f5fff-b878-4989-a209-0c45e6369316
retry: 1000
```


<hr>

### Used tools (for TON interaction):

1. SDK `tonapi-go` ([Github link](https://github.com/tonkeeper/tonapi-go))

Also use TON API Key from [TON Console](https://tonconsole.com/tonapi/api-keys) for SSE funcs

<hr>
