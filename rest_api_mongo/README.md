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
REST_API_MONGO_CORS_ALLOWED_METHODS="GET,HEAD,POST,PATCH"

```


<hr>

### Endpoints:

#### - jettons
1. GET `/api/jettons/get-many` - returns a list of jettons by filter
2. GET `/api/jettons/get-one` - returns one jetton by filter
3. POST `/api/jettons/get-many` - insert jetton into mongo
4. DELETE `/api/jettons/get-many` - delete a list of jettons by filter



### Swagger Docs can be found at `/api/swagger`

### Authorization use header "Authorization" and must be like"
```
Authorization: apiKey your-own-key-for-this-app
```
