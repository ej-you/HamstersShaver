# HamstersShaver

## Telegram bot for fast transactions on TON

#### About REST API for TON API interaction read [here](./rest_api/README.md)

#### File `.env` must be contained in project root
##### Content:

```dotenv
# Bot token (gotten form BotFather)
BOT_TOKEN="7589679:FslNG9krgnk4gihgl3h4MSDK-vjeH4t8"

# comma-separated allowed users' IDs
ALLOWED_USERS="123456789,012345678"

# My REST API
REST_API_HOST="https://domain.com"
# Key for my REST API
REST_API_KEY="your-own-key-for-your-rest-api"

# redis server settings
REDIS_HOST=172.17.0.3
REDIS_PORT=6379

# mongo server settings
MONGO_HOST=172.17.0.2
MONGO_PORT=27017
MONGO_DB="hamsters_shaver"

```
