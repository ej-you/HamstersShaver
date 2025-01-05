package ton_api_sse


// ошибка для SSE запросов
type SseError string

func (this SseError) Error() string {
	return string(this)
}
