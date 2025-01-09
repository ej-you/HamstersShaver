package serializers

import (
	validate "github.com/gobuffalo/validate/v3"
)


// структура входных данных для получения кол-ва монет по проценту от их баланса
type JettonAmountFromPercentIn struct {
	Percent int `query:"percent" json:"percent" myvalid:"required|min:1|max:100"`
	MasterAddress string `query:"masterAddress" json:"masterAddress" myvalid:"required"`
}

// дополнительная валидация входных данных (обязательный метод для всей валидации)
func (self *JettonAmountFromPercentIn) IsValid(errors *validate.Errors) {}

// структура выходных данных получения кол-ва монет по проценту от их баланса
type JettonAmountFromPercentOut struct {
	JettonAmount string `json:"jettonAmount" example:"124.533915351" description:"строковое кол-во монет, эквивалентное проценту от их баланса"`
}
