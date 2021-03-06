// Code generated by sqlc. DO NOT EDIT.

package repository

import (
	"context"
)

type Querier interface {
	CreateCar(ctx context.Context, arg CreateCarParams) (Car, error)
	CreateOwner(ctx context.Context, arg CreateOwnerParams) (Owner, error)
	CreateWorkOrder(ctx context.Context, arg CreateWorkOrderParams) (WorkOrder, error)
	EndWorkOrder(ctx context.Context, id int32) error
	EndWorkOrderServiceExecution(ctx context.Context, workOrderID int64) error
	GetCarByLicensePlate(ctx context.Context, licensePlate string) (Car, error)
	GetCarsByOwner(ctx context.Context, ownerID int64) ([]Car, error)
	GetOwnerByEmail(ctx context.Context, email string) (Owner, error)
	GetOwnerByID(ctx context.Context, id int32) (Owner, error)
	GetOwnerByNationalID(ctx context.Context, nationalID string) (Owner, error)
	GetRunningServices(ctx context.Context, workOrderID int64) ([]WorkOrderServiceExecution, error)
	GetRunningWorkOrders(ctx context.Context) ([]GetRunningWorkOrdersRow, error)
	RegisterWorkOrderServiceExecution(ctx context.Context, arg RegisterWorkOrderServiceExecutionParams) (WorkOrderServiceExecution, error)
	UpdateWorkOrderServiceStatus(ctx context.Context, arg UpdateWorkOrderServiceStatusParams) error
}

var _ Querier = (*Queries)(nil)
