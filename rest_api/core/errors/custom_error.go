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
