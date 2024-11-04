package main

import (
	"fmt"
	"os"
	// "errors"

	echo "github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/ej-you/HamstersShaver/rest_api/docs"

	coreErrorHandler "github.com/ej-you/HamstersShaver/rest_api/core/error_handler"
	coreUrls "github.com/ej-you/HamstersShaver/rest_api/core/urls"
	"github.com/ej-you/HamstersShaver/rest_api/settings"
)

// Настройка Swagger документации
//	@Title						RESTful API for TON API interaction
//	@Version					1.0
//	@Description				RESTful API for TON API interaction written on Golang using "Stonfi" API, SDK "tonapi-go" and SDK "tongo".
//	@Description				All resources is protected with api-key in query.
//	@Host						127.0.0.1:8000
//	@BasePath					/api
//	@Schemes					http
//	@Accept						json
//	@Produce					json
//	@SecurityDefinitions.apiKey	ApiKeyAuth
//	@In							query
//	@Name						api-key
//	@Description				Security api key. Please add it to URL like "?api-key=5how45gi54yi" to authorize your requests.
func main() {
	echoApp := echo.New()
	echoApp.HideBanner = true

	// если при запуске указан аргумент "dev"
	args := os.Args
	if len(args) > 1 {
		// запуск в dev режиме
		if args[1] == "dev" {
			echoApp.Debug = true
		}
	}

	// удаление последнего слеша
	echoApp.Pre(echoMiddleware.RemoveTrailingSlash())
	// кастомизация логирования
	echoApp.Use(echoMiddleware.LoggerWithConfig(echoMiddleware.LoggerConfig{
		Format: settings.LogFmt,
	}))
	// отлавливание паник для беспрерывной работы сервиса
	echoApp.Use(echoMiddleware.Recover())

	// настройка кастомного обработчика ошибок
	coreErrorHandler.CustomErrorHandler(echoApp)
	// настройка Swagger документации
	echoApp.GET("/api/swagger/*", echoSwagger.WrapHandler)

	// создание группы для ресурсов, защищённых API-ключом
	apiKeyProtected := echoApp.Group("/api")

	// добавление middleware для проверки API Key в строке запроса
	apiKeyProtected.Use(echoMiddleware.KeyAuthWithConfig(echoMiddleware.KeyAuthConfig{
		KeyLookup: "query:api-key",
		Validator: func(key string, context echo.Context) (bool, error) {
			// для более простой отладки делаем API-ключ "debug" доступным для авторизации
			if echoApp.Debug {
				return key == settings.RestApiKey || key == "debug", nil
			}
			return key == settings.RestApiKey, nil
		},
		ErrorHandler: coreErrorHandler.CustomApiKeyErrorHandler,
	}))

	// настройка роутеров для эндпоинтов
	coreUrls.InitUrlRouters(apiKeyProtected)

	// запуск приложения
	echoApp.Logger.Fatal(echoApp.Start(fmt.Sprintf(":%s", settings.Port)))
}
