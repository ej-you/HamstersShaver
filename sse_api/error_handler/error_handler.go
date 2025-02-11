package error_handler

import (
	"encoding/json"
	"fmt"

	echo "github.com/labstack/echo/v4"

	customErrors "github.com/ej-you/HamstersShaver/sse_api/errors"
	"github.com/ej-you/HamstersShaver/sse_api/settings"
)


type ResponseError struct {
	StatusCode 	int `json:"statusCode"`
	Status 		string `json:"status"`
	Errors 		map[string]string `json:"errors"`
}


// настройка обработчика ошибок
func CustomErrorHandler(echoApp *echo.Echo) {
	echoApp.HTTPErrorHandler = func(err error, ctx echo.Context) {
		// структура для обрабатываемой ошибки любого типа
		var errMessage ResponseError

		// если ошибка является customErrors.SseError ошибкой
		if sseErr, ok := err.(customErrors.SseError); ok {
			errMessage.StatusCode = 500
			errMessage.Status = "sseError"
			errMessage.Errors = map[string]string{"sseError": sseErr.Error()}

			// отправка ответа (для customErrors.SseError ошибки)
			respWriter := ctx.Response()
			respWriter.Header().Set("Content-Type", "text/event-stream")
			respWriter.Header().Set("Cache-Control", "no-cache")
			respWriter.Header().Set("Connection", "keep-alive")
			respWriter.WriteHeader(500)

			byteErrMessage, _ := json.Marshal(errMessage)
			message := fmt.Sprintf("event: error\ndata: %s\n\n", string(byteErrMessage))
			_, _ = fmt.Fprint(respWriter, message)
			respWriter.Flush()

			// логируем ошибку в STDERR
			settings.ErrorLog.Printf("Path: %v | Error: %#v", ctx.Path(), errMessage)
			return

		// если ошибка является структурой *echo.HTTPError
		} else if httpErr, ok := err.(*echo.HTTPError); ok {
			// приведение httpError.Message типа interface{} к map[string]string
			errorsInMessage, ok := httpErr.Message.(map[string]string)
			if !ok {
				errorsInMessage = map[string]string{"unknown": httpErr.Error()}
			}
			errMessage.StatusCode = httpErr.Code
			errMessage.Status = "error"
			errMessage.Errors = errorsInMessage

		// если неизвестная ошибка
		} else {
			errMessage.StatusCode = 500
			errMessage.Status = "unknownError"
			errMessage.Errors = map[string]string{"unknown": err.Error()}
		}

		// логируем ошибку в STDERR
		settings.ErrorLog.Printf("Path: %v | Error: %#v", ctx.Path(), errMessage)

		// отправка ответа (для *echo.HTTPError или неизвестной ошибки)
		respErr := ctx.JSON(errMessage.StatusCode, errMessage)
		if respErr != nil {
			settings.ErrorLog.Println("failed to send error response:", respErr)
		}
	}
}
