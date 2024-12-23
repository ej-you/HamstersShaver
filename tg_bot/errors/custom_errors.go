package errors


// ошибка от моего REST API
type RestAPIError string

func (this RestAPIError) Error() string {
	return string(this)
}


// ошибка от Redis
type RedisError string

func (this RedisError) Error() string {
	return string(this)
}


// ошибка от БД
type DBError string

func (this DBError) Error() string {
	return string(this)
}


// ошибка от валидации данных от юзера в боте
type ValidateError string

func (this ValidateError) Error() string {
	return string(this)
}


// неизвестная ошибка в боте
type InternalError string

func (this InternalError) Error() string {
	return string(this)
}
