package api

import (
	"github.com/ShadrackAdwera/ticket-assignment/utils"
	"github.com/go-playground/validator/v10"
)

var isValidAgentStatus validator.Func = func(fl validator.FieldLevel) bool {
	if status, ok := fl.Field().Interface().(string); ok {
		return utils.IsValidAgentStatus(status)
	}
	return false
}
