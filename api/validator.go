package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/karan0033/bank/utils"
)

var validateCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return utils.IsSupportedCurrency(currency) // Note the uppercase 'I'
	}
	return false
}
