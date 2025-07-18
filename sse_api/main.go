package main

import (
	"fmt"
	"os"

	echo "github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/ej-you/go-utils/env"

	"github.com/ej-you/HamstersShaver/sse_api/handlers"
	errorHandler "github.com/ej-you/HamstersShaver/sse_api/error_handler"
	"github.com/ej-you/HamstersShaver/sse_api/settings"
)


func main() {
	// проверка, что эти переменные окружения заданы
	env.MustBePresented(
		"TON_API_WALLET_HASH",
		"SSE_API_TON_API_PORT",
		"SSE_API_TON_API_CORS_ALLOWED_ORIGINS", "SSE_API_TON_API_CORS_ALLOWED_METHODS",
		"SSE_API_TON_API_TOKEN",
	)

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
	// отлавливание паник для беспрерывной работы сервиса (если приложение запущено не в debug режиме)
	if !echoApp.Debug {
		echoApp.Use(echoMiddleware.Recover())
	}
	
	// настройка кастомного обработчика ошибок
	errorHandler.CustomErrorHandler(echoApp)

	// настройка CORS
	echoApp.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: settings.CorsAllowedOrigins,
		AllowMethods: settings.CorsAllowedMethods,
	}))

	// настройка роутеров для эндпоинтов
	echoApp.GET("/sse/account-traces", handlers.SubscribeToAccountTraces)

	// запуск приложения
	echoApp.Logger.Fatal(echoApp.Start(fmt.Sprintf(":%s", settings.Port)))
}
