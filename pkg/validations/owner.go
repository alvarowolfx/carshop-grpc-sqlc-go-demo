package validations

import (
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RegisterOwnerParams struct {
	Email      string `validate:"required,email"`
	NationalID string `validate:"required,nationalID" json:"national_id,omitempty"`
}

func RegisterOwner(param *RegisterOwnerParams) []*ErrorResponse {
	return validateStruct(param)
}

func nationalIDValidation(nationalId string) error {
	if len(nationalId) != 11 {
		return status.Errorf(codes.InvalidArgument, "%s must have 11 digits", nationalId)
	}
	return nil
}

func ValidateNationalID(fl validator.FieldLevel) bool {
	err := nationalIDValidation(fl.Field().String())
	return err == nil
}
