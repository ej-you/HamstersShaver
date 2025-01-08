package api_client

import (
	"bytes"
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

const getRequestTimeout = 10*time.Second
const postRequestTimeout = 10*time.Second
const sseRequestTimeout = 6*time.Minute


// структура для query-параметров для GET-запросов
type QueryParams struct {
	// поддерживает значения типов string, int, float64
	Params map[string]interface{}
}

// структура для JSON-body для POST-запросов
type JsonBody struct {
	Data map[string]interface{}
	// Amount		float64 `json:"amount"`
	// JettonCA 	string `json:"jettonCA"`
	// Slippage 	int `json:"slippage"`
}


// отправка запроса и обработка ответа с указанием timeout времени на запрос
func sendRequest(req *http.Request, method, apiPath string, outStruct any, timeout time.Duration) error {
	var err error
	client := &http.Client{Timeout: timeout}

	// добавление query-параметра - API ключа
	queryParams := req.URL.Query()
	queryParams.Add("api-key", settings.RestApiKey)
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
			return fmt.Errorf("send %s-request to %q: %v: %w", method, apiPath, err, internalErr)
		}
		defer resp.Body.Close()

		if resp.StatusCode / 100 == 2 { // 2xx код
			break
		}
		// парсинг ответа с ошибкой в RestAPIError (или RestAPITimeoutError) ошибку
		apiErr = fmt.Errorf("send %s-request to %q: got %d response: %w", method, apiPath, resp.StatusCode, parseError(resp))

		// если полученная ошибка не timeout ошибка, то возвращаем её
		if !errors.As(apiErr, restTimeoutErr) {
			return apiErr
		}
	}
	// если все sendRequestAttemps попытки были неудачными
	if apiErr != nil {
		return apiErr
	}

	// если для ответа не передана структура, то пропускаем парсинг ответа в неё
	if outStruct == nil {
		return nil
	}
	// при успешном запросе - декодирование ответа в структуру
	if err := json.NewDecoder(resp.Body).Decode(outStruct); err != nil {
		internalErr := customErrors.InternalError("failed to decode answer from request")
	    return fmt.Errorf("send %s-request to %q: %w", method, apiPath, fmt.Errorf("%v: %w", err, internalErr))
	}
	return nil
}


// получение данных в структуру outStruct по GET-запросу на apiPath с query-параметрами params
func GetRequest(apiPath string, params *QueryParams, outStruct any) error {
	// обращение к API
	req, err := http.NewRequest("GET", settings.RestApiHost+apiPath, nil)
	if err != nil {
		internalErr := customErrors.InternalError("failed to create GET-request")
	    return fmt.Errorf("send GET-request to %q: %w", apiPath, fmt.Errorf("%v: %w", err, internalErr))
	}
	// добавление query-параметров
	queryParams := req.URL.Query()
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

	// отправка запроса и обработка ответа
	return sendRequest(req, "GET", apiPath, outStruct, getRequestTimeout)
}


// получение данных в структуру outStruct после отправки данных body POST-запросом на apiPath
func PostRequest(apiPath string, body *JsonBody, outStruct any) error {
	// перевод JSON-body в байты
	bytesBody, err := json.Marshal(body.Data)
	if err != nil {
		internalErr := customErrors.InternalError(fmt.Sprintf("failed to marshal json-body"))
		return fmt.Errorf("send POST-request to %q: %v: %w", apiPath, err, internalErr)
	}

	// обращение к API
	req, err := http.NewRequest("POST", settings.RestApiHost+apiPath, bytes.NewReader(bytesBody))
	if err != nil {
		internalErr := customErrors.InternalError("failed to create POST-request")
	    return fmt.Errorf("send POST-request to %q: %w", apiPath, fmt.Errorf("%v: %w", err, internalErr))
	}

	// установка типа контента JSON для данных в body
	req.Header.Add("Content-Type", "application/json")

	// отправка запроса и обработка ответа
	return sendRequest(req, "POST", apiPath, outStruct, postRequestTimeout)
}


// обращение к apiPath и получение данных в структуру outStruct после продолжительного ожидания
func SseRequest(apiPath string, outStruct any) error {
	// обращение к API
	req, err := http.NewRequest("GET", settings.RestApiHost+apiPath, nil)
	if err != nil {
		internalErr := customErrors.InternalError("failed to create GET-request")
	    return fmt.Errorf("send GET-request to %q: %w", apiPath, fmt.Errorf("%v: %w", err, internalErr))
	}

	// отправка запроса и обработка ответа
	return sendRequest(req, "GET", apiPath, outStruct, sseRequestTimeout)
}
