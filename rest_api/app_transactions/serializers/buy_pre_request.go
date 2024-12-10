package serializers

import (
	validate "github.com/gobuffalo/validate/v3"
)


// структура входных данных для получения информации о последующей транзакции покупки
type BuyPreRequestIn struct {
	JettonCA string `query:"jettonCA" json:"jettonCA" myvalid:"required"`
	Amount float64 `query:"amount" json:"amount" myvalid:"required"`
	Slippage int `query:"slippage" json:"slippage" myvalid:"required|min:1|max:100"`
}

// дополнительная валидация входных данных (обязательный метод для всей валидации)
func (self *BuyPreRequestIn) IsValid(errors *validate.Errors) {}
