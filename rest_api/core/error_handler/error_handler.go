package error_handler

import (
	"errors"
	"time"

	echo "github.com/labstack/echo/v4"

	"github.com/ej-you/HamstersShaver/rest_api/settings"
)


// настройка обработчика ошибок
func CustomErrorHandler(echoApp *echo.Echo) {
	echoApp.HTTPErrorHandler = func(err error, ctx echo.Context) {
		// является ли ошибка структурой *echo.HTTPError
		httpError := new(echo.HTTPError)
		if !errors.As(err, &httpError) {
			httpError = echo.NewHTTPError(500, map[string]string{"unknown": err.Error()})
		}
		// если пришла *echo.HTTPError ошибка со строкой в качестве httpError.Message
		stringErrorMessage, ok := (httpError.Message).(string)
		if ok {
			httpError.Message = map[string]string{"unknown": stringErrorMessage}
		}

		errMessage := map[string]interface{}{
			"status": "error",
			"statusCode": httpError.Code,
			"path": ctx.Path(),
			"timestamp": time.Now().Format(settings.TimeFmt),
			"errors": httpError.Message,
		}

		// отправка ответа
		respErr := ctx.JSON(httpError.Code, errMessage)
		if respErr != nil {
			settings.ErrorLog.Println("Failed to send error response:", respErr)
		}
	}
}
