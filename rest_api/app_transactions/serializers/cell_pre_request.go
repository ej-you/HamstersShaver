package serializers

import (
	validate "github.com/gobuffalo/validate/v3"
)


// структура входных данных для получения информации о последующей транзакции продажи
type CellPreRequestIn struct {
	JettonCA string `query:"jettonCA" json:"jettonCA" myvalid:"required"`
	Amount float64 `query:"amount" json:"amount" myvalid:"required"`
	Slippage int `query:"slippage" json:"slippage" myvalid:"required|min:0|max:100"`
}

// дополнительная валидация входных данных (обязательный метод для всей валидации)
func (self *CellPreRequestIn) IsValid(errors *validate.Errors) {}
