package error_handler

import (
	"errors"
	"time"

	echo "github.com/labstack/echo/v4"

	coreErrors "github.com/ej-you/HamstersShaver/rest_api/core/errors"
	"github.com/ej-you/HamstersShaver/rest_api/settings"
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
		// иначе приводим ошибку к APIError
		} else {
			apiErr := coreErrors.AssertAPIError(err)
			errMessage = ResponseError{
				Status: apiErr.ErrStatus,
				StatusCode: apiErr.ErrCode,
				Path: ctx.Path(),
				Timestamp: time.Now().Format(settings.TimeFmt),
				Errors: map[string]string{apiErr.ErrType: apiErr.Description},
			}
		}

		// отправка ответа
		respErr := ctx.JSON(httpError.Code, errMessage)
		if respErr != nil {
			settings.ErrorLog.Println("failed to send error response:", respErr)
		}
	}
}
