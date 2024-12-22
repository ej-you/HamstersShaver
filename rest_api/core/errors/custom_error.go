package errors

import (
	"github.com/pkg/errors"

	echo "github.com/labstack/echo/v4"
)


type APIError struct {
	rawError	error
	description string
	errType		string
	errCode		int
}


// создание новой API ошибки
func New(rawError error, description, errType string, errCode int) APIError {
	return APIError{
		rawError: rawError,
		description: description,
		errType: errType,
		errCode: errCode,
	}
}


// реализация интерфейса error
func (this APIError) Error() string {
	return this.rawError.Error()
}

// получение HTTP ошибки
func (this APIError) GetHTTPError() error {
	return echo.NewHTTPError(this.errCode, map[string]string{this.errType: this.description})
}


// приведение ошибки к типу APIError
func AssertAPIError(err error) APIError {
	apiErr := new(APIError)
	
	// ели ошибка является APIError структурой
	if errors.As(err, apiErr) {
		return *apiErr
	}
	// разворачивание ошибки до самой первой
	if errors.As(errors.Cause(err), apiErr) {
		return *apiErr
	}

	return New(err, err.Error(), "unknown", 500)
}
