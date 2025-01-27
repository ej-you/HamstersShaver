// package errors


// // ошибка от моего REST API
// type RestAPIError string

// func (this RestAPIError) Error() string {
// 	return string(this)
// }

// // timeout ошибка от моего REST API
// type RestAPITimeoutError string

// func (this RestAPITimeoutError) Error() string {
// 	return string(this)
// }



// // ошибка от Redis
// type RedisError string

// func (this RedisError) Error() string {
// 	return string(this)
// }


// // ошибка от БД
// type DBError string

// func (this DBError) Error() string {
// 	return string(this)
// }

// // ошибка от БД при отсутствии результатов по поиску
// type DBNotFoundError string

// func (this DBNotFoundError) Error() string {
// 	return string(this)
// }


// // ошибка от валидации данных от юзера в боте
// type ValidateError string

// func (this ValidateError) Error() string {
// 	return string(this)
// }

// // ошибка при попытке юзером отправить транзакцию, когда ещё не завершилась предыдущая
// type LastTransNotFinishedError string

// func (this LastTransNotFinishedError) Error() string {
// 	return string(this)
// }

// // неизвестная ошибка в боте
// type InternalError string

// func (this InternalError) Error() string {
// 	return string(this)
// }


// // ошибка доступа к боту
// type AccessError string

// func (this AccessError) Error() string {
// 	return string(this)
// }


package errors

import (
	"context"
	"github.com/pkg/errors"
)


type APIError struct {
	RawError	error
	ErrStatus	string
	ErrCode		int
	ErrType		string
	Description string
}


// создание новой API ошибки
func New(rawError error, description, errType string, errCode int) APIError {
	return APIError{
		RawError: rawError,
		ErrStatus: "error",
		ErrCode: errCode,
		ErrType: errType,
		Description: description,
	}
}
// создание новой API timeout ошибки
func NewTimeout(rawError error, description, errType string, errCode int) APIError {
	apiErr := New(rawError, description, errType, errCode)
	apiErr.ErrStatus = "timeout"
	return apiErr
}


// проверка ошибки на timeout ошибку (если timeout произошёл из-за контекста) и корректирование статуса APIError, если это так
func (this *APIError) CheckTimeout() {
	timeoutErr := context.DeadlineExceeded
	
	// разворачивание ошибки для логов до самой первой
	if errors.Is(errors.Cause(this.RawError), timeoutErr) {
		this.ErrStatus = "timeout"
	}
}


// реализация интерфейса error
func (this APIError) Error() string {
	return this.RawError.Error()
}

// приведение ошибки к типу APIError
func AssertAPIError(err error) APIError {
	apiErr := new(APIError)
	
	// если ошибка является APIError структурой
	if errors.As(err, apiErr) {
		return *apiErr
	}
	// разворачивание ошибки до самой первой
	if errors.As(errors.Cause(err), apiErr) {
		return *apiErr
	}

	// если ошибка не является APIError структурой, то делаем её из неё
	return New(err, err.Error(), "unknown", 500)
}
