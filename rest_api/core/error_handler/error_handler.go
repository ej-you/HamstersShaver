package error_handler

import (
	"errors"

	echo "github.com/labstack/echo/v4"

	validatorModule "github.com/go-playground/validator/v10"
	myValidatorModule "github.com/ej-you/go-utils/validator"
	coreValidator "github.com/ej-you/HamstersShaver/rest_api/core/validator"

	coreErrors "github.com/ej-you/HamstersShaver/rest_api/core/errors"
	"github.com/ej-you/HamstersShaver/rest_api/settings"
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

		// проверка на ошибки валидации
		validateErrors, ok := err.(validatorModule.ValidationErrors)
		if ok {
			errMessage.StatusCode = 400
			errMessage.Status = "validateError"
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
			errMessage.StatusCode = httpError.Code
			errMessage.Status = "error"
			errMessage.Errors = errorsInMessage
			sendErrorResponse(&ctx, &errMessage)
			return
		}

		// иначе приводим ошибку к APIError
		apiErr := coreErrors.AssertAPIError(err)
		errMessage.StatusCode = apiErr.ErrCode
		errMessage.Status = apiErr.ErrStatus
		errMessage.Errors = map[string]string{apiErr.ErrType: apiErr.Description}
		sendErrorResponse(&ctx, &errMessage)
	}
}


// отправка ответа с сообщением об ошибке
func sendErrorResponse(ctx *echo.Context, errMessage *ResponseError) {
	// логируем ошибку в STDERR
	settings.ErrorLog.Printf("Path: %v | Error: %#v", (*ctx).Path(), *errMessage)

	// отправляем ошибку клиенту
	respErr := (*ctx).JSON((*errMessage).StatusCode, *errMessage)
	if respErr != nil {
		settings.ErrorLog.Println("failed to send error response:", respErr)
	}
}
