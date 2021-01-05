package validations

import (
	"regexp"

	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RegisterCarParam struct {
	LicensePlate string `validate:"required,licensePlate"`
	OwnerID      int64  `validate:"required"`
	Size         int32  `validate:"gte=0,lte=2"`
	NumWheels    int32  `validate:"gte=2,lte=10"`
	Color        string `validate:"iscolor"`
}

func RegisterCar(param *RegisterCarParam) []*ErrorResponse {
	return validateStruct(param)
}

func ValidateLicensePlate(fl validator.FieldLevel) bool {
	err := licensePlateValidation(fl.Field().String())
	return err == nil
}

func licensePlateValidation(licensePlate string) error {
	ok, err := regexp.MatchString("^[A-Z]{3}-[0-9]{4}", licensePlate)
	if err != nil {
		return err
	}
	if !ok {
		return status.Errorf(codes.InvalidArgument, "%s should have format ABC-1234", licensePlate)
	}
	return nil

}
