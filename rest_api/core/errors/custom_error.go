package errors

import (
	"errors"
	"context"
)


var RestApiError = errors.New("rest api error") // code 500
var TonApiError = errors.New("ton api error") // code 500
var TimeoutError = errors.New("timeout error") // code 500
var JettonNotFoundError = errors.New("jetton not found") // code 400
var AccountHasNotJettonError = errors.New("account has not given jetton") // code 404


// проверяет, что ошибка из-за таймаута, и возвращает true, если timeout произошёл из-за контекста
func IsTimeout(err error) bool {
	timeoutErr := context.DeadlineExceeded
	
	if errors.Is(err, timeoutErr) {
		return true
	}
	return false
}


// type APIError struct {
// 	RawError	error
// 	ErrStatus	string
// 	ErrCode		int
// 	ErrType		string
// 	Description string
// }


// // создание новой API ошибки
// func New(rawError error, description, errType string, errCode int) APIError {
// 	return APIError{
// 		RawError: rawError,
// 		ErrStatus: "apiError",
// 		ErrCode: errCode,
// 		ErrType: errType,
// 		Description: description,
// 	}
// }
// // создание новой API timeout ошибки
// func NewTimeout(rawError error, description, errType string, errCode int) APIError {
// 	apiErr := New(rawError, description, errType, errCode)
// 	apiErr.ErrStatus = "timeout"
// 	return apiErr
// }



// // реализация интерфейса error
// func (this APIError) Error() string {
// 	return this.RawError.Error()
// }

// // приведение ошибки к типу APIError
// func AssertAPIError(err error) APIError {
// 	apiErr := new(APIError)
	
// 	// если ошибка является APIError структурой
// 	if errors.As(err, apiErr) {
// 		return *apiErr
// 	}
// 	// разворачивание ошибки до самой первой
// 	if errors.As(errors.Cause(err), apiErr) {
// 		return *apiErr
// 	}

// 	// если ошибка не является APIError структурой, то делаем её из неё
// 	return New(err, err.Error(), "unknown", 500)
// }









// // ошибка внутри SSE функции
// type SseError string

// func (this SseError) Error() string {
// 	return string(this)
// }

// func NewSseErrorf(format string, a ...any) error {
// 	return SseError(fmt.Sprintf(format, a...))
// }
