package serializers

import (
	validate "github.com/gobuffalo/validate/v3"
)


// @Description Cтруктура входных данных для получения информации о прошедшей транзакции по её хэшу
type InfoIn struct {
	// хэш транзакции
	TransactionHash string `query:"transactionHash" json:"transactionHash" myvalid:"required" example:"29a301e4d2a05713f4eab6c8f0daa3c58eed15d1d41678068cd50fe46ca7f6a5"`
	// действие с монетами в транзакции (покупка/продажа)
	Action string `query:"action" json:"action" myvalid:"required" example:"cell"`
}

// дополнительная валидация входных данных (обязательный метод для всей валидации)
func (self *InfoIn) IsValid(errors *validate.Errors) {
	if self.Action != "buy" && self.Action != "cell" {
		errors.Add("action", "Invalid action parameter was given. Only buy and cell are accepted")
	}
}
