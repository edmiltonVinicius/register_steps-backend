package domain

import (
	"github.com/edmiltonVinicius/register-steps/api/handlers/contracts"
	"github.com/go-playground/validator/v10"
)

var (
	ValidationErrors validator.ValidationErrors
	Validate         = validator.New()
)

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	case "email":
		return "Should be a email valid"
	case "min":
		return "Should has min " + fe.Param() + " characters"
	case "max":
		return "Should has max " + fe.Param() + " characters"
	case "contains":
		return "Should contains " + fe.Param() + " characters"
	case "eqfield":
		return fe.Field() + " should be equal to " + fe.Param()
	case "eq":
		return fe.Field() + " should be equal to " + fe.Param()
	case "containsany":
		return fe.Field() + " should contain any characters by type " + fe.Param()
	}
	return "Unknown error"
}

func RunValidator(errs validator.ValidationErrors) (msgs []contracts.ContractError) {
	for _, err := range errs {
		msgs = append(msgs, contracts.ContractError{
			Field:   err.Field(),
			Message: getErrorMsg(err),
		})
	}
	return
}
