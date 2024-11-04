package error_handler

import (
	"net/http"
	"reflect"

	echo "github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)


// настройка обработчика ошибок для JWT middleware
func CustomApiKeyErrorHandler(err error, context echo.Context) error {
	// токен не был отправлен в строке запроса
	apiKeyParsingError, ok := err.(*echoMiddleware.ErrKeyAuthMissing)
	if ok {
		httpError := &echo.HTTPError{
			Code: http.StatusBadRequest,
			Message: map[string]string{"apiKey": apiKeyParsingError.Error()},
		}
		return httpError
	}

	// неверный API key
	if reflect.TypeOf(err).String() == "*errors.errorString" {
		httpError := &echo.HTTPError{
			Code: http.StatusUnauthorized,
			Message: map[string]string{"apiKey": err.Error()},
		}
		return httpError
	}

	return err
}
