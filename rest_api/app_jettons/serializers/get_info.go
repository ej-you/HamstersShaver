package serializers

import (
	validate "github.com/gobuffalo/validate/v3"
)


// @Description Структура входных данных для получения информации о монете по её адресу
type GetInfoIn struct {
	// мастер-адрес монеты (jetton_master)
	MasterAddress string `query:"masterAddress" json:"masterAddress" myvalid:"required"`
}

// дополнительная валидация входных данных (обязательный метод для всей валидации)
func (self *GetInfoIn) IsValid(errors *validate.Errors) {}
