package serializers

import (
	echo "github.com/labstack/echo/v4"
	validate "github.com/gobuffalo/validate/v3"

	coreValidator "github.com/ej-you/HamstersShaver/rest_api/core/validator"
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


// базовая валидация полей по тегам
func (self *BuyPreRequestIn) IsValid(errors *validate.Errors) {
	coreValidator.BaseValidator(self, errors)
}

// более глубокая валидация с возвратом ошибок валидации
func (self *BuyPreRequestIn) Validate() error {
	// базовая валидация полей по тегам
	var validateErrors *validate.Errors = validate.Validate(self)

	if len(validateErrors.Errors) > 0 {
		// словарь для ошибок
		errMap := make(map[string]string, len(validateErrors.Errors))

		for key, value := range validateErrors.Errors {
			errMap[key] = value[0]
		}
		// возвращаем *echo.HTTPError
		httpError := echo.NewHTTPError(400, errMap)
		return httpError
	}
	return nil
}
