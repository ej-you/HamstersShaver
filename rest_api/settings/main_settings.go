package settings

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)


// путь до директории с конфигом
var configPath = func() string {
	fullPath := os.Getenv("CONFIG_PATH")
	// если переменная окружения не задана, то ставим дефолтное значение
	if fullPath == "" {
		return "./settings/config/"
	}
	return fullPath
}()


// загрузка переменных окружения
var _ error = godotenv.Load(configPath + ".env")

// распаковка переменных окружения
var Port string = os.Getenv("GO_PORT")
var RestApiKey string = os.Getenv("REST_API_KEY")

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
