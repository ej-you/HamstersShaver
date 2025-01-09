package serializers

import (
	validate "github.com/gobuffalo/validate/v3"
)


// структура входных данных для получения кол-ва TON проценту от их баланса
type TonAmountFromPercentIn struct {
	Percent int `query:"percent" json:"percent" myvalid:"required|min:1|max:100"`
}

// дополнительная валидация входных данных (обязательный метод для всей валидации)
func (self *TonAmountFromPercentIn) IsValid(errors *validate.Errors) {}

// структура выходных данных получения кол-ва TON по проценту от их баланса
type TonAmountFromPercentOut struct {
	TonAmount string `json:"tonAmount" example:"1.533915351" description:"строковое кол-во TON, эквивалентное проценту от баланса"`
}
