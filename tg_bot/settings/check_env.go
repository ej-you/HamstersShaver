package settings

import (
	"fmt"
	"os"
)


// проверка, что все переменные окружения были предоставлены
func CheckEnv() {
	if (
		os.Getenv("BOT_TOKEN") == "" ||
		os.Getenv("ALLOWED_USERS") == "" ||

		os.Getenv("REDIS_HOST") == "" ||
		os.Getenv("REDIS_PORT") == "" ||

		os.Getenv("MONGO_HOST") == "" ||
		os.Getenv("MONGO_PORT") == "" ||
		os.Getenv("MONGO_DB") == "" ||

		os.Getenv("REST_API_HOST") == "" ||
		os.Getenv("REST_API_KEY") == "") {
		// if body
		panic(fmt.Errorf("Not all env variables is presented!"))
	}
}
