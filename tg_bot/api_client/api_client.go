package api_client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	customErrors "github.com/ej-you/HamstersShaver/tg_bot/errors"
	"github.com/ej-you/HamstersShaver/tg_bot/settings"
)


// структура для query-параметров для GET-запросов
type QueryParams struct {
	Params map[string]string
}


// получение данных в структуру outStruct по GET-запросу на apiPath с query-параметрами params
func GetRequest(apiPath string, params *QueryParams, outStruct any, timeout time.Duration) error {
	client := &http.Client{Timeout: timeout}

	// обращение к API
	req, err := http.NewRequest("GET", settings.RestApiHost+apiPath, nil)
	if err != nil {
		internalErr := customErrors.InternalError("failed to create GET-request")
	    return fmt.Errorf("send GET-request to %q: %w", apiPath, fmt.Errorf("%v: %w", err, internalErr))
	}
	// добавление query-параметров
	queryParams := req.URL.Query()
	queryParams.Add("api-key", settings.RestApiKey)
	if params != nil {
		for k, v := range params.Params {
			queryParams.Add(k, v)
		}
	}
	req.URL.RawQuery = queryParams.Encode()

	// отправка запроса
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to get account jettons: %v", err)
	}
	defer resp.Body.Close()

	// успешный запрос
	if resp.StatusCode == 200 {
		// декодирование ответа в структуру
		if err := json.NewDecoder(resp.Body).Decode(outStruct); err != nil {
			internalErr := customErrors.InternalError("failed to decode answer from GET-request")
		    return fmt.Errorf("send GET-request to %q: %w", apiPath, fmt.Errorf("%v: %w", err, internalErr))
		}
		return nil
	}

	// парсинг ответа с ошибкой в RestAPIError ошибку
	return fmt.Errorf("send GET-request to %q: got %d response: %w", apiPath, resp.StatusCode, parseError(resp))
}
