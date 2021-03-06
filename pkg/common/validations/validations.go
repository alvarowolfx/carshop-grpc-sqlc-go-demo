package validations

import (
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	FailedField string `json:"field"`
	Value       string `json:"value"`
	Error       string `json:"error"`
}

var (
	validate = validator.New()
)

// ValidateStruct Check if struct have errors and return properly formatted errors
func ValidateStruct(requestStruct interface{}) []*ErrorResponse {
	var errors []*ErrorResponse

	err := validate.Struct(requestStruct)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.Field()
			element.Value = err.Param()
			//element.Error = err.ActualTag()
			element.Error = err.Tag()
			errors = append(errors, &element)
		}
	}
	return errors
}

func RegisterValidation(tag string, fn validator.Func) error {
	return validate.RegisterValidation(tag, fn)
}
