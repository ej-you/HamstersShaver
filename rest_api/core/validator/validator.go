package validator

import (
	"sync"

	uniTrans "github.com/go-playground/universal-translator"
	validatorModule "github.com/go-playground/validator/v10"
	
	myValidatorModule "github.com/ej-you/go-utils/validator"
)


var (
	onceTranslator sync.Once
	onceValidator sync.Once

	// структура для валидации входных данных
	validator *validatorModule.Validate
	// "переводчик" для обработки сообщений ошибок
	translator *uniTrans.Translator
)


// получение "переводчика" для обработки сообщений ошибок
func GetTranslator() *uniTrans.Translator {
	onceTranslator.Do(func() {
		translator = myValidatorModule.GetTranslator()
	})
	return translator
}

// получение структуры для валидации входных данных
func GetValidator() *validatorModule.Validate {
	onceValidator.Do(func() {
		validator = myValidatorModule.GetValidator(GetTranslator())
	})
	return validator
}
