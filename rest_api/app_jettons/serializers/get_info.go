package serializers

import (
	echo "github.com/labstack/echo/v4"
	validate "github.com/gobuffalo/validate/v3"

	coreValidator "github.com/ej-you/HamstersShaver/rest_api/core/validator"
)


// @Description Структура входных данных для получения информации о монете по её адресу
type GetInfoIn struct {
	// мастер-адрес монеты (jetton_master)
	MasterAddress string `query:"masterAddress" json:"masterAddress" myvalid:"required" example:"EQC47093oX5Xhb0xuk2lCr2RhS8rj-vul61u4W2UH5ORmG_O"`
}


// базовая валидация полей по тегам
func (self *GetInfoIn) IsValid(errors *validate.Errors) {
	coreValidator.BaseValidator(self, errors)
}

// более глубокая валидация с возвратом ошибок валидации
func (self *GetInfoIn) Validate() error {
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
