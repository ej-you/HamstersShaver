package main

import (
	"fmt"
	"os"

	echo "github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"

	coreErrorHandler "github.com/Danil-114195722/HamstersShaver/core/error_handler"
	coreUrls "github.com/Danil-114195722/HamstersShaver/core/urls"
	"github.com/Danil-114195722/HamstersShaver/settings"
)


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

	// добавление middleware для проверки API Key в строке запроса
	echoApp.Use(echoMiddleware.KeyAuthWithConfig(echoMiddleware.KeyAuthConfig{
		KeyLookup: "query:api-key",
		Validator: func(key string, context echo.Context) (bool, error) {
			return key == settings.RestApiKey, nil
		},
		ErrorHandler: coreErrorHandler.CustomApiKeyErrorHandler,
	}))

	// настройка кастомного обработчика ошибок
	coreErrorHandler.CustomErrorHandler(echoApp)
	// настройка роутеров для эндпоинтов
	coreUrls.InitUrlRouters(echoApp)

	// запуск приложения
	echoApp.Logger.Fatal(echoApp.Start(fmt.Sprintf(":%s", settings.Port)))
}
