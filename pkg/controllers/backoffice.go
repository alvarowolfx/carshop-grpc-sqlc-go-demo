package controllers

import (
	"context"
	"database/sql"
	"net/http"

	carshop "com.aviebrantz.carshop/api"
	"com.aviebrantz.carshop/pkg/repository"
	"com.aviebrantz.carshop/pkg/validations"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BackOfficeController interface {
	carshop.BackOfficeServiceServer
}

type BackOfficeControllerDeps struct {
	DB repository.Querier
}

type backOfficeController struct {
	deps    BackOfficeControllerDeps
	client  *http.Client
	encoder *jsonpb.Marshaler
}

// NewBackOfficeController
func NewBackOfficeController(deps BackOfficeControllerDeps) BackOfficeController {
	return &backOfficeController{
		deps: deps,
	}
}

func (cs *backOfficeController) RegisterOwner(ctx context.Context, param *carshop.Owner) (*empty.Empty, error) {
	errs := validations.RegisterOwner(&validations.RegisterOwnerParams{
		Email:      param.Email,
		NationalID: param.NationalId,
	})
	if len(errs) != 0 {
		err := validations.FromErrorResponsesToGrpcError("invalid parameters for owner registration", errs)
		return nil, err
	}
	_, err := cs.deps.DB.CreateOwner(ctx, repository.CreateOwnerParams{
		Email:      param.Email,
		NationalID: param.NationalId,
	})
	if err != nil {
		return nil, validations.CheckUniqueConstraintError(err)
	}
	return &empty.Empty{}, err
}

func (cs *backOfficeController) RegisterCar(ctx context.Context, param *carshop.Car) (*empty.Empty, error) {
	errs := validations.RegisterCar(&validations.RegisterCarParam{
		LicensePlate: param.LicensePlate,
		OwnerID:      param.OwnerId,
		Size:         int32(param.Size),
		NumWheels:    param.NumWheels,
		Color:        param.Color,
	})
	if len(errs) != 0 {
		err := validations.FromErrorResponsesToGrpcError("invalid parameters for car registration", errs)
		return nil, err
	}

	_, err := cs.deps.DB.GetOwnerByID(ctx, int32(param.OwnerId))
	if err == sql.ErrNoRows {
		return nil, status.Error(codes.InvalidArgument, "Owner not found")
	}

	if err != nil {
		return nil, err
	}

	carSize := []repository.CarSize{repository.CarSizeSmall, repository.CarSizeMedium, repository.CarSizeLarge}[int(param.Size)]
	_, err = cs.deps.DB.CreateCar(ctx, repository.CreateCarParams{
		LicensePlate: param.LicensePlate,
		Size:         carSize,
		NumWheels:    int16(param.NumWheels),
		Color:        param.Color,
		OwnerID:      param.OwnerId,
	})
	return &empty.Empty{}, validations.CheckUniqueConstraintError(err)
}
