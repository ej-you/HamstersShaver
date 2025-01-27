package error_handler

import (
	"errors"

	echo "github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)


// настройка обработчика ошибок для JWT middleware
func CustomApiKeyErrorHandler(err error, context echo.Context) error {
	// токен не был отправлен в строке запроса
	apiKeyParsingError := new(echoMiddleware.ErrKeyAuthMissing)
	if errors.As(err, &apiKeyParsingError) {
		return echo.NewHTTPError(400, map[string]string{"api_key": apiKeyParsingError.Error()})
	}

	// неверный API key
	return echo.NewHTTPError(401, map[string]string{"api_key": err.Error()})
}
