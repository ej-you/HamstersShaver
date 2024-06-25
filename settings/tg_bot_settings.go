package settings

import (
	"os"
	"github.com/joho/godotenv"
)


// загрузка переменных окружения
var _ error = godotenv.Load("./settings/config/.env")

// токен бота
var BotToken string = os.Getenv("BOT_TOKEN")

// настройки redis
var redisHost string = os.Getenv("REDIS_HOST")
var redisPort string = os.Getenv("REDIS_PORT")

var RedisAddr string = redisHost + ":" + redisPort
var RedisPassword string = os.Getenv("REDIS_PASSWORD")
