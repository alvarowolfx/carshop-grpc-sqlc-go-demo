package controllers

import (
	"context"
	"database/sql"
	"net/http"

	"com.aviebrantz.carshop/pkg/backoffice/domain"
	carshop "com.aviebrantz.carshop/pkg/common/api"
	"com.aviebrantz.carshop/pkg/common/converters"
	"com.aviebrantz.carshop/pkg/common/repository"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BackOfficeController interface {
	carshop.BackOfficeServiceServer
}

type ControllerDeps struct {
	DB repository.Querier
}

type backOfficeController struct {
	deps    ControllerDeps
	client  *http.Client
	encoder *jsonpb.Marshaler
}

// NewController
func NewController(deps ControllerDeps) BackOfficeController {
	return &backOfficeController{
		deps: deps,
	}
}

func (cs *backOfficeController) RegisterOwner(ctx context.Context, param *carshop.Owner) (*empty.Empty, error) {
	errs := domain.RegisterOwner(&domain.RegisterOwnerParams{
		Email:      param.Email,
		NationalID: param.NationalId,
	})
	if len(errs) != 0 {
		err := converters.FromErrorResponsesToGrpcError("invalid parameters for owner registration", errs)
		return nil, err
	}
	_, err := cs.deps.DB.CreateOwner(ctx, repository.CreateOwnerParams{
		Email:      param.Email,
		NationalID: param.NationalId,
	})
	if err != nil {
		return nil, converters.CheckUniqueConstraintError(err)
	}
	return &empty.Empty{}, err
}

func (cs *backOfficeController) RegisterCar(ctx context.Context, param *carshop.Car) (*empty.Empty, error) {
	errs := domain.RegisterCar(&domain.RegisterCarParam{
		LicensePlate: param.LicensePlate,
		OwnerID:      param.OwnerId,
		Size:         int32(param.Size),
		NumWheels:    param.NumWheels,
		Color:        param.Color,
	})
	if len(errs) != 0 {
		err := converters.FromErrorResponsesToGrpcError("invalid parameters for car registration", errs)
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

	return &empty.Empty{}, converters.CheckUniqueConstraintError(err)
}
