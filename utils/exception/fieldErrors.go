package exception

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func FieldErrors(err error) map[string]string {
	errorMap := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, ve := range validationErrors {
			fieldName := ve.Field()
			tagName := ve.Tag()
			paramValue := ve.Param()

			switch tagName {
			case "required":
				errorMap[fieldName] = fmt.Sprintf("%s is required", fieldName)
			case "min":
				errorMap[fieldName] = fmt.Sprintf("%s must be at least %s characters", fieldName, paramValue)
			case "max":
				errorMap[fieldName] = fmt.Sprintf("%s must not exceed %s characters", fieldName, paramValue)
			case "numeric":
				errorMap[fieldName] = fmt.Sprintf("%s must be numeric", fieldName)
			case "alpha":
				errorMap[fieldName] = fmt.Sprintf("%s must contain only alphabetic characters", fieldName)
			case "alphanum":
				errorMap[fieldName] = fmt.Sprintf("%s must be alphanumeric", fieldName)
			case "len":
				errorMap[fieldName] = fmt.Sprintf("%s must be exactly %s characters", fieldName, paramValue)
			default:
				errorMap[fieldName] = fmt.Sprintf("%s is invalid", fieldName)
			}
		}
	}

	return errorMap
}
