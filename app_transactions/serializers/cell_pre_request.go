package serializers

import (
	echo "github.com/labstack/echo/v4"
	validate "github.com/gobuffalo/validate/v3"

	coreValidator "github.com/Danil-114195722/HamstersShaver/core/validator"
)


// структура входных данных для получения информации о последующей транзакции
type CellPreRequestIn struct {
	JettonCA string `json:"jettonCA" myvalid:"required"`
	Amount float64 `json:"amount" myvalid:"required"`
	Slippage int `json:"slippage" myvalid:"required|min:0|max:100"`
}


// базовая валидация полей по тегам
func (self *CellPreRequestIn) IsValid(errors *validate.Errors) {
	coreValidator.BaseValidator(self, errors)
}

// более глубокая валидация с возвратом ошибок валидации
func (self *CellPreRequestIn) Validate() error {
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
