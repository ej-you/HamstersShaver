package settings

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)


// путь до .env файла
var envFilePath = func() string {
	return os.Getenv("ENV_FILE_PATH") + "./.env"
}()

// загрузка переменных окружения
var _ error = godotenv.Load(envFilePath)


// токен бота
var BotToken string = os.Getenv("BOT_TOKEN")

// список ID юзеров с доступом к боту
var AllowedUsers []string = strings.Split(os.Getenv("ALLOWED_USERS"), ",")

// настройки для REST API
var RestApiHost string = os.Getenv("REST_API_HOST")
var RestApiKey string = os.Getenv("REST_API_KEY")

// настройки redis
var RedisAddr string = os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")
var RedisPassword string = os.Getenv("REDIS_PASSWORD")

// логеры
var InfoLog *log.Logger = log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
var ErrorLog *log.Logger = log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime)
var fatalLog *log.Logger = log.New(os.Stderr, "[FATAL]\t", log.Ldate|log.Ltime|log.Lshortfile)

// функция для обработки критических ошибок
func DieIf(err error) {
	if err != nil {
		fatalLog.Panic(err)
	}
}
