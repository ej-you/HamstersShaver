package settings

import (
	"fmt"
	"log"
	"os"
	"strings"
)


// данные REST API
var Port string = os.Getenv("REST_API_TON_API_PORT")
var MyApisKey string = os.Getenv("MY_APIS_KEY")

// разрешённые источники и методы
var CorsAllowedOrigins []string = strings.Split(os.Getenv("REST_API_TON_API_CORS_ALLOWED_ORIGINS"), ",")
var CorsAllowedMethods []string = strings.Split(os.Getenv("REST_API_TON_API_CORS_ALLOWED_METHODS"), ",")


// настройки mongo
var MongoAddr string = fmt.Sprintf("mongodb://%s:%s/", os.Getenv("MONGO_HOST"), os.Getenv("MONGO_PORT"))
var MongoDB string = os.Getenv("MONGO_DB")


// формат логов (для Echo)
var LogFmt string = "[${time_rfc3339}] -- ${status} -- from ${remote_ip} to ${host} (${method} ${uri}) [time: ${latency_human}] | ${bytes_in} ${bytes_out} | error: ${error} |\n"
// формат времени (для Echo)
var TimeFmt string = "06-01-02 15:04:05 -07"

// логеры
var InfoLog *log.Logger = log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
var ErrorLog *log.Logger = log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime)
