package errors

import "fmt"


// ошибка внутри SSE функции
type SseError string

func (this SseError) Error() string {
	return string(this)
}

func NewSseErrorf(format string, a ...any) error {
	return SseError(fmt.Sprintf(format, a...))
}
