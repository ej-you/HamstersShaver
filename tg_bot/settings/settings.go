package settings

import (
	"fmt"
	"log"
	"os"
	"strings"
)


// токен бота
var BotToken string = os.Getenv("TG_BOT_TOKEN")
// список ID юзеров с доступом к боту
var AllowedUsers []string = strings.Split(os.Getenv("TG_BOT_ALLOWED_USERS"), ",")

// настройки redis
var RedisAddr string = os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")

// настройки mongo
var MongoAddr string = fmt.Sprintf("mongodb://%s:%s/", os.Getenv("MONGO_HOST"), os.Getenv("MONGO_PORT"))
var MongoDB string = os.Getenv("MONGO_DB")

// настройки для REST API
var RestApiHost string = os.Getenv("REST_API_TON_API_HOST")
var RestApiKey string = os.Getenv("MY_APIS_KEY")


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
