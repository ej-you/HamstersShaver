package api_client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	customErrors "github.com/ej-you/HamstersShaver/tg_bot/errors"
	"github.com/ej-you/HamstersShaver/tg_bot/settings"
)


const sendRequestAttemps = 3


// структура для query-параметров для GET-запросов
type QueryParams struct {
	// поддерживает значения типов string, int, float64
	Params map[string]interface{}
}


// получение данных в структуру outStruct по GET-запросу на apiPath с query-параметрами params
func GetRequest(apiPath string, params *QueryParams, outStruct any) error {
	client := &http.Client{Timeout: 10*time.Second}

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
			// конвертация всех типов в string значения
			switch v := v.(type) {
				case string:
					queryParams.Add(k, v)
				case int:
					queryParams.Add(k, fmt.Sprint(v))
				case float64:
					queryParams.Add(k, strconv.FormatFloat(v, 'f', -1, 64))
				default:
					internalErr := customErrors.InternalError(fmt.Sprintf("failed to add query-param %q", k))
					return fmt.Errorf("send GET-request to %q: %w", apiPath, fmt.Errorf("unsupported query-param value type %T: %w", v, internalErr))
			}
		}
	}
	req.URL.RawQuery = queryParams.Encode()

	var resp *http.Response
	restTimeoutErr := new(customErrors.RestAPITimeoutError)
	var apiErr error
	// отправляем запрос и, если получаем timeout ошибку, то пытаемся ещё sendRequestAttemps-1 раз
	for i := 0; i < sendRequestAttemps; i++ {
		// отправка запроса
		resp, err = client.Do(req)
		if err != nil {
			internalErr := customErrors.InternalError(fmt.Sprintf("failed to do request"))
			return fmt.Errorf("failed to get account jettons: %v: %w", err, internalErr)
		}
		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			break
		}
		// парсинг ответа с ошибкой в RestAPIError (или RestAPITimeoutError) ошибку
		apiErr = fmt.Errorf("send GET-request to %q: got %d response: %w", apiPath, resp.StatusCode, parseError(resp))

		// если полученная ошибка не timeout ошибка, то возвращаем её
		if !errors.As(apiErr, restTimeoutErr) {
			return apiErr
		}
	}
	// если все sendRequestAttemps попытки были неудачными
	if apiErr != nil {
		return apiErr
	}

	// при успешном запросе - декодирование ответа в структуру
	if err := json.NewDecoder(resp.Body).Decode(outStruct); err != nil {
		internalErr := customErrors.InternalError("failed to decode answer from GET-request")
	    return fmt.Errorf("send GET-request to %q: %w", apiPath, fmt.Errorf("%v: %w", err, internalErr))
	}
	return nil
}
