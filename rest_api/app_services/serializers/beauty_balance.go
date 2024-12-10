package serializers

import (
	validate "github.com/gobuffalo/validate/v3"
)


// структура входных данных для получения преобразованного баланса монеты из int64 формата в округлённый float в виде строки
type BeautyBalanceIn struct {
	RawBalance int64 `query:"rawBalance" json:"rawBalance" myvalid:"required"`
	Decimals int `query:"decimals" json:"decimals" myvalid:"required|min:1"`
}

// дополнительная валидация входных данных (обязательный метод для всей валидации)
func (self *BeautyBalanceIn) IsValid(errors *validate.Errors) {}

// структура выходных данных получения преобразованного баланса монеты из int64 формата в округлённый float в виде строки
type BeautyBalanceOut struct {
	BeautyBalance string `json:"beautyBalance" example:"326.167" description:"преобразованный баланс монеты из int64 формата в округлённый float в виде строки"`
}
