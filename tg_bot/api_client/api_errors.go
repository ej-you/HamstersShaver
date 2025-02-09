package api_client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	customErrors "github.com/ej-you/HamstersShaver/tg_bot/errors"
)


// структура для парсинга JSON-ответа от API с ошибкой
type rawRestAPIError struct {
	Status 	string `json:"status"`
	StatusCode 	int `json:"statusCode"`
	Path 	string `json:"path"`
	Errors 		map[string]string `json:"errors"`
}

func parseError(response *http.Response) error {
	var rawErr rawRestAPIError

	// декодирование ответа с ошибкой в структуру
	if err := json.NewDecoder(response.Body).Decode(&rawErr); err != nil {
	    return fmt.Errorf("%v: %w", err, customErrors.InternalError("failed to decode error answer from GET-request"))
	}

	// объединение словаря ошибок в строку ошибок с кодом ответа API
	errorsString := ""
	for errKey, errValue := range rawErr.Errors {
		errorsString += fmt.Sprintf("%s: %s && ", errKey, errValue)
	}
	errorsString = fmt.Sprintf("code %d: %s", rawErr.StatusCode, strings.TrimSuffix(errorsString, " && "))

	// если timeout ошибка - возврат ошибки RestAPITimeoutError
	if rawErr.Status == "timeout" {
		return customErrors.RestAPITimeoutError(errorsString)
	}
	// возврат ошибки RestAPIError
	return customErrors.RestAPIError(errorsString)
}
