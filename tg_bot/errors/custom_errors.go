package errors


// ошибка от моего REST API
type RestAPIError string

func (this RestAPIError) Error() string {
	return string(this)
}

// timeout ошибка от моего REST API
type RestAPITimeoutError string

func (this RestAPITimeoutError) Error() string {
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

// ошибка от БД при отсутствии результатов по поиску
type DBNotFoundError string

func (this DBNotFoundError) Error() string {
	return string(this)
}


// ошибка от валидации данных от юзера в боте
type ValidateError string

func (this ValidateError) Error() string {
	return string(this)
}

// ошибка при попытке юзером отправить транзакцию, когда ещё не завершилась предыдущая
type LastTransNotFinishedError string

func (this LastTransNotFinishedError) Error() string {
	return string(this)
}

// неизвестная ошибка в боте
type InternalError string

func (this InternalError) Error() string {
	return string(this)
}


// ошибка доступа к боту
type AccessError string

func (this AccessError) Error() string {
	return string(this)
}
