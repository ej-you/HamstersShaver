package main

import (
	"fmt"
	"os"

	echo "github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	openapidocs "github.com/kohkimakimoto/echo-openapidocs"

	"github.com/ej-you/go-utils/env"

	coreErrorHandler "github.com/ej-you/HamstersShaver/rest_api_mongo/core/error_handler"
	coreUrls "github.com/ej-you/HamstersShaver/rest_api_mongo/core/urls"
	"github.com/ej-you/HamstersShaver/rest_api_mongo/settings"
)

// Настройка Swagger документации
// @Version 1.0.0
// @Title RESTful API for MongoDB
// @Description RESTful API for MongoDB written on Golang. All resources is protected with api-key in header.
// @Server http://150.241.82.68:8002/api Remote server
// @Server http://127.0.0.1:8002/api Local machine
// @SecurityScheme APIKey apiKey header Authorization
// @Security APIKey
func main() {
	// проверка, что эти переменные окружения заданы
	env.MustBePresented(
		"REST_API_MONGO_PORT", "MY_APIS_KEY",
		"REST_API_MONGO_CORS_ALLOWED_ORIGINS", "REST_API_MONGO_CORS_ALLOWED_METHODS",
		"MONGO_HOST", "MONGO_PORT", "MONGO_DB",
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
	coreErrorHandler.CustomErrorHandler(echoApp)

	// настройка Swagger документации
	echoApp.File("/api/docs/swagger_v3.yml", "./docs/swagger_v3.yml")
	echoApp.File("/favicon.ico", "./docs/favicon.ico")
	openapidocs.SwaggerUIDocuments(echoApp, "/api/swagger", openapidocs.SwaggerUIConfig{
		SpecUrl: "/api/docs/swagger_v3.yml",
		Title:   "REST API for TON API",
	})

	// создание группы для ресурсов, защищённых API-ключом
	apiKeyProtected := echoApp.Group("/api")

	// добавление middleware для проверки API Key в заголовках запроса
	apiKeyProtected.Use(echoMiddleware.KeyAuthWithConfig(echoMiddleware.KeyAuthConfig{
		AuthScheme: "apiKey",
		Validator: func(key string, context echo.Context) (bool, error) {
			// для более простой отладки делаем API-ключ "debug" доступным для авторизации
			if echoApp.Debug {
				return key == settings.MyApisKey || key == "debug", nil
			}
			return key == settings.MyApisKey, nil
		},
		ErrorHandler: coreErrorHandler.CustomApiKeyErrorHandler,
	}))

	// настройка CORS
	echoApp.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: settings.CorsAllowedOrigins,
		AllowMethods: settings.CorsAllowedMethods,
	}))

	// настройка роутеров для эндпоинтов
	coreUrls.InitUrlRouters(apiKeyProtected)

	// запуск приложения
	echoApp.Logger.Fatal(echoApp.Start(fmt.Sprintf(":%s", settings.Port)))
}
