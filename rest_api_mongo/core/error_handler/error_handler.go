package error_handler

import (
	"errors"

	echo "github.com/labstack/echo/v4"

	coreValidator "github.com/ej-you/HamstersShaver/rest_api_mongo/core/validator"
	"github.com/ej-you/HamstersShaver/rest_api_mongo/settings"
)


type ResponseError struct {
	StatusCode 	int `json:"-"`
	Status 		string `json:"status"`
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
			errMessage.StatusCode = httpError.Code
			errMessage.Status = "error"
			errMessage.Errors = errorsInMessage
			sendErrorResponse(&ctx, &errMessage)
			return
		}

		// проверка на ошибки валидации
		errMap, ok := coreValidator.GetValidator().GetMapFromValidationError(err)
		if ok {
			errMessage.StatusCode = 400
			errMessage.Status = "validateError"
			errMessage.Errors = errMap
			sendErrorResponse(&ctx, &errMessage)
			return
		}

		// иначе
		errMessage.StatusCode = 500
		errMessage.Status = "unknownError"
		errMessage.Errors = map[string]string{"unknown": err.Error()}
		sendErrorResponse(&ctx, &errMessage)
	}
}


// отправка ответа с сообщением об ошибке
func sendErrorResponse(ctx *echo.Context, errMessage *ResponseError) {
	respErr := (*ctx).JSON((*errMessage).StatusCode, *errMessage)
	if respErr != nil {
		settings.ErrorLog.Println("failed to send error response:", respErr)
		return
	}
	// логируем ошибку в STDERR
	settings.ErrorLog.Printf("Path: %v | Error: %#v", (*ctx).Path(), *errMessage)
}
