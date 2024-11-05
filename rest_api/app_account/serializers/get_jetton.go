package serializers

import (
	validate "github.com/gobuffalo/validate/v3"
)


// @Description Структура входных данных для получения информации о монете аккаунта по её адресу
type GetJettonIn struct {
	// мастер-адрес монеты (jetton_master)
	MasterAddress string `query:"masterAddress" json:"masterAddress" myvalid:"required" example:"EQC47093oX5Xhb0xuk2lCr2RhS8rj-vul61u4W2UH5ORmG_O"`
}

// дополнительная валидация входных данных (обязательный метод для всей валидации)
func (self *GetJettonIn) IsValid(errors *validate.Errors) {}
