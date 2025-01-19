package settings

import (
	"log"
	"os"
	"strings"
)


// даннные кошелька
var hash string = os.Getenv("WALLET_HASH")
var seedPhrase string = os.Getenv("WALLET_SEED_PHRASE")

// данные REST API
var Port string = os.Getenv("REST_API_PORT")
var RestApiKey string = os.Getenv("REST_API_KEY")

// разрешённые источники и методы
var CorsAllowedOrigins []string = strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ",")
var CorsAllowedMethods []string = strings.Split(os.Getenv("CORS_ALLOWED_METHODS"), ",")

// TON API ключ для SSE запросов
var TonApiToken string = os.Getenv("TON_API_TOKEN")


// формат логов (для Echo)
var LogFmt string = "[${time_rfc3339}] -- ${status} -- from ${remote_ip} to ${host} (${method} ${uri}) [time: ${latency_human}] | ${bytes_in} ${bytes_out} | error: ${error} |\n"
// формат времени (для Echo)
var TimeFmt string = "06-01-02 15:04:05 -07"

// логеры
var InfoLog *log.Logger = log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
var ErrorLog *log.Logger = log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)
var fatalLog *log.Logger = log.New(os.Stderr, "[FATAL]\t", log.Ldate|log.Ltime|log.Lshortfile)

// функция для обработки критических ошибок
func DieIf(err error) {
	if err != nil {
		fatalLog.Panic(err)
	}
}
