package serializers

import (
	validate "github.com/gobuffalo/validate/v3"
)


// структура входных данных для отправки транзакции на продажу
type CellSendIn struct {
	JettonCA string `json:"jettonCA" myvalid:"required" example:"EQC47093oX5Xhb0xuk2lCr2RhS8rj-vul61u4W2UH5ORmG_O" description:"мастер-адрес продаваемой монеты (jetton_master)"`
	Amount float64 `json:"amount" myvalid:"required" example:"200" description:"кол-во используемых монет на продажу в формате, удобном для человека"`
	Slippage int `json:"slippage" myvalid:"required|min:0|max:100" example:"20" description:"процент проскальзывания"`
}

// дополнительная валидация входных данных (обязательный метод для всей валидации)
func (self *CellSendIn) IsValid(errors *validate.Errors) {}


// успешная отправка транзакции на продажу
type CellSendOut struct {
	Success bool `json:"success" example:"true" description:"успех"`
}
