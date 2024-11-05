package validator

import (
	"fmt"
	"reflect"
	"strings"
	"strconv"

	echo "github.com/labstack/echo/v4"
	validate "github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
)


// интерфейс валидируемыемых структур с входными данными
type inputDataToValidate interface {
	IsValid(errors *validate.Errors)
}


// объединение *validate.Errors в одну ошибку *echo.HTTPError
func collectHttpError(validateErrors *validate.Errors) error {
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


// возврат ошибок валидации
func Validate(dataToValidate inputDataToValidate) error {
	validateErrors := validate.NewErrors()

	// базовая валидация входных данных по тегам структуры
	baseValidator(dataToValidate, validateErrors)
	// если базовая валидация не прошла, то возвращаем ошибку
	if err := collectHttpError(validateErrors); err != nil {
		return err
	}

	// дополнительная валидация входных данных (в методе IsValid у структуры)
	validateErrors = validate.Validate(dataToValidate)
	// если дополнительная валидация не прошла, то возвращаем ошибку
	if err := collectHttpError(validateErrors); err != nil {
		return err
	}
	return nil
}


func baseValidator(givenStruct validate.Validator, errors *validate.Errors) {
	// Получаем значение структуры (с разыменовыванием поинтера через Elem())
	structValue := reflect.ValueOf(givenStruct).Elem()
	// получаем кол-во полей в структуре
	structNumFields := structValue.NumField()

	var field reflect.StructField
	var tag string
	// перебираем поля и проверяем теги структуры
	for i:=0; i < structNumFields; i++ {
		// полная информация по полю (название, тип, значение и т.д.)
		field = structValue.Type().Field(i)

		// значение тега myvalid
		tag = field.Tag.Get("myvalid")
		if tag != "" {
			// перебираем значения тега myvalid
			for _, tagValue := range strings.Split(tag, "|") {
				switch {
					// обязательное поле
					case tagValue == "required":
						// проверка значения поля на ненулевое
						if structValue.Field(i).String() == "" {
							errors.Add(field.Tag.Get("json"), fmt.Sprintf("%s field must not be blank", field.Name))
						}

					// валидация email
					case tagValue == "email":
						// валидация средствами библиотеки
						errors.Append(validate.Validate(
							&validators.EmailIsPresent{
								Name: "Email",  // название поля
								Field: structValue.Field(i).String(),  // значение поля
								Message: "Email is not in the right format",  // сообщение при ошибке валидации
							},
						))

					// длина больше чем ... (пример, "minLen:8")
					case strings.HasPrefix(tagValue, "minLen"):
						// парсинг минимальной длины из тега
						minInt, _ := strconv.Atoi(strings.TrimPrefix(tagValue, "minLen:"))

						// проверка типа поля на string и его соответствие длины
						if field.Type == reflect.TypeOf("") && len(structValue.Field(i).String()) < minInt {
							errors.Add(field.Tag.Get("json"), fmt.Sprintf("%s field must contain at least %d symbols", field.Name, minInt))
						}

					// длина меньше чем ... (пример, "maxLen:100")
					case strings.HasPrefix(tagValue, "maxLen"):
						// парсинг максимальной длины из тега
						maxInt, _ := strconv.Atoi(strings.TrimPrefix(tagValue, "maxLen:"))

						// проверка типа поля на string и его соответствие длины
						if field.Type == reflect.TypeOf("") && len(structValue.Field(i).String()) > maxInt {
							errors.Add(field.Tag.Get("json"), fmt.Sprintf("%s field must contain less than %d symbols", field.Name, maxInt))
						}

					// число (int) больше чем ... (пример, "min:8")
					case strings.HasPrefix(tagValue, "min"):
						// парсинг минимального значения из тега
						minInt, _ := strconv.ParseInt(strings.TrimPrefix(tagValue, "min:"), 10, 64)

						// проверка типа поля на int и сравнение с минимальным порогом
						if field.Type.String() == "int" && structValue.Field(i).Int() < minInt {
							errors.Add(field.Tag.Get("json"), fmt.Sprintf("%s field must be greater than or equal to %d", field.Name, minInt))
						}

					// число (int) меньше чем ... (пример, "max:100")
					case strings.HasPrefix(tagValue, "max"):
						// парсинг максимального значения из тега
						maxInt, _ := strconv.ParseInt(strings.TrimPrefix(tagValue, "max:"), 10, 64)

						// проверка типа поля на int и сравнение с максимальным порогом
						if field.Type.String() == "int" && structValue.Field(i).Int() > maxInt {
							errors.Add(field.Tag.Get("json"), fmt.Sprintf("%s field must be less than or equal to %d", field.Name, maxInt))
						}
				}
			}
		}		
	}
}
