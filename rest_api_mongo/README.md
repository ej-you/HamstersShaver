# HamstersShaver

## RESTful API for MongoDB


### Needed `.env` variables:

```dotenv
MONGO_HOST=172.17.0.2
MONGO_PORT=27017
MONGO_DB="db_name"

# Key for my REST APIs
MY_APIS_KEY="your-own-key-for-this-app"

# port for my REST API for MongoDB
REST_API_MONGO_PORT=8002

# comma-separated allowed origins for CORS for my REST API for MongoDB
REST_API_MONGO_CORS_ALLOWED_ORIGINS="*"
# comma-separated allowed methods for CORS for my REST API for MongoDB
REST_API_MONGO_CORS_ALLOWED_METHODS="GET,HEAD,POST,PATCH,DELETE"

```


<hr>

### Endpoints:

#### - jettons
1. POST `/api/jettons/get-many` - insert jetton into mongo
2. GET `/api/jettons/get-many` - returns a list of jettons by filter
3. GET `/api/jettons/get-one` - returns one jetton by filter
4. DELETE `/api/jettons/get-many` - delete a list of jettons by filter

#### - transactions
1. POST `/api/transactions/create-trade` - insert "trade" transaction into mongo
2. POST `/api/transactions/create-auto` - insert "auto" transaction into mongo
3. GET `/api/transactions/get-one` - returns one transaction by filter
4. GET `/api/transactions/get-many` - returns a list of transactions by filter
5. PATCH `/api/transactions` - update a list of transactions by filter
6. PATCH `/api/transactions/commit-init-trans` - commit init transaction info into initTrans sub-object


### Swagger Docs can be found at `/api/swagger`

### Authorization use header "Authorization" and must be like"
```
Authorization: apiKey your-own-key-for-this-app
```
