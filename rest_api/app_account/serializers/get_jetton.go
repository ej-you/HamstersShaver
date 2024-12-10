package serializers

import (
	validate "github.com/gobuffalo/validate/v3"
)


// структура входных данных для получения информации о монете аккаунта по её адресу
type GetJettonIn struct {
	MasterAddress string `query:"masterAddress" json:"masterAddress" myvalid:"required"`
}

// дополнительная валидация входных данных (обязательный метод для всей валидации)
func (self *GetJettonIn) IsValid(errors *validate.Errors) {}
