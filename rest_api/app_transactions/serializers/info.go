package serializers

import (
	validate "github.com/gobuffalo/validate/v3"
)


// @Description Cтруктура входных данных для получения информации о прошедшей транзакции по её хэшу
type InfoIn struct {
	TransactionHash string `query:"transactionHash" json:"transactionHash" myvalid:"required"`
	Action string `query:"action" json:"action" myvalid:"required"`
}

// дополнительная валидация входных данных (обязательный метод для всей валидации)
func (self *InfoIn) IsValid(errors *validate.Errors) {
	if self.Action != "buy" && self.Action != "cell" {
		errors.Add("action", "Invalid action parameter was given. Only buy and cell are accepted")
	}
}
