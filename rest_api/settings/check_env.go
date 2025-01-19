package settings

import "fmt"


// проверка, что все переменные окружения были предоставлены
func CheckEnv() {
	if (
		hash == "" ||
		seedPhrase == "" ||

		Port == "" ||
		RestApiKey == "" ||

		CorsAllowedOrigins == nil ||
		CorsAllowedMethods == nil ||

		TonApiToken == "") {
		// if body
		panic(fmt.Errorf("Not all env variables is presented!"))
	}
}
