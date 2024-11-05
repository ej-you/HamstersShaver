package serializers

import (
	validate "github.com/gobuffalo/validate/v3"
)


// @Description Cтруктура входных данных для получения информации о последующей транзакции покупки
type BuyPreRequestIn struct {
	// мастер-адрес покупаемой монеты (jetton_master)
	JettonCA string `query:"jettonCA" json:"jettonCA" myvalid:"required" example:"EQC47093oX5Xhb0xuk2lCr2RhS8rj-vul61u4W2UH5ORmG_O"`
	// кол-во используемых TON для покупки в формате, удобном для человека
	Amount float64 `query:"amount" json:"amount" myvalid:"required" example:"0.1"`
	// процент проскальзывания 
	Slippage int `query:"slippage" json:"slippage" myvalid:"required|min:0|max:100" example:"20"`
}

// дополнительная валидация входных данных (обязательный метод для всей валидации)
func (self *BuyPreRequestIn) IsValid(errors *validate.Errors) {}
