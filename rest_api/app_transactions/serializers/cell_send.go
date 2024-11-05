package serializers

import (
	validate "github.com/gobuffalo/validate/v3"
)


// @Description Cтруктура входных данных для отправки транзакции на продажу
type CellSendIn struct {
	// мастер-адрес продаваемой монеты (jetton_master)
	JettonCA string `json:"jettonCA" myvalid:"required" example:"EQC47093oX5Xhb0xuk2lCr2RhS8rj-vul61u4W2UH5ORmG_O"`
	// кол-во используемых монет на продажу в формате, удобном для человека
	Amount float64 `json:"amount" myvalid:"required" example:"200"`
	// процент проскальзывания 
	Slippage int `json:"slippage" myvalid:"required|min:0|max:100" example:"20"`
}

// дополнительная валидация входных данных (обязательный метод для всей валидации)
func (self *CellSendIn) IsValid(errors *validate.Errors) {}


// @Description Успешная отправка транзакции на продажу
type CellSendOut struct {
	// успех
	Success bool `json:"success" example:"true"`
}
