package error_handler

import (
	"errors"
	"time"

	echo "github.com/labstack/echo/v4"

	"github.com/ej-you/HamstersShaver/sse_api/settings"
)


type ResponseError struct {
	Status 		string `json:"status"`
	StatusCode 	int `json:"statusCode"`
	Path 		string `json:"path"`
	Timestamp 	string `json:"timestamp"`
	Errors 		map[string]string `json:"errors"`
}


// настройка обработчика ошибок
func CustomErrorHandler(echoApp *echo.Echo) {
	echoApp.HTTPErrorHandler = func(err error, ctx echo.Context) {
		var errMessage ResponseError
		var httpErrorStatus int

		// если ошибка является структурой *echo.HTTPError
		httpError := new(echo.HTTPError)
		if errors.As(err, &httpError) {
			// приведение httpError.Message типа interface{} к map[string]string
			errorsInMessage, ok := httpError.Message.(map[string]string)
			if !ok {
				errorsInMessage = map[string]string{"unknown": httpError.Error()}
			}
			errMessage = ResponseError{
				Status: "error",
				StatusCode: httpError.Code,
				Path: ctx.Path(),
				Timestamp: time.Now().Format(settings.TimeFmt),
				Errors: errorsInMessage,
			}
			httpErrorStatus = httpError.Code
		} else {
			errMessage = ResponseError{
				Status: "error",
				StatusCode: 500,
				Path: ctx.Path(),
				Timestamp: time.Now().Format(settings.TimeFmt),
				Errors: map[string]string{"unknown": err.Error()},
			}
		}

		// отправка ответа
		respErr := ctx.JSON(httpErrorStatus, errMessage)
		if respErr != nil {
			settings.ErrorLog.Println("failed to send error response:", respErr)
		}
	}
}
