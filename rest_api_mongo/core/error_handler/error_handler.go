package error_handler

import (
	"errors"
	"time"

	echo "github.com/labstack/echo/v4"

	validatorModule "github.com/go-playground/validator/v10"
	myValidatorModule "github.com/ej-you/go-utils/validator"
	coreValidator "github.com/ej-you/HamstersShaver/rest_api_mongo/core/validator"

	"github.com/ej-you/HamstersShaver/rest_api_mongo/settings"
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
		errMessage := ResponseError{
			Path: ctx.Path(),
			Timestamp: time.Now().Format(settings.TimeFmt),
		}

		// проверка на ошибки валидации
		validateErrors, ok := err.(validatorModule.ValidationErrors)
		if ok {
			errMessage.Status = "validateError"
			errMessage.StatusCode = 400
			errMessage.Errors = myValidatorModule.GetTranslatedMap(validateErrors, coreValidator.GetTranslator())
			sendErrorResponse(&ctx, &errMessage)
			return
		}

		// если ошибка является структурой *echo.HTTPError
		httpError := new(echo.HTTPError)
		if errors.As(err, &httpError) {
			// приведение httpError.Message типа interface{} к map[string]string
			errorsInMessage, ok := httpError.Message.(map[string]string)
			if !ok {
				errorsInMessage = map[string]string{"unknown": httpError.Error()}
			}
			errMessage.Status = "error"
			errMessage.StatusCode = httpError.Code
			errMessage.Errors = errorsInMessage
			sendErrorResponse(&ctx, &errMessage)
			return
		}

		// иначе
		errMessage.Status = "error"
		errMessage.StatusCode = 500
		errMessage.Errors = map[string]string{"unknown": err.Error()}
		sendErrorResponse(&ctx, &errMessage)
	}
}


// отправка ответа с сообщением об ошибке
func sendErrorResponse(ctx *echo.Context, errMessage *ResponseError) {
	respErr := (*ctx).JSON((*errMessage).StatusCode, *errMessage)
	if respErr != nil {
		settings.ErrorLog.Println("failed to send error response:", respErr)
	}
}
